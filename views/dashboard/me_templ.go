// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package dashboard

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/TimRobillard/handicap_tracker/views/layout"
	"strings"
)

func BigNum(num string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(num)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 10, Col: 12}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Card(title, value, url, fontColor string, loaded bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if loaded {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 15, Col: 17}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"bg-white rounded-3xl flex flex-col items-center justify-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 = []any{fmt.Sprintf("text-5xl font-semibold text-%s", fontColor)}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(value)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 17, Col: 11}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"text-slate-600\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strings.Title(title))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 19, Col: 51}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-slate-100 rounded-3xl flex flex-col items-center justify-center\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(url)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 22, Col: 94}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"load\" hx-swap=\"outerHTML\"><p class=\"text-5xl font-semibold\"><span class=\"block animate-spin text-slate-500 text-lg\"><i class=\"fa-solid fa-spinner\"></i></span></p><p class=\"text-slate-600\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(strings.Title(title))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 28, Col: 51}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Chart() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var11 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var11 == nil {
			templ_7745c5c3_Var11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center justify-center h-full\"><canvas id=\"myChart\" height=\"200px\" width=\"240px\"></canvas></div><script>\n  const ctx = document.getElementById('myChart');\n\n  new Chart(ctx, {\n    type: 'line',\n    data: {\n      labels: ['', 'May','','June', 'July', ''],\n      datasets: [{\n\t\tlabel: '',\n        data: [20.1, 20.3, 20.0, 19.2, 18.9, 21],\n        borderWidth: 2,\n      }]\n    },\n\t\tresponsive: true,\n    options: {\n\t\tresponsive: true,\n      scales: {\n        y: {\n          beginAtZero: true,\n\t\t  min: 18, max: 22\n        }\n      },\n\t  fill: 'rgb(22 101 52 )',\n\t  elements: {\n\t\tline: {\n\t\ttension: 0.2\n\t\t},\n\t  },\n\t\tlayout: {\n\t\t\tpadding: 20\n\t\t},\n\t\tplugins: {\n\t\t\tlegend: {\n\t\t\t\tdisplay: false,\n\t\t\t}\n\t\t}\n    }\n  });\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Me() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = layouts.Base("Dashboard | Golf IndX").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"h-screen bg-slate-200 w-screen flex flex-col sm:flex-row\"><aside class=\"bg-white w-full sm:w-16 lg:w-3/12 flex flex-col\"><a href=\"/\" class=\"text-black font-semibold text-sm pt-2 px-2 sm:p-4 sm:mb-16\">Golf IndX</a><div class=\"p-2 sm:p-0\"><ul class=\"w-full no-scrollbar overflow-x-scroll sm:overflow-visible flex-1 flex sm:flex-col gap-2 sm:gap-10\"><li class=\"group min-w-24 sm:min-w-0 sm:w-full p-2 sm:p-0 sm:pl-4 flex bg-green-800 sm:bg-white rounded-full border sm:border-none border-green-800\"><a href=\"#me\" class=\"flex items-center justify-center lg:justify-start text-xs w-full font-semibold text-white sm:text-black group-hover:text-black transition-all\"><i class=\"fa-solid fa-user mr-4 text-white sm:text-green-800 transition-al\"></i><span class=\"block sm:hidden lg:block\">Me</span></a><div class=\"hidden sm:block bg-green-800 w-1 rounded-full translate-x-0.5\"></div></li><li class=\"group min-w-24 sm:min-w-0 sm:w-full p-2 sm:p-0 sm:pl-4 flex bg-white sm:bg-white rounded-full border sm:border-none border-green-800\"><a href=\"#me\" class=\"flex items-center justify-center lg:justify-start text-xs w-full text-slate-500 sm:group-hover:text-black transition-all\"><i class=\"fa-solid fa-user-group mr-4 text-slate-500 sm:group-hover:text-black transition-al\"></i><span class=\"block sm:hidden lg:block\">Friends</span></a><div class=\"hidden sm:group-hover:block bg-opacity-20 bg-green-800 w-1 rounded-full translate-x-0.5\"></div></li><li class=\"group min-w-24 sm:min-w-0 sm:w-full p-2 sm:p-0 sm:pl-4 flex bg-white sm:bg-white rounded-full border sm:border-none border-green-800\"><a href=\"#me\" class=\"flex items-center justify-center lg:justify-start text-xs w-full text-slate-500 sm:group-hover:text-black transition-all\"><i class=\"fa-solid fa-location-dot mr-4 text-slate-500 sm:group-hover:text-black transition-al\"></i><span class=\"block sm:hidden lg:block\">Local</span></a><div class=\"hidden sm:group-hover:block bg-opacity-20 bg-green-800 w-1 rounded-full translate-x-0.5\"></div></li><li class=\"group min-w-24 sm:min-w-0 sm:w-full p-2 sm:p-0 sm:pl-4 flex bg-white sm:bg-white rounded-full border sm:border-none border-green-800\"><a href=\"#me\" class=\"flex items-center justify-center lg:justify-start text-xs w-full text-slate-500 sm:group-hover:text-black transition-all\"><i class=\"fa-solid fa-earth-americas mr-4 text-slate-500 sm:group-hover:text-black transition-al\"></i><span class=\"block sm:hidden lg:block\">Global</span></a><div class=\"hidden sm:group-hover:block bg-opacity-20 bg-green-800 w-1 rounded-full translate-x-0.5\"></div></li></ul></div></aside><section class=\"w-full h-screen overflow-scroll\"><h2 class=\"hidden sm:block p-8 font-semibold text-3xl\">My IndX</h2><div class=\"grid grid-cols-3 md:grid-cols-5 dashboard-rows gap-2 sm:gap-4 p-4 w-full overflow-scroll\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Card("current", "", "/indx/current", "green-800", false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-white rounded-3xl flex flex-col items-center justify-center\"><p class=\"text-5xl font-semibold text-slate-700\">92</p><p class=\"text-slate-600\">Latest Round</p></div><div class=\"bg-white rounded-3xl flex flex-col items-center justify-center\"><p class=\"text-5xl font-semibold text-slate-700\">88</p><p class=\"text-slate-600\">Lowest Round</p></div><div class=\"hidden md:block bg-white rounded-3xl col-span-2 row-span-3 p-4 overflow-y-scroll\"><ul><li>102 | Manderley on the Green</li><li>92 | Dragonfly</li><li>96 | Manderley on the Green</li><li>96 | Cedarhill</li><li>34 | Amberwood</li><li>101 | Manderley on the Green</li><li>100 | Cedarhill</li><li>98 | Cedarhill</li><li>94 | Manderley on the Green</li><li>93 | Manderley on the Green</li><li>92 | Manderley on the Green</li><li>102 | Manderley on the Green</li><li>92 | Dragonfly</li><li>96 | Manderley on the Green</li><li>96 | Cedarhill</li><li>34 | Amberwood</li><li>101 | Manderley on the Green</li><li>100 | Cedarhill</li><li>98 | Cedarhill</li><li>94 | Manderley on the Green</li><li>93 | Manderley on the Green</li><li>92 | Manderley on the Green</li></ul></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Card("lowest", "", "/indx/lowest", "", false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-white rounded-3xl col-span-2 row-span-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Chart().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Card("highest", "", "/indx/highest", "", false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"block md:hidden bg-white rounded-3xl col-span-3 row-span-3\">8</div></div></section><div class=\"hidden text-green-400\">generate class</div></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
