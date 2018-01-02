package render

import (
	"io"

	"github.com/dusanmilanov/langless/lesson"
)

type textRenderer struct{}

// NewTextRenderer returns a simple text renderer
func NewTextRenderer() Renderer {
	return textRenderer{}
}

type textWriter struct {
	w   io.Writer
	err error
}

func (t *textWriter) writeString(s string) {
	if t.err != nil {
		return
	}
	_, t.err = io.WriteString(t.w, s)
}

func (t textRenderer) Render(l lesson.Lesson, w io.Writer) error {
	tw := textWriter{w, nil}

	tw.writeString(l.Title + "\n")
	for i, elem := range l.Elements {
		if i > 0 {
			tw.writeString("\n")
		}
		tw.writeString(elem.GetText())
	}
	return tw.err
}
