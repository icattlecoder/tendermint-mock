package tendermint_mock

type State struct {
	Height int
	Round  int
	VoteSet
}

type PeerRoundState struct {
	VoteSetMap
}

type Proposal struct {
	Height int
	Round  int
	Data   []byte
	DataId []byte
}

type VoteType int

const (
	VoteTypePrevote   VoteType = iota
	VoteTypePrecommit
)

type Vote struct {
	ValidatorId int
	Height      int
	Round       int
	Type        VoteType
	DataId      []byte
}

type VoteSet struct {
	prevotes   []*Vote
	precommits []*Vote
}

type VoteSetMap struct {
	prevotes   []bool
	precommits []bool
}

// find out which vote does not sync yet
func subVotes(v1 []*Vote, v2 []bool) *Vote {
	for _, v1_ := range v1 {
		if v1_ != nil && !v2[v1_.ValidatorId] {
			return v1_
		}
	}
	return nil
}

func (s *VoteSet) AddVote(v *Vote) {
	if v.Type == VoteTypePrevote {
		s.prevotes[v.ValidatorId] = v
	} else {
		s.precommits[v.ValidatorId] = v
	}
}
