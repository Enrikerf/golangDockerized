package CreateTask

type Command struct {
	Host              string
	Port              string
	Sentences         []string
	CommunicationMode string
	Status            string
	ExecutionMode     string
}
