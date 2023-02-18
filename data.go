package main

type Question struct {
	Id       int
	Question string
	OptionA  string
	OptionB  string
	OptionC  string
	OptionD  string
	Answer   string
}

type QuizResponse struct {
	Id       int
	Question string
	OptionA  string
	OptionB  string
	OptionC  string
	OptionD  string
	Answer   string
	ErrorMsg string
	Status   int64
}

var questions = [][]string{
	{"How many Infinity Stones are there?", "3", "5", "6", "10", "6"},
	{"What is the only food that cannot go bad?", "Dark chocolate", "Peanut butter", "Canned tuna", "Honey", "Honey"},
	{"Which was René Magritte’s first surrealist painting?", "Not to Be Reproduced", "Personal Values", "The Lovers", "The Lost Jockey", "The Lost Jockey"},
	{"What 90s boy band member bought Myspace in 2011?", "Nick Lachey", "Justin Timberlake", "Shawn Stockman", "AJ McLean", "Justin Timberlake"},
	{"What is the most visited tourist attraction in the world?", "Eiffel Tower", "Statue of Liberty", "Great Wall of China", "Colosseum", "Eiffel Tower"},
}

func getQuestions() [][]string {
	return questions
}
