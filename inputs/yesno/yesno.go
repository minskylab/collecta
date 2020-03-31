package yesno

import (
	"github.com/minskylab/collecta"
)

// YesNo represent an YesNo Runtime type of any question
type YesNo struct {
	value        bool
	defaultValue bool
}

// New creates a new YesNo modified with options
func New(mods ...Mod) (*YesNo, error) {
	yesNo := new(YesNo)
	for _, option := range mods {
		if err := option(yesNo); err != nil {
			return nil, err
		}
	}
	return yesNo, nil
}

// Type returns the type of text, in this case: "text"
func (yesNo *YesNo) Type() collecta.InputType {
	return collecta.YesNo
}

func (yesNo *YesNo) Value() ([]string, error) {
	if yesNo.value {
		return []string{"yes"}, nil
	}
	return []string{"no"}, nil
}
