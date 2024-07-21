package parser

import (
	"fmt"
)

/*
DECLARATION OF HEADING ELEMENTS
We want to enforce heading to be of type h2...h6
*/
func (h h2) isHeading() {}

//	func (h h2) String() string {
//		return string(h)
//	}
func (h h2) HTMLTag() string {
	return "h2"
}

// func (h h2) Class() string {
// 	return h.style.
// }

func (h h3) isHeading() {}

//	func (h h3) String() string {
//		return string(h)
//	}
func (h h3) HTMLTag() string {
	return "h3"
}

func (h h4) isHeading() {}

//	func (h h4) String() string {
//		return string(h)
//	}
func (h h4) HTMLTag() string {
	return "h4"
}

func (h h5) isHeading() {}

// func (h h5) String() string {
// 	return string(h)
// }

func (h h6) isHeading() {}

// func (h h6) String() string {
// 	return string(h)
// }

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

func (t strong) isTextContent() {}
func (t strong) String() string {
	return string(t)
}

func (t italic) isTextContent() {}
func (t italic) String() string {
	return string(t)
}

func (t code) isTextContent() {}
func (t code) String() string {
	return t.String()
}

func (t codeblock) isTextContent() {}
func (t codeblock) String() string {
	return string(t)
}

func (t pre) isTextContent() {}
func (t pre) String() string {
	return string(t)
}

func (t blockquote) isTextContent() {}
func (t blockquote) String() string {
	return string(t)
}

func (t hr) isTextContent() {}
func (t hr) String() string {
	return string(t)
}

func (t br) isTextContent() {}
func (t br) String() string {
	return string(t)
}

func (t table) isTextContent() {}
func (t table) String() string {
	return string(t)
}

func (t tr) isTextContent() {}
func (t tr) String() string {
	return string(t)
}

func (t th) isTextContent() {}
func (t th) String() string {
	return string(t)
}

func (t td) isTextContent() {}
func (t td) String() string {
	return string(t)
}

func (t ul) isTextContent() {}
func (t ul) String() string {
	return string(t)
}

func (t ol) isTextContent() {}
func (t ol) String() string {
	return string(t)
}

func (t li) isTextContent() {}
func (t li) String() string {
	return string(t)
}

func (s Section) String() string {
	content := ""
	for _, c := range s.TextContent {
		switch c.(type) {
		case textCodeblock, textTable, textUl, textOl:
			// content += fmt.Sprintf("%s", c.String())
			content += c.String()
		default:
			content += fmt.Sprintf("<p>%s</p>\n", c.String())
		}
	}
	return fmt.Sprintf("<section class=\"flex flex-col align-start overflow-x-auto\"><%s class=\"%s\">%s</%s>\n%s</section>\n", s.Title.HTMLTag(), s.Title.Class(), s.Title.String(), s.Title.HTMLTag(), content)
}
		