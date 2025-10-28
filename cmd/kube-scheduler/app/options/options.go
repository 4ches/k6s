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

	nfs := cliflag.NewNamedFlagSets()
	fs := nfs.FlagSet("misc")
	fs.StringVar(&o.ConfigFile, "config", o.ConfigFile, "The path to the configuration file.")
	fs.StringVar(&o.WriteConfigTo, "write-config-to", o.WriteConfigTo, "If set, write the configuration values to this file and exit.")
	fs.StringVar(&o.Master, "master", o.Master, "The address of the Kubernetes API server (overrides any value in kubeconfig)")

	o.Flags = &nfs
}
