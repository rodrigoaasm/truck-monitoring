package pubsub

type EventPublisherPayload interface {
}

type EventPublisherInterface interface {
	SendEvent(payload EventPublisherPayload) error
}
