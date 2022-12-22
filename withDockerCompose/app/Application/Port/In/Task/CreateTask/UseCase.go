package CreateTask

import (
	"fmt"
	"prove1/Domain/Task"
	"prove1/Domain/Task/CommunicationMode"
	"prove1/Domain/Task/ExecutionMode"
	"prove1/Domain/Task/Repository"
	"prove1/Domain/Task/Service"
	"prove1/Domain/Task/Step"
)

type UseCase interface {
	Create(task Command) (Task.Task, error)
}

func New(recorder Repository.Recorder) UseCase {
	return &useCase{Service.TaskCreator{Recorder: recorder}}
}

type useCase struct {
	TaskCreator Service.TaskCreator
}

func (useCase useCase) Create(command Command) (Task.Task, error) {
	var err error
	var stepVos []Step.Vo
	for _, commandSentence := range command.Sentences {
		stepVo, err := Step.NewStepVo(commandSentence)
		if err != nil {
			fmt.Println("error on use case 1")
			return nil, err
		}
		stepVos = append(stepVos, stepVo)
	}
	communicationMode, err := CommunicationMode.CreateFromString(command.CommunicationMode)
	if err != nil {
		fmt.Println("error on use case 2")
		return nil, err
	}
	executionMode, err := ExecutionMode.CreateFromString(command.ExecutionMode)
	if err != nil {
		fmt.Println("error on use case 3")
		return nil, err
	}
	task, nil := useCase.TaskCreator.Create(command.Host,
		command.Port,
		stepVos,
		communicationMode,
		executionMode)

	return task, nil
}
