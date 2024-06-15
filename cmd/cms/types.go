package cms

type Command struct {
	HelpString string

	Run func(any) any
}

func (c Command) Help(args ...string) string {
	return c.HelpString
}
