package email

import "html/template"

var RsvpEmailTmpl = template.Must(template.New("rsvp-email").Parse(`
<h2>Ny OSNING</h2>
<p><strong>E-post:</strong> {{.Email}}</p>
<p><strong>Kommer:</strong> {{if eq .Coming "yes"}}Ja{{else}}Nej{{end}}</p>
<p><strong>Antal:</strong> {{len .People}}</p>
{{if .People}}
  <h3>Personer</h3>
  <ul>
    {{range .People}}<li>{{.}}</li>{{end}}
  </ul>
{{end}}
{{if .Message}}
  <h3>Meddelande</h3>
  <p>{{.Message}}</p>
{{end}}
`))
