package controllers

import (
	"github.com/amelia/gomvc"
	"net/http"
)

type ProsesController struct {
	gomvc.Controller
}

func (c *ProsesController) hasil(rw http.ResponseWriter, req *http.Request) {
	c.ServeText("text")
}
