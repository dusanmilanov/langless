package lesson

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Grading(t *testing.T) {
	_, err := SampleQuestion1.Grade([]string{})
	require.NotNil(t, err)

	grade, err := SampleQuestion1.Grade([]string{"is"})
	require.Nil(t, err)
	require.False(t, grade.IsCorrect())
	require.Len(t, grade.Individual, 1)
	require.Len(t, grade.Hints, 1)
	require.False(t, grade.Individual[0])
	require.EqualValues(t, SampleQuestion1.answers[0].Hint, grade.Hints[0])

	grade, err = SampleQuestion1.Grade([]string{"es"})
	require.Nil(t, err)
	require.True(t, grade.IsCorrect())
	require.Len(t, grade.Individual, 1)
	require.Len(t, grade.Hints, 1)
	require.True(t, grade.Individual[0])
	require.EqualValues(t, SampleQuestion1.answers[0].Accepted["es"], grade.Hints[0])

	grade, err = SampleQuestion3.Grade([]string{"?", "es"})
	require.Nil(t, err)
	require.False(t, grade.IsCorrect())
	require.Len(t, grade.Individual, 2)
	require.Len(t, grade.Hints, 2)
	require.False(t, grade.Individual[0])
	require.True(t, grade.Individual[1])
	require.EqualValues(t, SampleQuestion3.answers[0].Hint, grade.Hints[0])
	require.EqualValues(t, SampleQuestion3.answers[1].Accepted["es"], grade.Hints[1])
}
