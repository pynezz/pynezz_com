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
	Title     = "text-4xl font-bold text-text font-mono underline-offset-2 text-left px-2 pr-10"
	Icon      = "w-8 h-8 px-2 transform hover:scale-105 transition-transform duration-200 ease-in-out"
	Article   = "flex flex-col text-left max-w-64"
	Wrapper   = "flex flex-col justify-center items-center h-screen w-full"

	// Card
	CardStyle  = "flex flex-row content-between font-normal text-text rounded-md shadow-lg bg-mantle flex-wrap p-4 m-4	"
	CardBtn    = "button bg-surface0 text-text rounded px-4 py-2 transition duration-200 ease-in-out w-max stroke-2 stroke-mauve"
	BtnBlue    = "bg-blue hover:bg-blue-700"
	BtnHover   = "hover:text-text hover:bg-surface1 hover:stroke-lavender stroke-2 str"
	SubHeading = "text-lg font-bold bg-clip-text text-transparent bg-gradient-to-br from-lavender to-mauve text-wrap first-letter:text-xl px-2 pr-8 pt-4"
	CardIcon   = "w-10 transform hover:scale-105 transition-transform duration-200 ease-in-out"
	CardImg    = "rounded-md w-auto h-auto p-4 m-4"
)

const baseURL = "https://pynezz.dev"
