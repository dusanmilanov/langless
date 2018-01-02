package render

import (
	"bytes"
	"testing"

	. "github.com/dusanmilanov/langless/lesson"
	"github.com/embroker/testify/require"
)

func Test_TextRenderer(t *testing.T) {
	r := NewTextRenderer()
	lessonText := "This is a lesson!"
	lessonTitle := "Title"
	l := Lesson{Title: lessonTitle, Description: "", Elements: []LessonElement{NewTextLessonElement(lessonText)}}

	b := bytes.NewBuffer(nil)
	err := r.Render(l, b)
	require.Nil(t, err)

	require.EqualValues(t, lessonTitle+"\n"+lessonText, b.String())

	l = Lesson{Title: lessonTitle, Description: "", Elements: []LessonElement{
		NewTextLessonElement(lessonText),
		NewTextLessonElement(lessonText),
	}}
	b = bytes.NewBuffer(nil)
	err = r.Render(l, b)
	require.Nil(t, err)
	require.EqualValues(t, lessonTitle+"\n"+lessonText+"\n"+lessonText, b.String())
}
