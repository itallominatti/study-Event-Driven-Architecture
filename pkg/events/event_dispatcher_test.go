package events

import "time"

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GateDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct{}

func (h *TestEventHandler) Handle(event EventInterface) {

}
