package worker

import jobtype "github.com/codingbot24-s/distributed-job-system/internal/job"


type Processor interface {
	Process(job *jobtype.Job) error
}