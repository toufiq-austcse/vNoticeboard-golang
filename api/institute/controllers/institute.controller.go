package controllers

import (
	"gihub.com/toufiq-austcse/vNoticeboard/api/institute/services"
)


type  InstituteController struct {
	service services.InstituteService
}

func NewInstituteController(service services.InstituteService) InstituteController {
	return InstituteController{
		service: service,
	}
}

