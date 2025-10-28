module k6s

go 1.25.1

require (
	github.com/spf13/cobra v1.9.1
	k6s/component-base v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)

replace k6s/component-base => ./k6s/component-base
