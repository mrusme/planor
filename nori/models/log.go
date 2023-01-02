package models

import (
	"fmt"
	"time"
)

type LogEvent struct {
	Message    string
	Timestamp  time.Time
	IngestedAt time.Time
}

type LogStream struct {
	ID             string
	Name           string
	GroupName      string
	CreatedAt      time.Time
	FirstEventAt   time.Time
	LastEventAt    time.Time
	LastIngestedAt time.Time
	LogEvents      []LogEvent
}

func (logStream LogStream) FilterValue() string {
	return logStream.Name
}

func (logStream LogStream) Title() string {
	return logStream.Name
}

func (logStream LogStream) Description() string {
	return logStream.LastEventAt.String()
}

type LogGroup struct {
	ID        string
	Name      string
	Streams   []LogStream
	SizeBytes int64
}

func (logGroup LogGroup) FilterValue() string {
	return logGroup.Name
}

func (logGroup LogGroup) Title() string {
	return logGroup.Name
}

func (logGroup LogGroup) Description() string {
	return fmt.Sprintf("%d KB", logGroup.SizeBytes/1024)
}
