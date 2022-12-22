package Repository

import (
	"prove1/Domain/Core"
	"prove1/Domain/Task"
)

type Finder interface {
	Find(id Core.Id) Task.Task
}
