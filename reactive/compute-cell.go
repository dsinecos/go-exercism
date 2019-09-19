package reactive

// ComputeCell - TODO
type ComputeCell struct {
	Value      int
	listeners  map[string][]func(v int)
	dependency []*InputCell
	ComputeFn  func(inputs ...*InputCell) int
}

// NewComputeCell - Create a new ComputeCell
func NewComputeCell() *ComputeCell {
	i := ComputeCell{
		dependency: []*InputCell{},
		listeners:  make(map[string][]func(v int)),
	}
	return &i
}

// On - Attach a listener for an event
func (c *ComputeCell) On(e string, f func(v int)) {
	if _, ok := c.listeners[e]; !ok {
		c.listeners[e] = []func(v int){}
	}
	c.listeners[e] = append(c.listeners[e], f)
}

// SetDependency - Add InputCells which the ComputeCell depends upon
func (c *ComputeCell) SetDependency(inputs ...*InputCell) {
	for _, v := range inputs {
		v.On(UPDATED, c.reCompute)
		c.dependency = append(c.dependency, v)
	}
}

func (c *ComputeCell) emit(e string) {
	if _, ok := c.listeners[e]; ok {
		for _, v := range c.listeners[e] {
			v(c.Value)
		}
	}
}

func (c *ComputeCell) reCompute(v int) {
	c.Value = c.ComputeFn(c.dependency...)
	c.emit(UPDATED)
}
