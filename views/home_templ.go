// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/markbates/goth"

func Home(user goth.User) templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1>mnemstart</h1><form action=\"https://duckduckgo.com\" method=\"get\"><input type=\"text\" id=\"search\" name=\"q\" placeholder=\"Search...\"></form><script>\r\n            let inputSequence = [];\r\n            const keymaps = {\r\n                ' gh': 'https://github.com',\r\n                ' yt': 'https://youtube.com',\r\n            };\r\n\r\n            document.addEventListener('keydown', (event) => {\r\n                const key = event.key;\r\n\r\n                // ignore keypresses when an input is focused\r\n                const activeElement = document.activeElement;\r\n                if (activeElement && (activeElement.tagName.toLowerCase() === 'input')) {\r\n                    return;\r\n                }\r\n\r\n                inputSequence.push(key);\r\n\r\n                const longestSequence = Math.max(...Object.keys(keymaps).map(s => s.length));\r\n\r\n                if (inputSequence.length > longestSequence) {\r\n                    inputSequence.shift();\r\n                }\r\n\r\n                const inputString = inputSequence.join('');\r\n\r\n                for (const mapping in keymaps) {\r\n                    if (inputString.endsWith(mapping)) {\r\n                        window.location.href = keymaps[mapping];\r\n                        inputSequence = [];\r\n                        break;\r\n                    }\r\n                }\r\n            });\r\n\r\n            // default mappings\r\n            // 'esc' clears the input sequence and unfocuses any input fields\r\n            document.addEventListener('keydown', (event) => {\r\n                if (event.key === 'Escape') {\r\n                    event.preventDefault();\r\n                    inputSequence = [];\r\n                    document.activeElement.blur();\r\n                }\r\n            });\r\n\r\n            // 'i' focuses the search input - i.e. 'insert mode'\r\n            document.addEventListener('keydown', (event) => {\r\n                if (event.key === 'i') {\r\n                    event.preventDefault();\r\n                    document.getElementById('search').focus();\r\n                }\r\n            });\r\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Page(true, user).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
