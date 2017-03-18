package paxos

type Proposer struct {
	ProposerID                  uint64
	HighestOtherProposerID      uint64
	Proposal                    []byte
	HighestOtherPreAcceptBallot Ballot
}

func (p *Proposer) AddPreAcceptValue(otherPreAcceptBallot *Ballot, otherPreAcceptProposal []byte) {

}

func (p *Proposer) ResetHighestOtherPreAcceptBallot() {
	p.HighestOtherPreAcceptBallot.Reset()
}

func (p *Proposer) Prepare() {

}

func (p *Proposer) OnPrepareReply() {
}

func (p *Proposer) Accept() {

}

func (p *Proposer) OnAcceptReply() {

}

func (p *Proposer) ExitPrepare() {

}

func (p *Proposer) ExitAccept() {

}

func (p *Proposer) CancelSkipPrepare() {

}
