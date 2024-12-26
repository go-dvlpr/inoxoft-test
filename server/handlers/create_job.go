package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CreateJobRequest struct {
	Name                string `json:"name"`
	MillisecondDuration int    `json:"millisecond_duration"`
}

type CreateJobResponse struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	MillisecondDuration int    `json:"millisecond_duration"`
}

func (h *Handlers) CreateJob(w http.ResponseWriter, r *http.Request) {
	var req CreateJobRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.Error("failed to parse req body", logrus.WithError(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if req.MillisecondDuration < 30000 {
		logrus.Error("invalid duration")
		http.Error(w, "duration should be at least 30000 millisecond (30s)", http.StatusBadRequest)
		return
	}

	job := h.jobProcessor.NewJob(req.Name, time.Millisecond*time.Duration(req.MillisecondDuration), time.Second*10)
	err = h.jobProcessor.AddJob(job)
	if err != nil {
		logrus.Error("failed to add new job", logrus.WithError(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := CreateJobResponse{
		ID:                  job.GetID(),
		Name:                req.Name,
		MillisecondDuration: req.MillisecondDuration,
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		logrus.Error("failed to marshal response", logrus.WithError(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(respBytes)
	return
}
