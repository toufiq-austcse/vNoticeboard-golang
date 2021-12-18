package controllers

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/services"
)

type InstituteController struct {
	service services.InstituteService
}

func NewInstituteController(service services.InstituteService) InstituteController {
	return InstituteController{
		service: service,
	}
}
