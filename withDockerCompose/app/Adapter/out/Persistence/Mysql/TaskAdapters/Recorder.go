package TaskAdapters

import (
	"fmt"
	"gorm.io/gorm"
	"prove1/Domain/Task"
)

type Recorder interface {
	Persist(task Task.Task)
}

type recorder struct {
	Orm *gorm.DB
}

func (r *recorder) Persist(task Task.Task) {
	fmt.Println("hasta abajo")
}

func NewRecorder(orm *gorm.DB) Recorder {
	return &recorder{orm}
}
