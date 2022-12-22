package Service

import (
	"fmt"
	"prove1/Domain/Task"
	"prove1/Domain/Task/CommunicationMode"
	"prove1/Domain/Task/ExecutionMode"
	"prove1/Domain/Task/Repository"
	"prove1/Domain/Task/Step"
)

type TaskCreator struct {
	Recorder Repository.Recorder
}

func (taskCreator *TaskCreator) Create(
	host string,
	port string,
	stepVos []Step.Vo,
	communicationMode CommunicationMode.Mode,
	executionMode ExecutionMode.Mode,
) (Task.Task, error) {

	var task = Task.NewTask(
		host,
		port,
		stepVos,
		communicationMode,
		executionMode,
	)
	fmt.Println("TaskCreatedEvent")
	taskCreator.Recorder.Persist(task)
	return task, nil

}
