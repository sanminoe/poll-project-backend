package pollControllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	pollModel "github.com/sanminoe/poll-project-backend/pkg/models"
	"github.com/sanminoe/poll-project-backend/pkg/utils"
)

var NewPoll pollModel.Poll

func CreatePoll(w http.ResponseWriter, r *http.Request) {

	CreatePoll := &pollModel.Poll{}
	utils.ParseBody(r, CreatePoll)

	p := CreatePoll.CreatePoll()
	res, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAllPolls(w http.ResponseWriter, r *http.Request) {
	polls := pollModel.GetAllPolls()

	res, _ := json.Marshal(polls)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPollById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pollId := vars["pollId"]
	ID, err := strconv.Atoi(pollId)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	poll, _ := pollModel.GetPollById(ID)

	res, _ := json.Marshal(poll)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePoll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pollId := vars["pollId"]

	ID, parseErr := strconv.Atoi(pollId)

	if parseErr != nil {
		fmt.Println("Error while parsing")
	}
	pollMod := pollModel.Poll{}

	pollErr := pollMod.DeletePoll(ID)

	if pollErr != nil {
		fmt.Println("Error while deleting!")
	}

	res, _ := json.Marshal("ok")

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
