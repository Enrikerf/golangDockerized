package TaskAdapters

import (
	"gorm.io/gorm"
	"prove1/Domain/Core"
	"prove1/Domain/Task"
	"prove1/Domain/Task/CommunicationMode"
	"prove1/Domain/Task/ExecutionMode"
	"prove1/Domain/Task/Step"
)

type Finder interface {
}

type finder struct {
	Orm *gorm.DB
}

func (r *finder) Find(id Core.Id) Task.Task {
	var stepVos []Step.Vo
	stepVo, _ := Step.NewStepVo("")
	stepVos = append(stepVos, stepVo)
	task := Task.NewTask("host", "Port", stepVos, CommunicationMode.Unary, ExecutionMode.Manual)
	return task
}

func NewFinder() Finder {
	return &finder{}
}
