package wow

type Wow interface {
	Wow()
	wow()
}

type wow struct {
	w0w string
}

func (wow *wow) wow() string {

	return "wow"
}

func (wow *wow) Wow() string {
	return "wow"
}

// wow
func NewWow() *wow {
	return &wow{
		w0w: "wow",
	}
}
