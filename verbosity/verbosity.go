package verbosity

type Level int

const (
	Quiet       Level = -1
	Normal      Level = 0
	Verbose     Level = 1
	VeryVerbose Level = 2
	Debug       Level = 3
)
