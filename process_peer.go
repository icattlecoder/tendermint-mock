package tendermint_mock

import "time"

type Receiver interface {
	Receive(interface{})
}

type ProcessPeer struct {
	PeerId int
	peer   Receiver
	Data   map[string]interface{}
	PeerRoundState
}

func (p *ProcessPeer) UpdateProposal(prop *Proposal) {
	p.Data["proposal"] = prop
}

func (p *ProcessPeer) GetProposal() (*Proposal, bool) {

	i, ok := p.Data["proposal"]
	if !ok {
		return nil, ok
	}
	return i.(*Proposal), true
}

func (p *ProcessPeer) Send(d interface{}) {
	time.Sleep(time.Millisecond * time.Duration(_rand(30)))
	p.peer.Receive(d)
}
