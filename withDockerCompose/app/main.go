package main

import (
	"prove1/Application/Port/In/Task/CreateTask"
	"prove1/Config"
)

func main() {
	app := Config.NewApp()
	app.GetSdkFacade().TaskSdk.CreateTask(CreateTask.Command{
		Host:              "",
		Port:              "",
		Sentences:         nil,
		CommunicationMode: "UNARY",
		Status:            "PENDING",
		ExecutionMode:     "MANUAL",
	})
}

/*func homeHandler(w http.ResponseWriter, r *http.Request) {
	myOs, myArch := runtime.GOOS, runtime.GOARCH
	inContainer := "inside"
	_, _ = fmt.Fprintf(w, "hello")
	_, _ = fmt.Fprintf(w, "hello6")
	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		inContainer = "outside"
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Hello, %s\n", r.UserAgent())
	_, _ = fmt.Fprintf(w, "%s,%s\n", myOs, myArch)
	_, _ = fmt.Fprintf(w, "%s\n", inContainer)
	_, _ = fmt.Fprintf(w, "%s\n", runtime.Version())
}

func main() {
	http.HandleFunc("/d", homeHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("hello %s\n", runtime.Version())
	}
}
*/
