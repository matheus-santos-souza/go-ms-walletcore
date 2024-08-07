package events

<<<<<<< HEAD
import "time"
=======
import (
	"sync"
	"time"
)
>>>>>>> master

type IEvent interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

type IEventHandler interface {
<<<<<<< HEAD
	Handle(event IEvent)
=======
	Handle(event IEvent, wg *sync.WaitGroup)
>>>>>>> master
}

type IEventDispatcher interface {
	Register(eventName string, handler IEventHandler) error
	Dispatch(event IEvent) error
	Remove(eventName string, handler IEventHandler) error
	Has(eventName string, handler IEventHandler) bool
	Clear() error
}
