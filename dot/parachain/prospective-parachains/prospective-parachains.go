package prospectiveparachains

import (
	"context"
	"errors"

	fragmentchain "github.com/ChainSafe/gossamer/dot/parachain/prospective-parachains/fragment-chain"
	parachaintypes "github.com/ChainSafe/gossamer/dot/parachain/types"
	"github.com/ChainSafe/gossamer/internal/log"
	"github.com/ChainSafe/gossamer/lib/common"
)

var logger = log.NewFromGlobal(log.AddContext("pkg", "prospective_parachains"), log.SetLevel(log.Debug))

type ProspectiveParachains struct {
	SubsystemToOverseer chan<- any
	View                *View
}

type RelayBlockViewData struct {
	fragmentChains map[parachaintypes.ParaID]fragmentchain.FragmentChain
}

type View struct {
	perRelayParent map[common.Hash]RelayBlockViewData
	activeLeaves   map[common.Hash]struct{}
	// implicitView   backing.ImplicitView
}

func (pp *ProspectiveParachains) AnswerGetBackableCandidates(
	view *View,
	relayParent common.Hash,
	para parachaintypes.ParaID,
	count uint32,
	ancestors Ancestors,
	tx chan []parachaintypes.CandidateHashAndRelayParent,
) {
	if _, ok := view.activeLeaves[relayParent]; !ok {
		logger.Tracef("Requested backable candidates for relay parent %s, but it is not an active leaf", relayParent)
		tx <- nil
		return
	}

	data, exists := view.perRelayParent[relayParent]

	if !exists {
		logger.Tracef("Requested backable candidates for relay parent %s, but it has no view data", relayParent)
		tx <- nil
		return
	}

	chain, exists := data.fragmentChains[para]

	if !exists {
		logger.Tracef("Requested backable candidates for relay parent %s and para %d, but no fragment chain exists", relayParent, para)
		tx <- nil
		return
	}

	logger.Tracef("Candidate chain for para %d: %s", para, chain.BestChainVec())
	logger.Tracef("Potential candidate storage for para %d: %s", para, chain.Unconnected())

	backableCandidates := chain.FindBackableChain(ancestors, count)

	if len(backableCandidates) == 0 {
		logger.Tracef("Could not find any backable candidate for para %d", para)
		tx <- nil
		return
	}

	logger.Tracef("Found backable candidates for para %d: %s", para, backableCandidates)

	candidateHashes := make([]parachaintypes.CandidateHashAndRelayParent, len(backableCandidates))

	for i, candidate := range backableCandidates {
		candidateHashes[i] = parachaintypes.CandidateHashAndRelayParent{
			CandidateHash:        candidate.CandidateHash,
			CandidateRelayParent: candidate.RealyParentHash,
		}
	}

	logger.Tracef("candidateHashesForDebug %v", candidateHashes) // just for test, i need to run to see the result

	tx <- candidateHashes
}

// Name returns the name of the subsystem
func (*ProspectiveParachains) Name() parachaintypes.SubSystemName {
	return parachaintypes.ProspectiveParachains
}

// NewProspectiveParachains creates a new ProspectiveParachain subsystem
func NewProspectiveParachains(overseerChan chan<- any) *ProspectiveParachains {
	prospectiveParachain := ProspectiveParachains{
		SubsystemToOverseer: overseerChan,
	}
	return &prospectiveParachain
}

// Run starts the ProspectiveParachains subsystem
func (pp *ProspectiveParachains) Run(ctx context.Context, overseerToSubsystem <-chan any) {
	for {
		select {
		case msg := <-overseerToSubsystem:
			pp.processMessage(msg)
		case <-ctx.Done():
			if err := ctx.Err(); err != nil && !errors.Is(err, context.Canceled) {
				logger.Errorf("ctx error: %s\n", err)
			}
			return
		}
	}
}

func (*ProspectiveParachains) Stop() {}

func (pp *ProspectiveParachains) processMessage(msg any) {
	switch msg := msg.(type) {
	case parachaintypes.Conclude:
		pp.Stop()
	case parachaintypes.ActiveLeavesUpdateSignal:
		_ = pp.ProcessActiveLeavesUpdateSignal(msg)
	case parachaintypes.BlockFinalizedSignal:
		_ = pp.ProcessBlockFinalizedSignal(msg)
	case IntroduceSecondedCandidate:
		panic("not implemented yet: see issue #4308")
	case CandidateBacked:
		panic("not implemented yet: see issue #4309")
	case GetBackableCandidates:
		pp.AnswerGetBackableCandidates(pp.View, msg.RelayParentHash, msg.ParaId, msg.Count, msg.Ancestors, msg.Response)
	case GetHypotheticalMembership:
		panic("not implemented yet: see issue #4311")
	case GetMinimumRelayParents:
		panic("not implemented yet: see issue #4312")
	case GetProspectiveValidationData:
		panic("not implemented yet: see issue #4313")
	default:
		logger.Errorf("%w: %T", parachaintypes.ErrUnknownOverseerMessage, msg)
	}

}

// ProcessActiveLeavesUpdateSignal processes active leaves update signal
func (pp *ProspectiveParachains) ProcessActiveLeavesUpdateSignal(parachaintypes.ActiveLeavesUpdateSignal) error {
	panic("not implemented yet: see issue #4305")
}

// ProcessBlockFinalizedSignal processes block finalized signal
func (*ProspectiveParachains) ProcessBlockFinalizedSignal(parachaintypes.BlockFinalizedSignal) error {
	// NOTE: this subsystem does not process block finalized signal
	return nil
}
