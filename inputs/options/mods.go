package options

type Mod func(options *Options) error

func WithDefaults(defaults []int) Mod {
	return func(options *Options) error {
		options.defaultValue = defaults
		return nil
	}
}

func WithMultipleSelection(multiple bool) Mod {
	return func(options *Options) error {
		options.multipleSelect = multiple
		return nil
	}
}
