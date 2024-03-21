package html

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

func CreateBareHTMLTemplate(templateName, content string, data any) string {
	var s strings.Builder
	w := io.MultiWriter(&s)

	c := fmt.Sprintf(`{{define %q}} %v {{end}}`, templateName, content)

	t := template.Must(template.New(templateName).Parse(c))

	err := t.ExecuteTemplate(w, templateName, data)
	if err != nil {
		panic(err)
	}

	return s.String()
}
