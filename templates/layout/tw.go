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

func (t TwCSS) Article() string {
	return Article
}

const (
	Container = "flex flex-1 justify-normal	mx-auto px-4 bg-overlay0 h-fit w-max"
	Title     = "text-4xl font-bold text-center text-text font-mono"
	Icon      = "w-8 h-8 transform hover:scale-110 transition-transform duration-500 ease-in-out"
	Article   = "flex flex-col text-left max-w-64"
	Wrapper   = "flex flex-col justify-center items-center h-full w-full"
)
