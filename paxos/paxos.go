package paxos

import (
	"github.com/masterlvng/myla/node"
	"github.com/masterlvng/myla/proto/paxosmsg"
)

//Ballot ...
type Ballot struct {
	ProposalID uint64
	NodeID     uint64
}

//Eq ...
func (b *Ballot) Eq(other Ballot) bool {
	if b.ProposalID == other.ProposalID && b.NodeID == other.ProposalID {
		return true
	}
	return false
}

//NoLT ...
func (b *Ballot) NoLT(other Ballot) bool {
	if b.ProposalID == other.ProposalID {
		return b.NodeID >= other.NodeID
	}
	return b.ProposalID >= other.ProposalID
}

//IsNULL ..
func (b *Ballot) IsNULL() bool {
	return b.ProposalID == 0
}

//Reset ...
func (b *Ballot) Reset() {
	b.ProposalID = 0
	b.NodeID = 0
}

type base struct {
	InstanceID uint64
}

func (b *base) NewInstance() {

}

func (b *base) SendMsg(nid node.ID) {

}

func (b *base) BroadcastMsg(msg *paxosmsg.PaxosMsg) {

}
