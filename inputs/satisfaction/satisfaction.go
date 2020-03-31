package satisfaction

import (
	"fmt"

	"github.com/minskylab/collecta"
)

type FeelingOption string

// Satisfaction represent an Satisfaction Runtime type of any question
type Satisfaction struct {
	value        float64
	defaultValue float64
	totalOptions int
	options      map[int]FeelingOption
}

// New creates a new Satisfaction modified with options
func New(mods ...Mod) (*Satisfaction, error) {
	satisfaction := new(Satisfaction)
	satisfaction.totalOptions = 3                 // default
	satisfaction.options = map[int]FeelingOption{ // default
		0: "😡",
		1: "😐",
		2: "😊",
	}
	for _, option := range mods {
		if err := option(satisfaction); err != nil {
			return nil, err
		}
	}
	return satisfaction, nil
}

// Type returns the type of text, in this case: "text"
func (satisfaction *Satisfaction) Type() collecta.InputType {
	return collecta.Satisfaction
}

func (satisfaction *Satisfaction) Value() ([]string, error) {
	val := fmt.Sprintf("%.2f", satisfaction.value)
	return []string{val}, nil
}
