package events

import "time"

type EventInterface interface {
	GetName() string
	GateDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(evantName string, handler EventHandlerInterface) error
	Has(evantName string, handler EventHandlerInterface) bool
	Clear() error
}
