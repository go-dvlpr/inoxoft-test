package handlers

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (h *Handlers) StreamAllLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	filename := "logs/logs.log"

	file, err := os.Open(filename)
	if err != nil {
		logrus.Warn("failed to open log file", logrus.WithError(err))
	}
	defer file.Close()

	var existedLogs string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		existedLogs += scanner.Text() + "\n"
	}

	_, _ = w.Write([]byte(existedLogs))
	flusher.Flush()

	stream := h.jobProcessor.SubscribeToStream(-1)

	go func() {
		<-r.Context().Done()
		close(stream)
	}()

	for log := range stream {
		w.Write([]byte(log))
		flusher.Flush()
	}

}

func (h *Handlers) StreamLogs(w http.ResponseWriter, r *http.Request) {
	jobIDStr := chi.URLParam(r, "id")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		http.Error(w, "id should be int", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	filename := "logs/logs.log"

	file, err := os.Open(filename)
	if err != nil {
		logrus.Warn("failed to open log file", logrus.WithError(err))
	}
	defer file.Close()

	var existedLogs string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, fmt.Sprintf("Job %d:", jobID)) { // Фильтруем по условию
			existedLogs += scanner.Text() + "\n"
		}
	}

	_, _ = w.Write([]byte(existedLogs))
	flusher.Flush()

	stream := h.jobProcessor.SubscribeToStream(jobID)

	go func() {
		<-r.Context().Done()
		close(stream)
	}()

	for log := range stream {
		w.Write([]byte(log))
		flusher.Flush()
	}

}
