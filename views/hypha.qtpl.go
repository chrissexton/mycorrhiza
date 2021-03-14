// Code generated by qtc from "hypha.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/hypha.qtpl:1
package views

//line views/hypha.qtpl:1
import "path/filepath"

//line views/hypha.qtpl:2
import "strings"

//line views/hypha.qtpl:3
import "github.com/bouncepaw/mycorrhiza/hyphae"

//line views/hypha.qtpl:4
import "github.com/bouncepaw/mycorrhiza/util"

//line views/hypha.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/hypha.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/hypha.qtpl:6
func StreamNaviTitleHTML(qw422016 *qt422016.Writer, h *hyphae.Hypha) {
//line views/hypha.qtpl:6
	qw422016.N().S(`
`)
//line views/hypha.qtpl:8
	var (
		prevAcc = "/hypha/"
		parts   = strings.Split(h.Name, "/")
	)

//line views/hypha.qtpl:12
	qw422016.N().S(`
<h1 class="navi-title">
`)
//line views/hypha.qtpl:14
	qw422016.N().S(`<a href="/hypha/`)
//line views/hypha.qtpl:15
	qw422016.E().S(util.HomePage)
//line views/hypha.qtpl:15
	qw422016.N().S(`">`)
//line views/hypha.qtpl:16
	qw422016.N().S(util.SiteNavIcon)
//line views/hypha.qtpl:16
	qw422016.N().S(`<span aria-hidden="true" class="navi-title__colon">:</span></a>`)
//line views/hypha.qtpl:20
	for i, part := range parts {
//line views/hypha.qtpl:21
		if i > 0 {
//line views/hypha.qtpl:21
			qw422016.N().S(`<span aria-hidden="true" class="navi-title__separator">/</span>`)
//line views/hypha.qtpl:23
		}
//line views/hypha.qtpl:23
		qw422016.N().S(`<a href="`)
//line views/hypha.qtpl:25
		qw422016.E().S(prevAcc + part)
//line views/hypha.qtpl:25
		qw422016.N().S(`"rel="`)
//line views/hypha.qtpl:26
		if i == len(parts)-1 {
//line views/hypha.qtpl:26
			qw422016.N().S(`bookmark`)
//line views/hypha.qtpl:26
		} else {
//line views/hypha.qtpl:26
			qw422016.N().S(`up`)
//line views/hypha.qtpl:26
		}
//line views/hypha.qtpl:26
		qw422016.N().S(`">`)
//line views/hypha.qtpl:27
		qw422016.N().S(util.BeautifulName(part))
//line views/hypha.qtpl:27
		qw422016.N().S(`</a>`)
//line views/hypha.qtpl:29
		prevAcc += part + "/"

//line views/hypha.qtpl:30
	}
//line views/hypha.qtpl:31
	qw422016.N().S(`
</h1>
`)
//line views/hypha.qtpl:33
}

//line views/hypha.qtpl:33
func WriteNaviTitleHTML(qq422016 qtio422016.Writer, h *hyphae.Hypha) {
//line views/hypha.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/hypha.qtpl:33
	StreamNaviTitleHTML(qw422016, h)
//line views/hypha.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line views/hypha.qtpl:33
}

//line views/hypha.qtpl:33
func NaviTitleHTML(h *hyphae.Hypha) string {
//line views/hypha.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/hypha.qtpl:33
	WriteNaviTitleHTML(qb422016, h)
//line views/hypha.qtpl:33
	qs422016 := string(qb422016.B)
//line views/hypha.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/hypha.qtpl:33
	return qs422016
//line views/hypha.qtpl:33
}

//line views/hypha.qtpl:35
func StreamAttachmentHTML(qw422016 *qt422016.Writer, h *hyphae.Hypha) {
//line views/hypha.qtpl:35
	qw422016.N().S(`
	`)
//line views/hypha.qtpl:36
	switch filepath.Ext(h.BinaryPath) {
//line views/hypha.qtpl:38
	case ".jpg", ".gif", ".png", ".webp", ".svg", ".ico":
//line views/hypha.qtpl:38
		qw422016.N().S(`
	<div class="binary-container binary-container_with-img">
		<a href="/binary/`)
//line views/hypha.qtpl:40
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:40
		qw422016.N().S(`"><img src="/binary/`)
//line views/hypha.qtpl:40
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:40
		qw422016.N().S(`"/></a>
	</div>

	`)
//line views/hypha.qtpl:43
	case ".ogg", ".webm", ".mp4":
//line views/hypha.qtpl:43
		qw422016.N().S(`
	<div class="binary-container binary-container_with-video">
		<video controls>
			<source src="/binary/`)
//line views/hypha.qtpl:46
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:46
		qw422016.N().S(`"/>
			<p>Your browser does not support video. <a href="/binary/`)
//line views/hypha.qtpl:47
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:47
		qw422016.N().S(`">Download video</a></p>
		</video>
	</div>

	`)
//line views/hypha.qtpl:51
	case ".mp3":
//line views/hypha.qtpl:51
		qw422016.N().S(`
	<div class="binary-container binary-container_with-audio">
		<audio controls>
			<source src="/binary/`)
//line views/hypha.qtpl:54
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:54
		qw422016.N().S(`"/>
			<p>Your browser does not support audio. <a href="/binary/`)
//line views/hypha.qtpl:55
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:55
		qw422016.N().S(`">Download audio</a></p>
		</audio>
	</div>

	`)
//line views/hypha.qtpl:59
	default:
//line views/hypha.qtpl:59
		qw422016.N().S(`
	<div class="binary-container binary-container_with-nothing">
		<p><a href="/binary/`)
//line views/hypha.qtpl:61
		qw422016.N().S(h.Name)
//line views/hypha.qtpl:61
		qw422016.N().S(`">Download media</a></p>
	</div>
`)
//line views/hypha.qtpl:63
	}
//line views/hypha.qtpl:63
	qw422016.N().S(`
`)
//line views/hypha.qtpl:64
}

//line views/hypha.qtpl:64
func WriteAttachmentHTML(qq422016 qtio422016.Writer, h *hyphae.Hypha) {
//line views/hypha.qtpl:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/hypha.qtpl:64
	StreamAttachmentHTML(qw422016, h)
//line views/hypha.qtpl:64
	qt422016.ReleaseWriter(qw422016)
//line views/hypha.qtpl:64
}

//line views/hypha.qtpl:64
func AttachmentHTML(h *hyphae.Hypha) string {
//line views/hypha.qtpl:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/hypha.qtpl:64
	WriteAttachmentHTML(qb422016, h)
//line views/hypha.qtpl:64
	qs422016 := string(qb422016.B)
//line views/hypha.qtpl:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/hypha.qtpl:64
	return qs422016
//line views/hypha.qtpl:64
}
