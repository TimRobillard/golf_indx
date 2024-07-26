// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package dashboard

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"os"

	"github.com/TimRobillard/handicap_tracker/store"
	"github.com/TimRobillard/handicap_tracker/views/components"
)

func Me(u *store.UIUser, rounds [20]*store.CalcRound) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><title>Dashboard | Golf IndX</title><link rel=\"icon\" type=\"image/x-icon\" href=\"/public/favicon.ico\"><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" href=\"/public/styles.css\"><script src=\"https://cdn.jsdelivr.net/npm/chart.js\"></script><script src=\"https://code.jquery.com/jquery-3.7.1.min.js\" integrity=\"sha256-/JqT3SQfawRcv/BIHPThkBvs0OEvtFFmqPF/lYI/Cxo=\" crossorigin=\"anonymous\"></script><script src=\"https://unpkg.com/alpinejs\" defer></script><script src=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/js/all.min.js\"></script><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><script src=\"https://unpkg.com/htmx.org@1.9.12/dist/ext/loading-states.js\"></script><script src=\"https://unpkg.com/htmx.org@1.9.12/dist/ext/response-targets.js\"></script></head><body class=\"antialiased relative bg-slate-200 overflow-y-scroll\"><main class=\"relative h-auto min-h-screen w-screen flex flex-col sm:flex-row\"><a href=\"/dashboard/score\" class=\"sm:hidden fixed flex items-center justify-center bottom-4 right-4 h-16 w-16 bg-green-800 text-white font-semibold text-sm rounded-full shadow-md hover:shadow-lg hover:bg-green-900\"><i class=\"fa-solid fa-plus text-lg shadow-md\"></i></a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.SideBar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"w-full h-screen\"><div class=\"hidden sm:flex p-8 items-center\"><h2 class=\"font-semibold text-3xl\">My IndX</h2><div class=\"ml-auto\"><a href=\"/dashboard/score\" class=\"flex items-center justify-center h-12 w-12 bg-green-800 text-white font-semibold text-sm rounded-full shadow-md hover:shadow-lg hover:bg-green-900\"><i class=\"fa-solid fa-plus text-lg shadow-md\"></i></a></div></div><div class=\"grid grid-cols-3 md:grid-cols-5 dashboard-rows gap-2 sm:gap-4 p-4 w-full\"><div class=\"col-span-3 flex items-center justify-start gap-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.ProfilePic(u.Indx, &u.ProfilePic).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><span class=\"block text-2xl font-semibold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(u.Username)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 42, Col: 63}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"block text-slate-700\">Latest round ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(u.LastScore)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 43, Col: 69}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></div><div class=\"rounded-3xl col-start-1 col-end-4 row-start-4 row-end-6 md:col-start-4 md:col-end-6 md:row-start-1 md:row-end-4\"><div class=\"flex flex-col bg-white rounded-3xl p-4 pr-0.5 md:h-full\"><span class=\"font-semibold text-lg\">Calculated Rounds</span><div class=\"md:overflow-y-scroll\"><ul class=\"md:h-full sm:overflow-y-scroll flex flex-col gap-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, round := range rounds {
			if round != nil {
				templ_7745c5c3_Err = components.CalculatedRound(*round).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></div></div><div class=\"bg-white rounded-3xl col-span-3 row-span-2 p-4\"><span class=\"md:flex-1 font-semibold text-lg\">Indx</span><div class=\"flex items-center justify-center h-full\"><div id=\"chart-spinner\" class=\"animate-spin\"><i class=\"fa-solid fa-spinner text-4xl text-green-800\"></i></div><canvas id=\"myChart\" class=\"hidden\"></canvas><span id=\"chart-no-data\" class=\"hidden text-lg text-slate-500 text-center\">No Data Found</span></div></div></div></section><div class=\"hidden text-green-400\">generate class</div></main><div id=\"app-url\" class=\"hidden\" app-url=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(os.Getenv("APP_URL"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/me.templ`, Line: 74, Col: 65}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></body><script>\n\t\t const appUrl = document.getElementById('app-url').getAttribute('app-url');\n\t\t\tdocument.addEventListener('DOMContentLoaded', \n\t\t\tasync () => {\n\t\t\t\ttry {\n\t\t\t\t\tconst response = await fetch(`${appUrl}/dashboard/chart/me`);\n\t\t\t\t\tconst data = await response.json();\n\t\t\t\t\tif (data && data.data.length) {\n\t\t\t\t\tconst ctx = document.getElementById('myChart');\n\t\t\t\t\tnew Chart(ctx, {\n\t\t\t\t\t\ttype: 'line',\n\t\t\t\t\t\tdata: {\n\t\t\t\t\t\t\tlabels: data.labels,\n\t\t\t\t\t\t\tdatasets: [{\n\t\t\t\t\t\t\t\tlabel: '',\n\t\t\t\t\t\t\t\tdata: data.data,\n\t\t\t\t\t\t\t\tborderWidth: 2,\n\t\t\t\t\t\t\t}]\n\t\t\t\t\t\t},\n\t\t\t\t\t\tresponsive: true,\n\t\t\t\t\t\toptions: {\n\t\t\t\t\t\t\tfill: {\n\t\t\t\t\t\t\t\ttarget: 'origin',\n\t\t\t\t\t\t\t\tabove: 'rgb(22, 101, 52)',\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\tpointBackgroundColor: '#F2F2F2',\n\t\t\t\t\t\t\tpointBorderColor: '#333',\n\t\t\t\t\t\t\tpointRadius: 4,\n\t\t\t\t\t\t\tborderColor: 'rgb(22, 101, 52)',\n\t\t\t\t\t\t\tresponsive: true,\n\t\t\t\t\t\t\tscales: {\n\t\t\t\t\t\t\t\ty: {\n\t\t\t\t\t\t\t\t\tbeginAtZero: true,\n\t\t\t\t\t\t\t\t\tmin: data.min, max: data.max\n\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\telements: {\n\t\t\t\t\t\t\t\tline: {\n\t\t\t\t\t\t\t\ttension: 0.2\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\tlayout: {\n\t\t\t\t\t\t\t\tpadding: 20\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\tplugins: {\n\t\t\t\t\t\t\t\tlegend: {\n\t\t\t\t\t\t\t\t\tdisplay: false,\n\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}\n\t\t\t\t\t});\n\t\t\t\t\tctx.classList.remove('hidden')\n\t\t\t\t\t} else {\n\t\t\t\t\t\tdocument.getElementById('chart-no-data').classList.remove('hidden')\n\t\t\t\t\t}\n\t\t\t\t} catch(e) {\n\t\t\t\t\tconsole.log(e)\n\t\t\t\t\t\tdocument.getElementById('chart-no-data').classList.remove('hidden')\n\t\t\t\t} finally {\n\t\t\t\t\tdocument.getElementById('chart-spinner').remove()\n\t\t\t\t}\n\t\t\t});\n\t\t\t// TODO move to score form\n\t\t\t// const scoreForm = document.getElementById('score-form');\n\t\t\t// const scoreInputs = document.querySelectorAll('#score-form input');\n\t\t\t// const scoreOut = document.getElementById('score-out');\n\t\t\t// const scoreIn = document.getElementById('score-in');\n\t\t\t// const scoreTot = document.getElementById('score-tot');\n\t\t\t// const saveTot = document.getElementById('save-tot');\n\n\t\t\t// scoreForm.addEventListener('input', calculateScore)\n\n\t\t\t// function calculateScore() {\n\t\t\t// \tlet outVal = 0;\n\t\t\t// \tlet inVal = 0;\n\t\t\t// \tscoreInputs.forEach(input => {\n\t\t\t// \t\tconst hole = Number(input.id.split('-')[1]);\n\t\t\t// \t\tconst value = Number(input.value);\n\t\t\t// \t\tif (hole < 10) {\n\t\t\t// \t\t\toutVal += value;\n\t\t\t// \t\t\treturn;\n\t\t\t// \t\t}\n\t\t\t// \t\tinVal += value;\n\t\t\t// \t})\n\t\t\t// \tscoreOut.innerHTML = outVal;\n\t\t\t// \tscoreIn.innerHTML = inVal;\n\t\t\t// \tscoreTot.innerHTML = outVal + inVal;\n\t\t\t// \tsaveTot.innerHTML = outVal + inVal;\n\t\t\t// }\n\t\t</script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
