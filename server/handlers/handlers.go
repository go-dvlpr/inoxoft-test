package handlers

import "inoxoft-test/jobs"

type Handlers struct {
	jobProcessor jobs.Jobber
}

func New(jobProcessor jobs.Jobber) *Handlers {
	return &Handlers{jobProcessor: jobProcessor}
}
