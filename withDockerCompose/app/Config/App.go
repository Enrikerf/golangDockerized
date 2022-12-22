package Config

import (
	"gorm.io/gorm"
	"prove1/Adapter/in/SdkFacade"
	"prove1/Adapter/in/SdkFacade/TaskSdk"
	"prove1/Adapter/out/Persistence/Mysql/TaskAdapters"
	"prove1/Application/Port/In/Task/CreateTask"
)

type App interface {
	GetSdkFacade() SdkFacade.Facade
}

func NewApp() App {
	var db *gorm.DB
	recorder := TaskAdapters.NewRecorder(db)
	var createTaskService = CreateTask.New(recorder)
	taskSdk := TaskSdk.NewSdk(createTaskService)
	facade := SdkFacade.Facade{TaskSdk: taskSdk}
	app := &app{facade}
	return app
}

type app struct {
	SdkFacade SdkFacade.Facade
}

func (app *app) GetSdkFacade() SdkFacade.Facade {
	return app.SdkFacade
}
