package store

import (
	"github.com/dusanmilanov/langless/lesson"
)

// PackageName provides type safety when accessing package maps
type PackageName string

// LessonPackage is an ordered list of lessons with a name
type LessonPackage struct {
	ID      EntityID
	Name    PackageName
	Lessons []PackageLesson
}

// PackageLesson is a wrapper around Lesson that provides additional information necessary for storing it and using it inside a package
type PackageLesson struct {
	lesson.Lesson
	ID     EntityID
	Group  string
	Number string
}

// LessonStore provides storage for Lessons and LessonPackages
type LessonStore interface {
	GetPackages() map[PackageName]LessonPackage
}
