package pollRoutes

import (
	"github.com/gorilla/mux"
	pollControllers "github.com/sanminoe/poll-project-backend/pkg/controllers"
)

var RegisterPollRoutes = func(router *mux.Router) {

	router.HandleFunc("/poll/", pollControllers.CreatePoll).Methods("GET")
}
