package lesson

import (
	"fmt"
)

// LessonElement is a part of the lesson that provides text or exercises
type LessonElement interface {
	GetText() string
}

type textLessonElement struct {
	text string
}

func (l textLessonElement) GetText() string {
	return l.text
}

// NewTextLessonElement creates a text LessonElement; can contain CommonMark, HTML, etc
func NewTextLessonElement(text string) LessonElement {
	return textLessonElement{text}
}

// Grade is result of grading a single LessonElement
// Provides a boolean result and a hint for each question in the graded LessonElement
type Grade struct {
	Individual []bool
	Hints      []string
}

// IsCorrect returns true if all the questions on a LessonElement have been answered correctly
func (g Grade) IsCorrect() bool {
	for _, b := range g.Individual {
		if !b {
			return false
		}
	}
	return true
}

// Answerable is a LessonElement that can be graded
type Answerable interface {
	Grade(answers []string) (Grade, error)
}

// Answer provides accepted values for a single question in a LessonElement
type Answer struct {
	Accepted map[string]string // Accepted answers with specific reason why
	Hint     string
}

// IndividualGrade evaluates an answer against the list of accepted values
func (a Answer) IndividualGrade(answer string) (accepted bool, hint string) {
	reason, ok := a.Accepted[answer]
	if ok {
		return true, reason
	}
	return false, a.Hint
}

type answerableLessonElement struct {
	answers []Answer
}

func errAnswerNumberMismatch(numRequired, numProvided int) error {
	return fmt.Errorf("Answer number mismatch: required %d, provided: %d", numRequired, numProvided)
}

// Grade evaluates answers for all questions in an answerableLessonElement
func (e answerableLessonElement) Grade(answers []string) (Grade, error) {
	g := Grade{}

	if len(answers) != len(e.answers) {
		return g, errAnswerNumberMismatch(len(e.answers), len(answers))
	}

	for i, answer := range answers {
		accepted, hint := e.answers[i].IndividualGrade(answer)
		g.Individual = append(g.Individual, accepted)
		g.Hints = append(g.Hints, hint)
	}
	return g, nil
}

type fillInLessonElement struct {
	textLessonElement
	answerableLessonElement
}
