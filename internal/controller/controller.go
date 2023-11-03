package controller

import "board_action/internal/service"

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}
