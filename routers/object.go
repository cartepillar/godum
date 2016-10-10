// object
package routers

import (
	"servertest/controllers"

	"github.com/gorilla/mux"
)

func SetObjectRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.ServerTest).Methods("POST")
	return router
}
