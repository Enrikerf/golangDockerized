package Task

import (
	"prove1/Domain/Core"
	"prove1/Domain/Task/CommunicationMode"
	"prove1/Domain/Task/ExecutionMode"
	"prove1/Domain/Task/Status"
	"prove1/Domain/Task/Step"
)

type Task interface {
	GetStatus() Status.Status
	GetSteps() []Step.Step
}

type task struct {
	uuid              Core.Id
	host              string
	port              string
	steps             []Step.Step
	communicationMode CommunicationMode.Mode
	status            Status.Status
	executionMode     ExecutionMode.Mode
}

func (t *task) GetStatus() Status.Status {
	return t.status
}

func (t *task) GetSteps() []Step.Step {
	return t.steps
}

func NewTask(
	host string,
	port string,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode,
) Task {
	task := &task{}
	task.uuid = Core.NewId()
	task.host = host
	task.port = port
	task.communicationMode = communicationMode
	task.executionMode = executionMode
	task.status = Status.Pending
	for _, stepVo := range stepVos {
		step := stepVo.CreateStep(task.uuid)
		task.steps = append(task.steps, step)
	}
	return task
}
