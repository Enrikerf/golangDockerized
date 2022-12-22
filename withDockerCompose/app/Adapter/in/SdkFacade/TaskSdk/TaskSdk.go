package TaskSdk

import (
	"fmt"
	"prove1/Application/Port/In/Task/CreateTask"
)

type Sdk interface {
	CreateTask(command CreateTask.Command)
}

type sdk struct {
	createTaskUseCase CreateTask.UseCase
}

func (c *sdk) CreateTask(command CreateTask.Command) {
	_, err := c.createTaskUseCase.Create(command)
	if err != nil {
		fmt.Println("error on sdk")
		return
	}
}

func NewSdk(createTaskUseCase CreateTask.UseCase) Sdk {
	sdk := &sdk{createTaskUseCase}
	return sdk
}
