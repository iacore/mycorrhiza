// Code generated by qtc from "stuff.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/stuff.qtpl:1
package views

//line views/stuff.qtpl:1
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/stuff.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/stuff.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/stuff.qtpl:3
func streamcommonScripts(qw422016 *qt422016.Writer) {
//line views/stuff.qtpl:3
	qw422016.N().S(`
`)
//line views/stuff.qtpl:4
	for _, scriptPath := range cfg.CommonScripts {
//line views/stuff.qtpl:4
		qw422016.N().S(`
<script src="`)
//line views/stuff.qtpl:5
		qw422016.E().S(scriptPath)
//line views/stuff.qtpl:5
		qw422016.N().S(`"></script>
`)
//line views/stuff.qtpl:6
	}
//line views/stuff.qtpl:6
	qw422016.N().S(`
`)
//line views/stuff.qtpl:7
}

//line views/stuff.qtpl:7
func writecommonScripts(qq422016 qtio422016.Writer) {
//line views/stuff.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/stuff.qtpl:7
	streamcommonScripts(qw422016)
//line views/stuff.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line views/stuff.qtpl:7
}

//line views/stuff.qtpl:7
func commonScripts() string {
//line views/stuff.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line views/stuff.qtpl:7
	writecommonScripts(qb422016)
//line views/stuff.qtpl:7
	qs422016 := string(qb422016.B)
//line views/stuff.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line views/stuff.qtpl:7
	return qs422016
//line views/stuff.qtpl:7
}
