package quiz_test

import (
	"testing"

	"github.com/janshercs/quizzer/pkg/quiz"
	"github.com/stretchr/testify/assert"
)

func TestGetQuestionFromStringSlice(t *testing.T) {
	q := quiz.GetQuestionFromStringSlice([]string{"I really like?", "Ipsum", "Lorem", "Possum", "Dolor", "1", "Not possum."})
	assert.Equal(t, quiz.Question{
		Query: "I really like?",
		Options: []quiz.Option{
			{
				ID:    "1",
				Value: "Ipsum",
			},
			{
				ID:    "2",
				Value: "Lorem",
			},
			{
				ID:    "3",
				Value: "Possum",
			},
			{
				ID:    "4",
				Value: "Dolor",
			},
		},
		Answer: quiz.Option{
			ID:    "1",
			Value: "Ipsum",
		},
		Comment: "Not possum.",
	}, q)
}

func TestConvertSheetValueToRows(t *testing.T) {
	rows := quiz.ConvertSheetValueToRows([][]interface{}{
		{"I really like?", "Ipsum", "Lorem", "Possum", "Dolor", "4", "Not possum."},
		{"Which of the following is a type of?", "Lorem", "Ipsum", "Dolor", "Sit", "2", "Not a type."},
	})

	assert.Equal(
		t,
		[][]string{
			{"I really like?", "Ipsum", "Lorem", "Possum", "Dolor", "4", "Not possum."},
			{"Which of the following is a type of?", "Lorem", "Ipsum", "Dolor", "Sit", "2", "Not a type."},
		},
		rows,
	)
}
