package pollControllers

import (
	"encoding/json"
	"net/http"

	pollModel "github.com/sanminoe/poll-project-backend/pkg/models"
	"github.com/sanminoe/poll-project-backend/pkg/utils"
)

var NewPoll pollModel.Poll

func CreatePoll(w http.ResponseWriter, r *http.Request) {
	CreatePoll := &pollModel.Poll{}
	utils.ParseBody(r, CreatePoll)

	p := CreatePoll.CreatePoll()
	res, _ := json.Marshal(p)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
