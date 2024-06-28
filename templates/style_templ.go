// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "strings"

const tw = `
@tailwind base;
@tailwind components;
@tailwind utilities;
`

func Style() templ.CSSClass {
	var templ_7745c5c3_CSSBuilder strings.Builder
	templ_7745c5c3_CSSBuilder.WriteString(`height:100%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:100%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:0;`)
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`background-color`, bg)))
	templ_7745c5c3_CSSBuilder.WriteString(string(templ.SanitizeCSS(`color`, subtxt)))
	templ_7745c5c3_CSSBuilder.WriteString(`font-family:'Inter', sans-serif;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:space-between;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Style`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

// css div() {
//     display: flex;
//     flex-direction: column;
//     align-items: center;
//     justify-content: center;
//     min-height: 100vh;
// }

// func RenderStyle() string {
//   return `
//     html, body {
//       height: 100%;
//       margin: 0;
//       padding: 0;
//       background-color: ` + bg + `;
//       color: ` + subtxt + `;
//       font-family: 'Inter', sans-serif;
//     }
//     .div {
//       display: flex;
//       flex-direction: column;
//       align-items: center;
//       justify-content: center;
//       min-height: 100vh;
//     }
//   `
// }
