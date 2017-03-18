package paxos

type MsgCounter struct {
}

func (m *MsgCounter) AddReject(nodeid uint64) {

}

func (m *MsgCounter) AddReceive(nodeid uint64) {

}

func (m *MsgCounter) AddPromiseOrAccept(nodeid uint64) {

}

func (m *MsgCounter) IsPassedOnThisRound() {

}

func (m *MsgCounter) IsRejectedOnThisRound() {

}

func (m *MsgCounter) IsAllReceiveOnThisRound() {

}

func (m *MsgCounter) StartNewRound() {

}
