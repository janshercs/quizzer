package quiz

type Quiz struct {
	Questions Questions
}

type Questions []Question

type Question struct {
	Query   string
	Options Options
	Answer  Option
	Comment string
}

type Options []Option

type Option struct {
	ID    string
	Value string
}
