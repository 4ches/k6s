package flag

import "github.com/spf13/pflag"

// NamedFlagSetsOption is a function to configure NamedFlagSets.
type NamedFlagSetsOption func(*NamedFlagSets)

// WithNormalizeNameFunc returns a NamedFlagSetsOption to set NormalizeNameFunc.
func WithNormalizeNameFunc(f func(f *pflag.FlagSet, name string) pflag.NormalizedName) NamedFlagSetsOption {
	return func(nfs *NamedFlagSets) {
		nfs.normalizeNameFunc = f
	}
}

// NamedFlagSets stores named flag sets in the order of calling FlagSet.
type NamedFlagSets struct {
	// FlagSets stores the flag sets by name.
	FlagSets map[string]*pflag.FlagSet

	// order is an ordered list of flag set names.
	order []string
	// normalizeNameFunc is the normalize function which used to initialize FlagSets created by NamedFlagSets.
	normalizeNameFunc func(f *pflag.FlagSet, name string) pflag.NormalizedName
}

// NewNamedFlagSets creates a named flag sets.
func NewNamedFlagSets(opts ...NamedFlagSetsOption) NamedFlagSets {
	nfs := NamedFlagSets{
		FlagSets:          make(map[string]*pflag.FlagSet),
		normalizeNameFunc: pflag.CommandLine.GetNormalizeFunc(),
	}
	for _, opt := range opts {
		opt(&nfs)
	}

	return nfs
}

// FlagSet returns the flag set with the given name and adds it to the
// ordered name list if it is not in there yet.
func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if _, ok := nfs.FlagSets[name]; !ok {
		nfs.FlagSets[name] = pflag.NewFlagSet(name, pflag.ExitOnError)
		nfs.order = append(nfs.order, name)
	}

	return nfs.FlagSets[name]
}
