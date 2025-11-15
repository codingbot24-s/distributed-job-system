package jobtype

// define the job struct here

type Job struct {
	JobID               string                 `json:"jobId,omitempty"`
	JobType             string                 `json:"jobType" binding:"required"`
	Payload             map[string]interface{} `json:"payload,omitempty"`
	MaxRetry            int8                   `json:"maxRetry,omitempty"`
	CurrentAttemptCount int8                   `json:"currentAttemptCount,omitempty"`
	EnqueuedAt          int64
	ScheduledAt         int64	
}
