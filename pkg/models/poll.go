package pollModel

import (
	"time"

	"github.com/sanminoe/poll-project-backend/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Option struct {
	gorm.Model
	PollID uint
	Value  string
	votes  int
}

type Poll struct {
	gorm.Model
	Title   string
	votes   int
	Closed  bool
	Options []Option `gorm:"foreignKey:PollID"`
	CloseBy *time.Time
}

func init() {
	config.Connect()

	db = config.GetDB()
	db.AutoMigrate(&Poll{}, &Option{})
}

func (p *Poll) CreatePoll() *Poll {
	db.Create(&p)

	db.Save(&p)

	return p
}
func (p *Poll) DeletePoll(pollId int) (err error) {
	p.ID = uint(pollId)
	db.Delete(&p).Where("ID=?", pollId)

	return
}

func GetAllPolls() []Poll {
	var polls []Poll

	db.Preload("Options").Find(&polls)

	return polls
}

func GetPollById(pollId int) (*Poll, *gorm.DB) {
	var poll Poll

	db := db.Preload("Options").Find(&poll, pollId)

	return &poll, db
}
