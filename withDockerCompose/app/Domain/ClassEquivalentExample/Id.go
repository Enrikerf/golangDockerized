package ClassEquivalentExample

type Class interface {
	Get() string
	Set(newValue string)
}

type class struct {
	variable string
}

func NewClass() Class {
	return &class{"default"}
}

func (class *class) Get() string {
	return class.variable
}

func (class *class) Set(newValue string) {
	class.variable = newValue
}
