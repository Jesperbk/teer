package options

// OutputOptions incapsulates the options that can be provided to teer through the command line.
type OutputOptions struct {
	OutputPath string `arg:"positional,required"`
	DoTruncate bool   `arg:"-t,--truncate" help:"Truncate existing files, rather than appending"`
}
