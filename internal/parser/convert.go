package parser

import "fmt"

/*
DECLARATION OF HEADING ELEMENTS
We want to enforce heading to be of type h2...h6
*/
func (h h2) isHeading() {}
func (h h2) String() string {
	return string(h)
}

func (h h3) isHeading() {}
func (h h3) String() string {
	return string(h)
}

func (h h4) isHeading() {}
func (h h4) String() string {
	return string(h)
}

func (h h5) isHeading() {}
func (h h5) String() string {
	return string(h)
}

func (h h6) isHeading() {}
func (h h6) String() string {
	return string(h)
}

func (t p) isTextContent() {}
func (t p) String() string {
	return string(t)
}

func (t a) isTextContent() {}
func (t a) String() string {
	return string(t)
}

func (t img) isTextContent() {}
func (t img) String() string {
	return string(t)
}

func (t em) isTextContent() {}
func (t em) String() string {
	return string(t)
}

func (t b) isTextContent() {}
func (t b) String() string {
	return string(t)
}

func (s Section) String() string {
	content := ""
	for _, c := range s.TextContent {
		content += fmt.Sprintf("<p>%s</p>\n", c.String())
	}
	return fmt.Sprintf("<section><%s>%s</%s>\n%s</section>\n", s.Title, s.Title.String(), s.Title, content)
}
