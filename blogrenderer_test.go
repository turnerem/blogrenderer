package blogrenderer_test

import (
  "bytes"
  "testing"
  "github.com/turnerem/blogrenderer"
  "github.com/approvals/go-approvals-tests"
)


func TestRender(t *testing.T) {
  var (
    aPost = blogrenderer.Post{
      Title:		"Hey",
      Body:		"Initial post",
      Description:	"This is a desc",
      Tags:		[]string{"go", "tdd"},
    }
  )

  t.Run("converts a single post to HTML", func(t *testing.T) {
    buf := bytes.Buffer{}
    err := blogrenderer.Render(&buf, aPost)

    if err != nil {
      t.Fatal(err)
    }

    approvals.VerifyString(t, buf.String())
  })
}
