package jobs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type JobStatus string

var (
	created  JobStatus = "created"
	started  JobStatus = "started"
	finished JobStatus = "finished"
)

type StreamMsg struct {
	jobID int
	msg   string
}

type Job struct {
	id            int
	name          string
	status        JobStatus
	duration      time.Duration
	writeInterval time.Duration
	logStream     chan StreamMsg
}

func NewJob(id int, name string, duration time.Duration, logInterval time.Duration) *Job {
	return &Job{
		id:            id,
		name:          name,
		status:        created,
		duration:      duration,
		writeInterval: logInterval,
		logStream:     make(chan StreamMsg),
	}
}

func (j *Job) GetLogStream() chan StreamMsg {
	return j.logStream
}

func (j *Job) GetID() int {
	return j.id
}

func (j *Job) RunInBackground() {
	go j.Run()
}

func (j *Job) Run() {
	writeIntervalTicker := time.NewTicker(j.writeInterval).C
	durationTicker := time.NewTicker(j.duration).C
	j.status = started

	filename := "logs/logs.log"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Error("failed to open log file", logrus.WithError(err))
		return
	}
	defer func() { _ = file.Close() }()
	defer func() { j.status = finished }()

	for {
		select {
		case <-writeIntervalTicker:
			logLine := fmt.Sprintf("Job %d: %s, Timestamp: %v \n", j.id, j.name, time.Now().UTC())

			_, err = file.WriteString(logLine)
			if err != nil {
				logrus.Error("failed to write data", logrus.WithError(err))
				return
			}

			j.logStream <- StreamMsg{jobID: j.id, msg: logLine}

		case <-durationTicker:
			close(j.logStream)
			return
		}
	}

}
