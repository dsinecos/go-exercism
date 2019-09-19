package reactive

// InputCell - TODO
type InputCell struct {
	value     int
	listeners map[string][]func(v int)
}

// NewInputCell - TODO
func NewInputCell() *InputCell {
	i := InputCell{
		listeners: make(map[string][]func(v int)),
	}
	return &i
}

// Set - TODO
func (i *InputCell) Set(v int) {
	i.value = v
	i.emit(UPDATED)
}

// Get - TODO
func (i *InputCell) Get() int {
	return i.value
}

// On - TODO
func (i *InputCell) On(e string, c func(v int)) {
	if _, ok := i.listeners[e]; !ok {
		// i.listeners[e] = make([]func(v int), 1)
		i.listeners[e] = []func(v int){}
	}
	i.listeners[e] = append(i.listeners[e], c)
}

func (i *InputCell) emit(e string) {
	if _, ok := i.listeners[e]; ok {
		for _, v := range i.listeners[e] {
			v(i.value)
		}
	}
}
