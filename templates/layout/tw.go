package layout

type TwCSS struct {
}

type ITW interface {
	String() string
}

func (t TwCSS) Container() string {
	return Container
}

func (t TwCSS) Title() string {
	return Title
}

func (t TwCSS) Icon() string {
	return Icon
}

const BaseURL = "https://pynezz.dev"

func (t TwCSS) Article() string {
	return Article
}

const (
	Container = "flex flex-1 justify-normal	mx-auto px-4 bg-overlay0 h-fit w-max"
	Title     = "text-4xl font-bold text-text font-mono text-left 	border-b border-surface1 mb-2 mt-4"
	Icon      = "w-8 h-8 px-2 transform hover:scale-105 transition-transform duration-200 ease-in-out"
	Article   = "flex flex-col items-center max-w-prose mx-auto p-4 bg-mantle rounded-md my-2 py-2"

	// Nav
	Wrapper = "flex flex-col justify-center items-center h-screen w-full"

	// About
	AboutLinks = "flex flex-col lg:flex-row items-start sm:items-center content-between font-normal text-text rounded-md shadow-lg bg-mantle mx-auto max-w-screen-md px-2"

	// Link Cards
	LinkCard      = "flex flex-col sm:flex-row items-start sm:items-center content-between font-normal text-text rounded-md shadow-lg bg-mantle mx-auto max-w-screen-md px-2	"
	LinkCardTitle = "text-4xl font-bold text-text px-2 border-b border-surface1 mb-2 mt-4 font-mono"

	// Card
	CardStyle  = "flex flex-col sm:flex-row items-start sm:items-center content-between font-normal text-text rounded-md shadow-lg bg-mantle my-4 mx-auto max-w-screen-md"
	CardBtn    = "button max-w-fit bg-sky text-base mx-1 p-2 px-4 rounded-full underline decoration-dotted transition-colors duration-200 ease-in-out "
	BtnBlue    = "bg-blue hover:bg-blue-700 "
	BtnHover   = "hover:text-text hover:bg-surface1 hover:stroke-lavender stroke-2 "
	SubHeading = "text-lg font-bold bg-clip-text text-transparent bg-gradient-to-br from-lavender to-mauve text-wrap first-letter:text-xl px-2 pr-8	"
	CardIcon   = "w-10 transform hover:scale-105 transition-transform duration-200 ease-in-out"
	CardImg    = "rounded-md w-auto h-auto p-4 m-2 "
)

const baseURL = "https://pynezz.dev"
const Link = "text-lavender hover:text-text hover:border-b hover:border-b-green visited:text-mauve "

const TagBtnHover = "hover:border-b-green"
const TagBtn = "text-green px-2 m-1 rounded no-underline border border-crust " + TagBtnHover
