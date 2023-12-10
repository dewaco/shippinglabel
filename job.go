package shippinglabel

import "time"

type JobStatusCode string

const (
	JobStatusCreated   JobStatusCode = "CREATED"
	JobStatusRunning   JobStatusCode = "RUNNING"
	JobStatusCancelled JobStatusCode = "CANCELLED"
	JobStatusCompleted JobStatusCode = "COMPLETED"
)

type ShipmentJob struct {
	ID                  int                  `json:"id,omitempty"`
	Status              JobStatusCode        `json:"status,omitempty"`
	ExecutionTime       *time.Time           `json:"executionTime,omitempty"`
	LastUpdate          time.Time            `json:"lastUpdate,omitempty"`
	Created             time.Time            `json:"created,omitempty"`
	QueueItems          []*ShipmentQueueItem `json:"queueItems,omitempty"`
	TotalQueueItems     int                  `json:"totalQueueItems,omitempty"`
	ProcessedQueueItems int                  `json:"processedQueueItems,omitempty"`
}
