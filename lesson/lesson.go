package lesson

// Lesson is focused on a single topic; explains it and provides exercises for it
type Lesson struct {
	Title       string
	Description string
	Elements    []LessonElement
}
