package options

import cliflag "k6s/component-base/cli/flag"

// Options has all the params needed to run a Scheduler.
type Options struct {
	// ConfigFile is the location of the scheduler server's configuration file.
	ConfigFile string

	// WriteConfigTo is the path where the default configuration will be written.
	WriteConfigTo string

	Master string

	// Flags hold the parsed CLI flags.
	Flags *cliflag.NamedFlagSets
}

// NewOptions returns default Scheduler' app options.
func NewOptions() *Options {
	o := &Options{}
	o.initFlags()

	return o
}

func (o *Options) initFlags() {
	if o.Flags != nil {
		return
	}

	_ = cliflag.NamedFlagSets{}

	// TODO: set flags.
}
