package quiz

import "strconv"

const (
	maxOpts = 4

	queryPos    = 0
	firstOptPos = 1
	lastOptPos  = 4
	answerPos   = 5
	commentPos  = 6
)

func GetQuizFromSheetValues(v [][]interface{}) Quiz {
	questions := make(Questions, 0, len(v))
	for _, q := range convertSheetValueToRows(v) {
		questions = append(questions, GetQuestionFromStringSlice(q))
	}
	return Quiz{Questions: questions}
}

func convertSheetValueToRows(v [][]interface{}) [][]string {
	data := make([][]string, 0, len(v))
	for _, r := range v {
		row := make([]string, 0, len(r))
		for _, value := range r {
			row = append(row, value.(string))
		}
		data = append(data, row)
	}
	return data
}

func GetQuestionFromStringSlice(s []string) Question {
	opts := make([]Option, 0, maxOpts)
	var answer Option
	for i := firstOptPos; i <= lastOptPos; i++ {
		id := strconv.Itoa(i)
		opt := Option{
			ID:    id,
			Value: s[i],
		}
		opts = append(opts, opt)

		if id == s[answerPos] {
			answer = opt
		}
	}

	return Question{
		Query:   s[queryPos],
		Options: opts,
		Answer:  answer,
		Comment: s[commentPos],
	}
}
