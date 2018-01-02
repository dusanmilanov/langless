package lesson

var SampleQuestion1 fillInLessonElement = fillInLessonElement{
	textLessonElement: textLessonElement{"Mi novia __ linda."},
	answerableLessonElement: answerableLessonElement{[]Answer{
		Answer{map[string]string{"es": `Third person of the verb "ser"`}, `Verb "ser"`},
	}},
}

var SampleQuestion2 fillInLessonElement = fillInLessonElement{
	textLessonElement: textLessonElement{"Yo __ alto."},
	answerableLessonElement: answerableLessonElement{[]Answer{
		Answer{map[string]string{"soy": `First person of the verb "ser"`}, `Verb "ser"`},
	}},
}

var SampleQuestion3 fillInLessonElement = fillInLessonElement{
	textLessonElement: textLessonElement{"__ novia __ alta."},
	answerableLessonElement: answerableLessonElement{[]Answer{
		Answer{map[string]string{"Mi": `My`, "Tu": "Your", "Su": "His/her"}, `A singular pronoun`},
		Answer{map[string]string{"es": `Third person of the verb "ser"`}, `Verb "ser"`},
	}},
}

var Lesson1 Lesson = Lesson{
	Title:       `Verb "ser" exercises`,
	Description: `Set of exercies for present tense of the verb "ser"`,
	Elements: []LessonElement{
		SampleQuestion1, SampleQuestion2, SampleQuestion3,
	},
}
