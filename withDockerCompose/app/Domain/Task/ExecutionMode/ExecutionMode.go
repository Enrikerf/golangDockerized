package ExecutionMode

type Mode string

const (
	Automatic Mode = "AUTOMATIC"
	Manual    Mode = "MANUAL"
)

func CreateFromString(modeString string) (Mode, error) {
	switch modeString {
	case string(Automatic):
		return Automatic, nil
	case string(Manual):
		return Manual, nil
	}
	return "", &UnknownError{modeString: modeString}
}
