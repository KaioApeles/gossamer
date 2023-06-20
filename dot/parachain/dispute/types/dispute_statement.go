package types

import (
	"fmt"

	"github.com/ChainSafe/gossamer/dot/parachain"
	parachainTypes "github.com/ChainSafe/gossamer/dot/parachain/types"
	"github.com/ChainSafe/gossamer/lib/babe/inherents"
	"github.com/ChainSafe/gossamer/pkg/scale"
)

// SecondedCompactStatement is the proposal of a parachain candidate.
type SecondedCompactStatement struct {
	CandidateHash parachainTypes.CandidateHash
}

// Index returns the index of the type SecondedCompactStatement.
func (SecondedCompactStatement) Index() uint {
	return 0
}

func (SecondedCompactStatement) SigningPayload() []byte {
	panic("implement me")
}

// ValidCompactStatement represents a valid candidate.
type ValidCompactStatement struct {
	CandidateHash parachainTypes.CandidateHash
}

// Index returns the index of the type ValidCompactStatement.
func (ValidCompactStatement) Index() uint {
	return 1
}

// CompactStatementVDT is the statement that can be made about parachain candidates
// These are the actual values that are signed.
type CompactStatementVDT scale.VaryingDataType

// Set will set a VaryingDataTypeValue using the underlying VaryingDataType
func (cs *CompactStatementVDT) Set(val scale.VaryingDataTypeValue) (err error) {
	vdt := scale.VaryingDataType(*cs)
	err = vdt.Set(val)
	if err != nil {
		return fmt.Errorf("setting value to varying data type: %w", err)
	}
	*cs = CompactStatementVDT(vdt)
	return nil
}

// Value returns the value from the underlying VaryingDataType
func (cs *CompactStatementVDT) Value() (val scale.VaryingDataTypeValue, err error) {
	vdt := scale.VaryingDataType(*cs)
	val, err = vdt.Value()
	if err != nil {
		return nil, fmt.Errorf("getting value from varying data type: %w", err)
	}
	return val, nil
}

func (cs *CompactStatementVDT) SigningPayload() ([]byte, error) {
	// scale encode the value
	encoded, err := scale.Marshal(*cs)
	if err != nil {
		return nil, fmt.Errorf("encode compact statement: %w", err)
	}

	return encoded, nil
}

// NewCompactStatement creates a new CompactStatementVDT.
func NewCompactStatement() CompactStatementVDT {
	vdt := scale.MustNewVaryingDataType(ValidCompactStatement{}, SecondedCompactStatement{})
	return CompactStatementVDT(vdt)
}

type CompactStatementKind uint

const (
	SecondedCompactStatementKind CompactStatementKind = iota
	ValidCompactStatementKind
)

func NewCustomCompactStatement(kind CompactStatementKind, candidateHash parachainTypes.CandidateHash) (CompactStatement, error) {
	vdt := NewCompactStatement()

	var err error
	switch kind {
	case SecondedCompactStatementKind:
		err = vdt.Set(SecondedCompactStatement{
			CandidateHash: candidateHash,
		})
	case ValidCompactStatementKind:
		err = vdt.Set(ValidCompactStatement{
			CandidateHash: candidateHash,
		})
	default:
		return CompactStatementVDT{}, fmt.Errorf("invalid compact statement kind")
	}

	if err != nil {
		return CompactStatementVDT{}, fmt.Errorf("set compact statement: %w", err)
	}

	return vdt, nil
}

// ExplicitDisputeStatement An explicit statement on a candidate issued as part of a dispute.
type ExplicitDisputeStatement struct {
	Valid         bool
	CandidateHash parachainTypes.CandidateHash
	Session       parachainTypes.SessionIndex
}

func (ExplicitDisputeStatement) SigningPayload() []byte {
	panic("implement me")
}

// ApprovalVote A vote of approval on a candidate.
type ApprovalVote struct {
	candidateHash parachainTypes.CandidateHash
}

func (a *ApprovalVote) SigningPayload() ([]byte, error) {
	encoded, err := scale.Marshal(&a)
	if err != nil {
		return nil, fmt.Errorf("encode approval vote: %w", err)
	}

	return encoded, nil
}

// SignedDisputeStatement A checked dispute statement from an associated validator.
type SignedDisputeStatement struct {
	DisputeStatement   inherents.DisputeStatement
	CandidateHash      parachainTypes.CandidateHash
	ValidatorPublic    parachainTypes.ValidatorID
	ValidatorSignature parachain.ValidatorSignature
	SessionIndex       parachainTypes.SessionIndex
}

func NewCheckedSignedDisputeStatement(disputeStatement inherents.DisputeStatement,
	candidateHash parachainTypes.CandidateHash,
	sessionIndex parachainTypes.SessionIndex,
	validatorPublic parachainTypes.ValidatorID,
	validatorSignature parachain.ValidatorSignature,
) (*SignedDisputeStatement, error) {
	payload, err := getDisputeStatementSigningPayload(disputeStatement, candidateHash, sessionIndex)
	if err != nil {
		return nil, fmt.Errorf("get dispute statement signing payload: %w", err)
	}

	if err := validatorSignature.Verify(payload, validatorPublic); err != nil {
		return nil, fmt.Errorf("verify validator signature: %w", err)
	}

	return &SignedDisputeStatement{
		DisputeStatement:   disputeStatement,
		CandidateHash:      candidateHash,
		ValidatorPublic:    validatorPublic,
		ValidatorSignature: validatorSignature,
		SessionIndex:       sessionIndex,
	}, nil
}

func getDisputeStatementSigningPayload(disputeStatement inherents.DisputeStatement,
	candidateHash parachainTypes.CandidateHash,
	session parachainTypes.SessionIndex,
) ([]byte, error) {
	statement, err := disputeStatement.Value()
	if err != nil {
		return nil, fmt.Errorf("failed to get dispute statement value: %w", err)
	}

	var payload []byte
	switch statement.(type) {
	case inherents.ExplicitValidDisputeStatementKind:
		data := ExplicitDisputeStatement{
			Valid:         true,
			CandidateHash: candidateHash,
			Session:       session,
		}
		payload = data.SigningPayload()
	case inherents.BackingSeconded:
		data, err := NewCustomCompactStatement(SecondedCompactStatementKind, candidateHash)
		if err != nil {
			return nil, fmt.Errorf("new custom compact statement: %w", err)
		}

		payload, err = data.SigningPayload()
		if err != nil {
			return nil, fmt.Errorf("signing payload: %w", err)
		}

	case inherents.BackingValid:
		data, err := NewCustomCompactStatement(ValidCompactStatementKind, candidateHash)
		if err != nil {
			return nil, fmt.Errorf("new custom compact statement: %w", err)
		}

		payload, err = data.SigningPayload()
		if err != nil {
			return nil, fmt.Errorf("signing payload: %w", err)
		}

	case inherents.ApprovalChecking:
		data := ApprovalVote{
			candidateHash: candidateHash,
		}
		payload, err = data.SigningPayload()
		if err != nil {
			return nil, fmt.Errorf("signing payload: %w", err)
		}

	case inherents.InvalidDisputeStatementKind:
		data := ExplicitDisputeStatement{
			Valid:         false,
			CandidateHash: candidateHash,
			Session:       session,
		}
		payload = data.SigningPayload()

	default:
		return nil, fmt.Errorf("invalid dispute statement kind")

	}

	return payload, nil
}

// Statement is the statement that can be made about parachain candidates.
type Statement struct {
	SignedDisputeStatement SignedDisputeStatement
	ValidatorIndex         parachainTypes.ValidatorIndex
}
