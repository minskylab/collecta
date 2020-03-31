package yesno

type Mod func(yesNo *YesNo) error

func WithDefaultValue(defaultValue bool) Mod {
	return func(yesNo *YesNo) error {
		yesNo.defaultValue = defaultValue
		return nil
	}
}
