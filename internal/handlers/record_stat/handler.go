package recordstat

import (
	"context"
	"encoding/json"
	"net/http"

	"example.com/interview/internal/stats"
)

type Handler struct {
	storage stats.Storage
	logger  Logger
}

func New(s stats.Storage, logger Logger) *Handler {
	return &Handler{s, logger}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)

	var dto Payload
	err := dec.Decode(&dto)
	if err != nil {
		h.logger.Error("failed to parse user input", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.storage.Add(context.TODO(), dto.UserId, dto.AuthorId)
	if err != nil {
		h.logger.Error("failed to update user stats", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

var _ http.Handler = &Handler{}
