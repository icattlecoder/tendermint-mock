package tendermint_mock

type Network struct {
	Processes    []Process
	ValidatorsId []int
}

func NewNetwork(cap int) *Network {
	n := Network{
		Processes:    make([]Process, 0, cap),
		ValidatorsId: make([]int, 0, cap),
	}
	return &n
}

func (n *Network) ValidatorId(h, r int) int {
	return n.ValidatorsId[(h+r)%len(n.ValidatorsId)]
}

func (n *Network) AddNewProcess() *Process {

	p := Process{Id: len(n.Processes), network: n}
	for i, v := range n.Processes {
		peer := ProcessPeer{PeerId: p.Id, peer: &p, Data: make(map[string]interface{})}
		v.Add(&peer)
		peer2 := ProcessPeer{PeerId: v.Id, peer: &n.Processes[i], Data: make(map[string]interface{})}
		p.Add(&peer2)
	}
	n.Processes = append(n.Processes, p)
	n.ValidatorsId = append(n.ValidatorsId, p.Id)
	return &p
}
