package domain

type EventType int8

const (
	ClientMeeting EventType = iota + 1
	Show
	ScheduledCall
)

type Event struct {
	Id        int       `json:"id"`
	Datetime  string    `json:"datetime"`
	Duration  string    `json:"duration"`
	EventType EventType `json:"eventType"`
	Comment   *string   `json:"comment,omitempty"`
}
