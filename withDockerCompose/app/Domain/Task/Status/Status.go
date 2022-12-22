package Status

type Status string

const (
	Pending Status = "PENDING"
	Running Status = "RUNNING"
	Done    Status = "DONE"
)

func CreateFromString(modeString string) (Status, error) {
	switch modeString {
	case string(Pending):
		return Pending, nil
	case string(Running):
		return Running, nil
	case string(Done):
		return Done, nil
	}
	return "", &UnknownError{modeString: modeString}
}
