package options

// OutputOptions incapsulates the options that can be provided to teer through the command line.
type OutputOptions struct {
	OutputPath   string
	ExistsAction ExistsAction
}
