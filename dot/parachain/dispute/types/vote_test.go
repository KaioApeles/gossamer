package types

import (
	"testing"

	parachainTypes "github.com/ChainSafe/gossamer/dot/parachain/types"
	"github.com/ChainSafe/gossamer/pkg/scale"
	"github.com/stretchr/testify/require"
)

func Test_CandidateVotesCodec(t *testing.T) {
	t.Parallel()
	// with
	receipt := parachainTypes.CandidateReceipt{
		Descriptor: parachainTypes.CandidateDescriptor{
			ParaID:                      100,
			RelayParent:                 GetRandomHash(),
			Collator:                    parachainTypes.CollatorID{2},
			PersistedValidationDataHash: GetRandomHash(),
			PovHash:                     GetRandomHash(),
			ErasureRoot:                 GetRandomHash(),
			Signature:                   parachainTypes.CollatorSignature{2},
			ParaHead:                    GetRandomHash(),
			ValidationCodeHash:          parachainTypes.ValidationCodeHash(GetRandomHash()),
		},
		CommitmentsHash: GetRandomHash(),
	}

	validVotes := NewValidCandidateVotes(32)
	inserted, err := validVotes.InsertVote(Vote{
		ValidatorIndex:     1,
		DisputeStatement:   DummyInvalidDisputeStatement(t),
		ValidatorSignature: [64]byte{1},
	})
	require.NoError(t, err)
	require.True(t, inserted)
	inserted, err = validVotes.InsertVote(Vote{
		ValidatorIndex:     2,
		DisputeStatement:   DummyValidDisputeStatement(t),
		ValidatorSignature: [64]byte{2},
	})
	require.NoError(t, err)
	require.True(t, inserted)

	invalidVotes := NewInvalidCandidateVotes(32)
	invalidVotes.Set(2, Vote{
		ValidatorIndex:     2,
		DisputeStatement:   DummyInvalidDisputeStatement(t),
		ValidatorSignature: [64]byte{2},
	})
	invalidVotes.Set(3, Vote{
		ValidatorIndex:     3,
		DisputeStatement:   DummyInvalidDisputeStatement(t),
		ValidatorSignature: [64]byte{3},
	})

	votes := CandidateVotes{CandidateReceipt: receipt, Valid: validVotes, Invalid: invalidVotes}

	// when
	encoded, err := scale.Marshal(votes)
	require.NoError(t, err)

	decoded := NewCandidateVotes()
	err = scale.Unmarshal(encoded, &decoded)
	require.NoError(t, err)

	// then
	require.Equal(t, votes.CandidateReceipt, decoded.CandidateReceipt)
	require.Equal(t, votes.Valid.Value.Len(), decoded.Valid.Value.Len())
	require.Equal(t, votes.Invalid.Len(), decoded.Invalid.Len())
}

func Test_Vote(t *testing.T) {
	t.Parallel()
	validVote := Vote{
		ValidatorIndex:     1,
		DisputeStatement:   DummyValidDisputeStatement(t),
		ValidatorSignature: GetRandomSignature(),
	}

	encoded, err := scale.Marshal(validVote)
	require.NoError(t, err)

	decoded := Vote{}
	err = scale.Unmarshal(encoded, &decoded)
	require.NoError(t, err)

	require.Equal(t, validVote, decoded)

	invalidVote := Vote{
		ValidatorIndex:     1,
		DisputeStatement:   DummyInvalidDisputeStatement(t),
		ValidatorSignature: GetRandomSignature(),
	}

	encoded, err = scale.Marshal(invalidVote)
	require.NoError(t, err)

	decoded = Vote{}
	err = scale.Unmarshal(encoded, &decoded)
	require.NoError(t, err)

	require.Equal(t, invalidVote, decoded)
}

func TestOwnVoteState_CannotVote(t *testing.T) {
	t.Parallel()
	// with
	ownVoteState, err := NewOwnVoteStateVDT(CannotVote{})
	require.NoError(t, err)

	// when
	encoded, err := scale.Marshal(ownVoteState)
	require.NoError(t, err)

	decoded := OwnVoteStateVDT{}
	err = scale.Unmarshal(encoded, &decoded)
	require.NoError(t, err)

	// then
	require.Equal(t, ownVoteState, decoded)
}

func TestOwnVoteState_Voted(t *testing.T) {
	t.Parallel()
	// with
	votes := []Vote{
		{
			ValidatorIndex:     1,
			DisputeStatement:   DummyValidDisputeStatement(t),
			ValidatorSignature: GetRandomSignature(),
		},
		{
			ValidatorIndex:     2,
			DisputeStatement:   DummyInvalidDisputeStatement(t),
			ValidatorSignature: GetRandomSignature(),
		},
	}

	ownVoteState, err := NewOwnVoteStateVDT(Voted{Votes: votes})
	require.NoError(t, err)

	// when
	encoded, err := scale.Marshal(ownVoteState)
	require.NoError(t, err)

	decoded, err := NewOwnVoteStateVDT(CannotVote{})
	require.NoError(t, err)
	err = scale.Unmarshal(encoded, &decoded)
	require.NoError(t, err)

	// then
	require.Equal(t, ownVoteState, decoded)
}
