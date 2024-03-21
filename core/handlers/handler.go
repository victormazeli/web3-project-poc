package handlers

import (
	"goApiStartetProject/internal/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		Service: srv,
	}
}
