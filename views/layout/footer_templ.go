// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Footer() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section id=\"footer\" class=\"w-full md:max-w-7xl flex flex-col md:flex-row px-8 mx-auto py-16 gap-6\"><div class=\"flex flex-col h-full flex-1\"><a href=\"/\" class=\"text-lg font-semibold text-center md:text-left hover:text-slate-800\">Golf IndX</a></div><div class=\"flex flex-col flex-1 gap-2\"><a href=\"#footer\" class=\"text-slate-700 text-center md:text-left text-sm hover:text-slate-900\">Products</a> <a href=\"#footer\" class=\"text-slate-700 text-center md:text-left text-sm hover:text-slate-900\">Clients</a> <a href=\"#footer\" class=\"text-slate-700 text-center md:text-left text-sm hover:text-slate-900\">Careers</a> <a href=\"#footer\" class=\"text-slate-700 text-center md:text-left text-sm hover:text-slate-900\">Support</a> <a href=\"/login\" class=\"text-slate-700 text-center md:text-left text-sm hover:text-slate-900\">Sign in</a></div><div class=\"flex flex-col gap-4 flex-1\"><div><h5 class=\"font-semibold text-center md:text-left\">Support</h5><a href=\"mailto:hello@timrobillard.ca\" class=\"block text-slate-700 hover:text-slate-900 text-center md:text-left\">hello@timrobillard.ca</a></div><div><h5 class=\"font-semibold text-center md:text-left\">Office</h5><p class=\"text-slate-700 text-center md:text-left\">123 Some St</p><p class=\"text-slate-700 text-center md:text-left\">Ottawa, Ontario</p><p class=\"text-slate-700 text-center md:text-left\">K1K 2M2</p></div></div></section><section class=\"w-full max-w-7xl hidden md:flex p-8\"><p class=\"text-slate-700 flex-1 text-sm\">&copy Tim Robillard</p><a href=\"/terms-of-service\" class=\"text-slate-700 flex-1 text-sm\">Terms & Privacy</a> <a href=\"/accessibility\" class=\"text-slate-700 flex-1 text-sm\">Accessibility</a></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
