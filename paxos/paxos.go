package paxos

type Ballot struct {
	ProposalID uint64
	NodeID     uint64
}

func (b *Ballot) Eq(other Ballot) bool {
	if b.ProposalID == other.ProposalID && b.NodeID == other.ProposalID {
		return true
	}
	return false
}

func (b *Ballot) NoLT(other Ballot) bool {
	if b.ProposalID == other.ProposalID {
		return b.NodeID >= other.NodeID
	}
	return b.ProposalID >= other.ProposalID
}

func (b *Ballot) IsNULL() bool {
	return b.ProposalID == 0
}

func (b *Ballot) Reset() {
	b.ProposalID = 0
	b.NodeID = 0
}

type base struct {
	InstanceID uint64
}

func (b *base) NewInstance() {

}

func (b *base) SendMsg() {

}

func (b *base) BroadcastMsg() {

}
