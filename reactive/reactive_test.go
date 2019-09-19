package reactive

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestComputeCell(t *testing.T) {
	t.Log("Testing ComputeCell functionality")
	{
		i1 := NewInputCell()
		i2 := NewInputCell()
		c1 := NewComputeCell()
		c1.SetDependency(i1, i2)
		c1.ComputeFn = func(inputs ...*InputCell) int {
			return inputs[0].Get() + inputs[1].Get()
		}
		i1.Set(2)
		i2.Set(3)

		if c1.Value == 5 {
			t.Log("\tCalculate value from InputCells", checkMark)
		} else {
			t.Fatal("\tCalculate value from InputCells", ballotX)
		}

		i1.Set(3)

		if c1.Value == 6 {
			t.Log("\tUpdate value when InputCells are updated", checkMark)
		} else {
			t.Fatal("\tUpdate value when InputCells are updated", ballotX)
		}
	}

}
