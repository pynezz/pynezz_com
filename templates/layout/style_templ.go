// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

const (
	bg       = "#1e1e2e" // catppuccin Mocha: base
	txt      = "#cdd6f4" // text
	subtxt   = "#b5bfe2" // text2
	text3    = "#a6adc8" // text3
	darkTxt  = "#313244" // surface0
	red      = "#f38ba8"
	green    = "#a6e3a1"
	overlay2 = "#9399b2" // light gray
	overlay1 = "#7f849c" // semi gray
	overlay0 = "#6c7086" // semi gray
	surface2 = "#585b70" // light gray
	surface0 = "#313244" // dark gray
	mantle   = "#181825" // darkest gray
	crust    = "#11111b" // darkest
)

func Style() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n    html {\n        font-size: 16px;\n        box-sizing: border-box;\n        padding: 0;\n        margin: 0;\n    }\n\n    body {\n        font-family: 'Roboto', sans-serif;\n        margin: 0;\n        padding: 0;\n        background-color: { bg };\n    }\n    a {\n        text-decoration: none;\n    }\n    a:link {\n      text-decoration: none;\n    }\n    a:visited {\n      text-decoration: none;\n    }\n    a:hover {\n      text-decoration: none;\n    }\n    a:active {\n      text-decoration: none;\n    }\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
