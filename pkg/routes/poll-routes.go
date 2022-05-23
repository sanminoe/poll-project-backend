package pollRoutes

import (
	"github.com/gorilla/mux"
	pollControllers "github.com/sanminoe/poll-project-backend/pkg/controllers"
)

var RegisterPollRoutes = func(router *mux.Router) {

	router.HandleFunc("/poll/", pollControllers.GetAllPolls).Methods("GET")
	router.HandleFunc("/poll/{pollId}/", pollControllers.GetPollById).Methods("GET")

	router.HandleFunc("/poll/{pollId}/", pollControllers.DeletePoll).Methods("DELETE")

	router.HandleFunc("/poll/", pollControllers.CreatePoll).Methods("POST")
}
