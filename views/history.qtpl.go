// Code generated by qtc from "history.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/history.qtpl:1
package views

//line views/history.qtpl:1
import "fmt"

//line views/history.qtpl:2
import "net/http"

//line views/history.qtpl:4
import "github.com/bouncepaw/mycorrhiza/l18n"

//line views/history.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/history.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/history.qtpl:6
func StreamHistory(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, list string, lc *l18n.Localizer) {
//line views/history.qtpl:6
	qw422016.N().S(`
<main class="main-width">
	<article class="history">
		<h1>`)
//line views/history.qtpl:9
	qw422016.N().S(fmt.Sprintf(lc.Get("ui.history_title"), beautifulLink(hyphaName)))
//line views/history.qtpl:9
	qw422016.N().S(`</h1>
		`)
//line views/history.qtpl:10
	qw422016.N().S(list)
//line views/history.qtpl:10
	qw422016.N().S(`
	</article>
</main>
`)
//line views/history.qtpl:13
}

//line views/history.qtpl:13
func WriteHistory(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, list string, lc *l18n.Localizer) {
//line views/history.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/history.qtpl:13
	StreamHistory(qw422016, rq, hyphaName, list, lc)
//line views/history.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line views/history.qtpl:13
}

//line views/history.qtpl:13
func History(rq *http.Request, hyphaName, list string, lc *l18n.Localizer) string {
//line views/history.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
//line views/history.qtpl:13
	WriteHistory(qb422016, rq, hyphaName, list, lc)
//line views/history.qtpl:13
	qs422016 := string(qb422016.B)
//line views/history.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
//line views/history.qtpl:13
	return qs422016
//line views/history.qtpl:13
}
