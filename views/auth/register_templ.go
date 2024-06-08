// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package auth

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/TimRobillard/handicap_tracker/views/layout"
)

func Register() templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center pt-4\"><h1 class=\"text-center font-bold text-4xl my-4\">Welcome to Tim's Todo List!</h1><form hx-ext=\"loading-states, response-targets\" hx-post=\"/auth/register\" hx-target-4*=\"#register-errors\" class=\"flex flex-col gap-4 md:border border-slate-300 sm:w-full md:shadow-md md:w-[696px] p-4 w-full rounded-md mb-4\"><input id=\"username\" autofocus class=\"border h-auto rounded-md flex-auto p-2\" type=\"text\" name=\"username\" value=\"\" placeholder=\"username\"> <input class=\"border h-auto rounded-md flex-auto p-2\" id=\"password\" type=\"password\" name=\"password\" value=\"\" placeholder=\"password\"><div id=\"register-errors\"></div><button type=\"submit\" class=\"px-4 py-2 bg-blue-400 text-white rounded-md sm:w-full\" data-loading-class=\"opacity-80\" data-loading-disable><span data-loading-class=\"hidden\">Login</span> <span data-loading>Loading</span></button></form><p>Have an account?</p><p><a href=\"/login\" class=\"text-blue-300\">login</a> instead!</p></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base("Register | Golf IndX").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
