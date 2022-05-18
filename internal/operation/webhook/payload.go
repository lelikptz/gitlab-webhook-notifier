package webhook

type Event interface {
	GetMessage() string
	IsSuccessful() bool
}

type Payload struct {
	Event
}

func NewPayload(event Event) *Payload {
	return &Payload{event}
}
