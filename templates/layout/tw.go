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
	Title     = "text-4xl font-bold text-text font-mono text-left pr-10 border-b border-surface1 mb-2 mt-4"
	Icon      = "w-8 h-8 px-2 transform hover:scale-105 transition-transform duration-200 ease-in-out"
	Article   = "flex flex-col text-left max-w-[960px] p-4 shadow-lg bg-mantle rounded-md text-sans text-text space-y-2"
	Wrapper   = "flex flex-col justify-center items-center h-screen w-full"

	// Card
	CardStyle  = "flex flex-row content-between max-w-fit font-normal text-text rounded-md shadow-lg bg-mantle flex-wrap p-4 m-4 "
	CardBtn    = "button bg-sky text-base p-2 px-4 rounded-full underline decoration-dotted	transition-colors duration-200 ease-in-out "
	BtnBlue    = "bg-blue hover:bg-blue-700 "
	BtnHover   = "hover:text-text hover:bg-surface1 hover:stroke-lavender stroke-2 "
	SubHeading = "text-lg font-bold bg-clip-text text-transparent bg-gradient-to-br from-lavender to-mauve text-wrap first-letter:text-xl px-2 pr-8	"
	CardIcon   = "w-10 transform hover:scale-105 transition-transform duration-200 ease-in-out"
	CardImg    = "rounded-md w-auto h-auto p-4 m-2	"
)

const baseURL = "https://pynezz.dev"
const Link = "text-lavender underline hover:text-overlay1"

const TagBtn = "text-text px-2 m-2 rounded no-underline"
