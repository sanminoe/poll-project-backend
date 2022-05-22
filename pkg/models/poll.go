package pollModel

import (
	"time"

	"github.com/sanminoe/poll-project-backend/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Option struct {
	Value string
	votes int
}

type Poll struct {
	gorm.Model
	Title  string
	votes  int
	Closed bool

	CloseBy *time.Time
}

func init() {
	config.Connect()

	db = config.GetDB()
	db.AutoMigrate(&Poll{})
}

func (p *Poll) CreatePoll() *Poll {
	db.Create(&p)

	return p
}
