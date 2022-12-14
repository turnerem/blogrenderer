package blogrenderer

import (
  "io"
  "html/template"
  "embed"
)

var (
  //go:embed "templates/*"
  postTemplates embed.FS
)

type Post struct {
  Title		string
  Body		string
  Description	string
  Tags		[]string
}

func Render(w io.Writer, p Post) error {
  templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

  if err != nil {
    return err
  }

  if err := templ.Execute(w, p); err != nil {
    return err
  }

  return nil
}
