package message

// EventHeaderKey represents the header key for the kafka messages.
type EventHeaderKey string

const (
	// HdrTopic is the type of event that match finally with the schema.
	HdrTopic EventHeaderKey = "Topic"
)
