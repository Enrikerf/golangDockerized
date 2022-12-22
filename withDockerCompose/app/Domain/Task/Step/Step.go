package Step

import "prove1/Domain/Core"

type Step interface {
	GetUuid() Core.Id
}
type step struct {
	Uuid     Core.Id
	TaskUuid Core.Id
	Sentence string
}

func (step *step) GetUuid() Core.Id {
	return step.Uuid
}

func NewStep(TaskId Core.Id, stepVo Vo) Step {
	self := &step{}
	self.Uuid = Core.NewId()
	self.TaskUuid = TaskId
	self.Sentence = stepVo.GetSentence()
	return self
}
