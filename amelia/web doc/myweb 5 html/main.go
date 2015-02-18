package main

import (
	"github.com/amelia/gomvc"
	"myweb/controllers"
)

func main() {
	gomvc.RouteFolder("/", "public")
	gomvc.Route("/proses/", &controllers.ProsesController{})
	gomvc.Run()
}
