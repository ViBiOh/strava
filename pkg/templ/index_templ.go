// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.590
package templ

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Root(uri, nonce, title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"utf-8\"><meta name=\"format-detection\" content=\"telephone=no\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = RootStyle(nonce).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</head><body class=\"no-margin\"><div class=\"content\"><header class=\"flex padding\"><h1 class=\"no-margin no-padding\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.URL(uri)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/templ/index.templ`, Line: 16, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></h1></header><article class=\"padding center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</article></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func RootStyle(nonce string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\" nonce=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(nonce))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">\n    :root {\n      --primary: cornflowerblue;\n      --success: limegreen;\n      --danger: salmon;\n      --dark: #000000;\n      --grey: #272727;\n      --white: silver;\n\n      --icon-small: 1.6rem;\n      --icon-size: 2.4rem;\n      --icon-large: 4.8rem;\n\n      --space-size: 1rem;\n    }\n\n    * {\n      box-sizing: border-box;\n    }\n\n    html {\n      font-size: 62.5%;\n    }\n\n    body {\n      -webkit-overflow-scrolling: touch;\n      background-color: var(--dark);\n      height: 100vh;\n    }\n\n    body,\n    button,\n    input {\n      color: var(--white);\n      font-family:\n        -apple-system,\n        'Segoe UI',\n        'Roboto',\n        'Oxygen-Sans',\n        'Ubuntu',\n        'Cantarell',\n        'Helvetica Nue',\n        sans-serif;\n      font-size: 1.6rem;\n      font-style: normal;\n      font-weight: 400;\n    }\n\n    input {\n      color: var(--dark);\n    }\n\n    a {\n      color: var(--white);\n      text-decoration: none;\n    }\n\n    a:hover {\n      color: var(--primary);\n      text-decoration: underline;\n    }\n\n    .primary {\n      color: var(--primary);\n    }\n\n    .success {\n      color: var(--success);\n    }\n\n    .danger {\n      color: var(--danger);\n    }\n\n    .grey {\n      color: var(--grey);\n    }\n\n    .white {\n      color: var(--white);\n    }\n\n    .bg-primary,\n    .bg-primary:hover {\n      background-color: var(--primary);\n      color: var(--dark);\n      text-decoration: none;\n    }\n\n    .bg-success,\n    .bg-success:hover {\n      background-color: var(--success);\n      color: var(--dark);\n      text-decoration: none;\n    }\n\n    .bg-danger,\n    .bg-danger:hover {\n      background-color: var(--danger);\n      color: var(--dark);\n      text-decoration: none;\n    }\n\n    .bg-grey,\n    .bg-grey:hover {\n      background-color: var(--grey);\n      color: var(--white);\n      text-decoration: none;\n    }\n\n    .button {\n      border-radius: 4px;\n      border: 0;\n      cursor: pointer;\n      display: inline-block;\n      margin: 0;\n      padding: var(--space-size);\n      text-decoration: none;\n    }\n\n    .button-icon {\n      background-color: transparent;\n    }\n\n    .icon {\n      background-position: center center;\n      background-repeat: no-repeat;\n      color: var(--white);\n      display: inline-block;\n      height: var(--icon-size);\n      text-decoration: none;\n      vertical-align: middle;\n      width: var(--icon-size);\n    }\n\n    .icon-large {\n      height: var(--icon-large);\n      width: var(--icon-large);\n    }\n\n    .icon-small {\n      height: var(--icon-small);\n      width: var(--icon-small);\n    }\n\n    .icon-overlay {\n      height: var(--icon-large);\n      left: calc((100% - var(--icon-large)) / 2);\n      pointer-events: none;\n      position: absolute;\n      top: calc((100% - var(--icon-large)) / 2);\n      width: var(--icon-large);\n    }\n\n    .modal {\n      align-items: center;\n      background-color: rgba(84, 84, 84, 0.75);\n      display: none;\n      height: 100vh;\n      justify-content: center;\n      left: 0;\n      pointer-events: none;\n      position: fixed;\n      top: 0;\n      width: 100vw;\n    }\n\n    .modal-content {\n      background-color: var(--dark);\n      display: flex;\n      flex-direction: column;\n      max-height: 100%;\n      max-width: 100%;\n      pointer-events: auto;\n    }\n\n    .header {\n      background-color: var(--grey);\n      margin-top: 0;\n      padding: calc(var(--space-size) / 2) var(--space-size);\n      text-align: left;\n    }\n\n    .flex {\n      align-items: center;\n      display: flex;\n    }\n\n    .flex-grow {\n      flex: 1 1;\n    }\n\n    .center {\n      text-align: center;\n    }\n\n    .padding {\n      padding: var(--space-size);\n    }\n\n    .padding-half {\n      padding: calc(var(--space-size) / 2);\n    }\n\n    .no-padding {\n      padding: 0;\n    }\n\n    .margin {\n      margin: var(--space-size);\n    }\n\n    .margin-top {\n      margin-top: var(--space-size);\n    }\n\n    .margin-right {\n      margin-right: var(--space-size);\n    }\n\n    .margin-bottom {\n      margin-bottom: var(--space-size);\n    }\n\n    .margin-left {\n      margin-left: var(--space-size);\n    }\n\n    .margin-half {\n      margin: calc(var(--space-size) / 2);\n    }\n\n    .no-margin {\n      margin: 0;\n    }\n\n    .ellipsis {\n      overflow: hidden;\n      text-align: left;\n      text-overflow: ellipsis;\n      white-space: nowrap;\n    }\n\n    .full {\n      width: 100%;\n    }\n\n    .max-full {\n      max-width: 100%;\n    }\n\n    @media print {\n      body::after {\n        content: 'Save ink, share link.';\n      }\n\n      body > * {\n        display: none !important;\n      }\n    }\n  </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
