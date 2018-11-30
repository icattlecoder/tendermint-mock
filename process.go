package tendermint_mock

import (
	"math/rand"
	"fmt"
	"time"
)

type Process struct {
	Id       int
	network  *Network
	state    State
	proposal *Proposal
}

func _rand(i int) int {
	return rand.Int()%i + i
}

func (p *Process) Add(peer *ProcessPeer) {
	go p.gossipVotes(peer)
	peer.Send(fmt.Sprintf("proc %d say hello to %d", p.Id, peer.PeerId))
}

func (p *Process) Received(d interface{}) {
	fmt.Printf("proc %d received: \"%v\"\n", p.Id, d)
}

func (p *Process) NewRound(r int) {

}

func (p *Process) isMyTurn() bool {
	return p.Id == p.network.ValidatorId(p.state.Height, p.state.Round)
}

func (p *Process) doProposal() {

	if !p.isMyTurn() {
		return
	}
	//pl := Proposal{
	//	Height: 0,
	//}
	// broadcast pl
}

func (p *Process) doPrevote() {

	var dataId []byte
	if p.proposal == nil {
		dataId = nil
	} else {
		dataId = p.proposal.DataId
	}

	v := Vote{
		ValidatorId: p.Id,
		Type:        VoteTypePrevote,
		Height:      p.state.Height,
		Round:       p.state.Round,
		DataId:      dataId,
	}
	p.state.AddVote(&v)
}

func (p *Process) DoPrecommit() {

}

func (p *Process) gossipData(peer *ProcessPeer) {

	for {
		time.Sleep(time.Millisecond * 100)
		if p.proposal == nil {
			continue
		}

		prop, ok := peer.GetProposal()
		if ok || prop != p.proposal {
			continue
		}
		peer.Send(prop)
		peer.UpdateProposal(prop)
	}
}

func (p *Process) gossipVotes(peer *ProcessPeer) {

	for {
		v := subVotes(p.state.prevotes, peer.PeerRoundState.prevotes)
		if v != nil {
			peer.Send(v)
			peer.PeerRoundState.prevotes[v.ValidatorId] = true
		}
		v = subVotes(p.state.precommits, peer.PeerRoundState.precommits)
		if v != nil {
			peer.Send(v)
			peer.PeerRoundState.precommits[v.ValidatorId] = true
		}
	}
}
