package handlers

import (
	"os"

	"github.com/Denvos/core/errors"
	"github.com/Denvos/core/logger"
)

type Handler interface {
	Handle(err *errors.Error)
}

type LoggerHandler struct {
	log *logger.Logger
}

func NewLoggerHandler(log *logger.Logger) *LoggerHandler {
	return &LoggerHandler{log: log}
}

func (h *LoggerHandler) Handle(err *errors.Error) {
	if h.log == nil {
		return
	}
	h.log.Error(err.Message(), "code", err.Code(), "cause", err.Cause())
}

type ExitHandler struct{}

func (h *ExitHandler) Handle(err *errors.Error) {
	if err != nil {
		os.Exit(1)
	}
}
