package events

type KwArgs map[string]interface{}

type Event struct {
	Name   string
	KwArgs KwArgs
	Error  error
}

func New(name string, kwargs KwArgs) *Event {
	return &Event{
		Name:   name,
		KwArgs: kwargs,
	}
}
