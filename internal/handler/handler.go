package handler

import "net"

type Handler struct {
	Conn *net.Conn
}

func NewHandler(Conn *net.Conn) *Handler {
	return &Handler{Conn: Conn}
}

func (h *Handler) handle() {
}
