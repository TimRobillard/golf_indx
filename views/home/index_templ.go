// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/TimRobillard/handicap_tracker/views/layout"
)

func Index() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"flex flex-col items-center p-8 max-w-7xl mx-auto\"><div class=\"flex flex-row items-center w-full\"><h4 class=\"text-black font-semibold text-sm flex-1\">Golf IndX</h4><div class=\"flex items-center\"><a href=\"/login\" class=\"mr-4 font-semibold text-sm\">Login</a> <a href=\"/signup\" class=\"px-4 py-2 bg-green-800 text-white rounded-md\">Sign Up</a></div></div><div class=\"w-full h-56 my-16 flex flex-col items-center justify-center\"><h1 class=\"text-5xl font-bold text-green-800\">Golf IndX</h1><h2 class=\"text-slate-400 text-lg mt-2\">Improve your score by keeping track of your handicap index</h2></div><section class=\"flex items-center gap-20\"><div class=\"flex-1\"><h4 class=\"text-2xl text-green-800 font-semibold\">Know your IndX</h4><p class=\"text-slate-400 text-lg mt-4\">Easily know what your handicap index is at all times</p></div><div class=\"flex-1 rounded-lg overflow-hidden\"><img src=\"/public/images/stock_golf_swing.jpeg\" class=\"w-full\"></div></section><section class=\"flex items-center mt-24 gap-20\"><div class=\"flex-1 rounded-lg overflow-hidden\"><img src=\"/public/images/stock_graph.jpeg\" class=\"w-full\"></div><div class=\"flex-1\"><h4 class=\"text-2xl text-green-800 font-semibold\">Track your history</h4><p class=\"text-slate-400 text-lg mt-4\">See how you are improving by watching your handicap index change over time</p></div></section><section class=\"mt-20\"><h3 class=\"font-semibold text-2xl text-green-800 mb-2\">Features</h3><div class=\"grid grid-cols-2 gap-x-12 md:gap-x-60 gap-y-10\"><div><h4 class=\"font-semibold\">Track your score</h4><p class=\"text-slate-400 text-sm md:text-md\">After your round, input your score, and Handitrak will automatically update your handicap. We’ll let you know if it’s changed, and how you are trending.</p></div><div><h4 class=\"font-semibold\">Compare with friends</h4><p class=\"text-slate-400 text-sm md:text-md\">View your friends recent rounds and progress. See how you stack up, so you’re ready for your next round of golf.</p></div><div><h4 class=\"font-semibold\">Keep rounds interesting</h4><p class=\"text-slate-400 text-sm md:text-md\">Use your score differential to calculate your net score. Know which holes to give or take a stroke, and keep the game fun for everyone.</p></div><div><h4 class=\"font-semibold\">Watch your progress</h4><p class=\"text-slate-400 text-sm md:text-md\">Use the analytics page to see your handicap index change over time. See your all time high or low index, or compare against your friends and local golfers.</p></div></div></section></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
