package ethereum

type EventHandle struct {
	//go-ethereum lib instance
}

func NewEventHandle() *EventHandle {
	ret := &EventHandle{}
	return ret
}

func (e *EventHandle) Do() {
	//step1: consume event

	//step2: parse log

	//step3: dispatch evnet
}

func (e *EventHandle) UpdateMintStatus() {

}

func (e *EventHandle) UpdateTransStatus() {

}
