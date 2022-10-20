package types

// Check displays a nagios compliant check as a struct
type Check struct {
	Output        string
	InMemoryValue string
	ExitCode      int
}
