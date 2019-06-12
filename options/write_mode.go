package options

// ExistsAction values indicate how an existing output file should be handled.
// Available options are truncate (ExistsActionTruncate) and append (ExistsActionAppend).
type ExistsAction uint

const (
	// ExistsActionTruncate indicates that when opening an existing file,
	// the existing data should be truncated when first opened.
	ExistsActionTruncate ExistsAction = iota
	// ExistsActionAppend indicates that when opening an existing file,
	// the existing data should be kept. New data will be appended to the end of the file.
	ExistsActionAppend ExistsAction = iota
)
