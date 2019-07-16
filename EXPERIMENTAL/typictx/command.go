package typictx

// Command represent the command in CLI
type Command struct {
	Name      string
	ShortName string
	Action    Action
}