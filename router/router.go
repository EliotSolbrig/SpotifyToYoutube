package router

import (
	// "context"
	// "main/service"
)

var basePasefiles []string = []string{
    "templates/components/header.html",
    "templates/components/footer.html",
    "templates/base.html",
}

type Router struct {
    // service service.IService
}

// func NewRouter(service service.IService) *Router {
func NewRouter() *Router {

    return &Router{
        // service: service,
    }
}
