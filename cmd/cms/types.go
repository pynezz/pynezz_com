package cms

type ICommand interface {
	Run(args ...interface{}) interface{}
	Help() string
	Name() string
}

// BaseCommand struct to implement common fields and methods
type Command struct {
	HelpStr string
	NameStr string
}

// Commands - reflecting the commands in the cms.go in cms.go
type Config struct{ *Command }
type ParseAll struct{ *Command }
type PageTags struct{ *Command }
type EditPage struct{ *Command }
type ShowPage struct{ *Command }
type ListPages struct{ *Command }
type PageStatus struct{ *Command }
type CreatePage struct{ *Command }
type DeletePage struct{ *Command }
type PublishPage struct{ *Command }
type UnpublishPage struct{ *Command }

// A no operation command for typo checking
// Perfect opportunity to implement a spell checking algorithm
type Nop struct{ *Command }

// Help - returns the help message of the command
func (bc *Command) Help() string {
	return bc.HelpStr
}

// Name - returns the name of the command
func (bc *Command) Name() string {
	return bc.NameStr
}
