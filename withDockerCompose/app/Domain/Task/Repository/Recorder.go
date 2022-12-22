package Repository

import "prove1/Domain/Task"

type Recorder interface {
	Persist(task  Task.Task)
}
