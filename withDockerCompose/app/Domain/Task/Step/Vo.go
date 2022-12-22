package Step

import (
	"errors"
	"prove1/Domain/Core"
)

type Vo interface {
	GetSentence() string
	CreateStep(taskId Core.Id) Step
}

type vo struct {
	Sentence string
}

func (vo *vo) GetSentence() string {
	return vo.Sentence
}

func (vo *vo) CreateStep(taskId Core.Id) Step {

	newStep := &step{}
	newStep.Uuid = Core.NewId()
	newStep.TaskUuid = taskId
	newStep.Sentence = vo.GetSentence()
	return newStep

}

func NewStepVo(sentence string) (Vo, error) {
	self := &vo{}
	if len(sentence) > 255 {
		return self, errors.New("step.sentence length must be less than 255")
	}
	self.Sentence = sentence
	return self, nil
}
