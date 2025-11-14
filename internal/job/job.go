package job

// define the job struct here

type Job struct {
	JobID              string
	JobType            string
	Payload            map[string]interface{}
	MaxRetry           int8
	CurentAttemptCount int8
	EnqueuedAt         int64
	ScheduledAt        int64
}
