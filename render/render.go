package render

import (
	"io"

	"github.com/dusanmilanov/langless/lesson"
)

// Renderer renders a lesson into a writer
type Renderer interface {
	Render(lesson.Lesson, io.Writer) error
}
