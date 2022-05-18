package webhook

type Event interface {
	GetMessage() string
}

type Payload struct {
	Event
}

func NewPayload(event Event) *Payload {
	return &Payload{event}
}
