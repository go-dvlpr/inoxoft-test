package jobs

import (
	"errors"
	"sync"
	"time"
)

type Jobber interface {
	GetNextJobID() int
	NewJob(name string, duration time.Duration, logInterval time.Duration) *Job
	AddJob(job *Job) error
	SubscribeToStream(jobID int) chan string
}

type ClientStream struct {
	stream chan string
	jobID  int
}

type JobProcessor struct {
	nextJobID      int
	jobs           map[int]*Job
	mu             sync.Mutex
	clientsStreams []ClientStream
}

func NewJobProcessor() Jobber {
	return &JobProcessor{
		nextJobID: 0,
		jobs:      make(map[int]*Job),
		mu:        sync.Mutex{},
	}
}

func (s *JobProcessor) NewJob(name string, duration time.Duration, logInterval time.Duration) *Job {
	s.mu.Lock()
	jobID := s.nextJobID
	s.nextJobID++
	s.mu.Unlock()

	return NewJob(jobID, name, duration, logInterval)
}

func (s *JobProcessor) GetNextJobID() int {
	return s.nextJobID
}

func (s *JobProcessor) AddJob(job *Job) error {
	s.mu.Lock()
	_, ok := s.jobs[job.id]
	if ok {
		return errors.New("job with specific id already exists")
	}

	s.jobs[job.id] = job
	s.mu.Unlock()

	job.RunInBackground()
	s.AddJobToStream(job.logStream)

	return nil
}

func (s *JobProcessor) SubscribeToStream(jobID int) chan string {
	stream := make(chan string)
	s.clientsStreams = append(s.clientsStreams, ClientStream{
		stream: stream,
		jobID:  jobID,
	})

	return stream
}

func (s *JobProcessor) AddJobToStream(jobStream chan StreamMsg) {
	go func() {
		for log := range jobStream {
			for i, cs := range s.clientsStreams {

				go func(cs ClientStream, i int) {
					defer func(i int) {
						if r := recover(); r != nil {
							s.mu.Lock()
							s.RemoveClientStream(i)
							s.mu.Unlock()
						}
					}(i)

					if cs.jobID == log.jobID || cs.jobID == -1 {
						cs.stream <- log.msg
					}

				}(cs, i)

			}
		}
	}()
}

func (s *JobProcessor) RemoveClientStream(index int) {
	s.clientsStreams = append(s.clientsStreams[:index], s.clientsStreams[index+1:]...)
}
