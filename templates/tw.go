package templates

import "github.com/pynezz/pynezz_com/templates/layout"

const (
	twContainer = string(layout.Container)
	twTitle     = string(layout.Title)
	twIcon      = string(layout.Icon)
	twArticle   = string(layout.Article)
	twAboutMe   = "flex flex-col items-center max-w-prose mx-auto p-4 bg-mantle rounded-md my-2 py-2"
	postsList   = "flex flex-col gap-5 text-left max-w-[960px] p-4 shadow-lg bg-mantle rounded-md text-sans text-text"
)

const BaseURL = "https://pynezz.dev"

// css aboutMe() {
//   margin-top: 20px;
//   margin-left: 20px;
//   margin-right: 20px;
//   flex-direction: column;
//   align-items: center;
//   font-family: 'Hack Nerd Font', 'Fira Mono', 'Fira Code', monospace;
// }
