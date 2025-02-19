package data

type Question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"` // Index of correct answer
}

func GetQuestions() []Question {
	return []Question{
		{"What is the zero value of an int in Go?", []string{"0", "nil", "undefined", "-1"}, 0},
		{"Which keyword is used to create a goroutine?", []string{"go", "thread", "routine", "async"}, 0},
		{"How do you declare a constant in Go?", []string{"var", "const", "let", "final"}, 1},
		{"What is the default visibility of functions in Go?", []string{"Public", "Private", "Protected", "None"}, 1},
		{"Which data type is used for Unicode characters in Go?", []string{"string", "rune", "char", "byte"}, 1},
		{"Which package is used for I/O operations in Go?", []string{"fmt", "io", "os", "log"}, 1},
		{"What does the `defer` keyword do in Go?", []string{"Runs the function immediately", "Skips execution", "Delays execution until the surrounding function returns", "None"}, 2},
		{"What is the purpose of the blank identifier `_` in Go?", []string{"Ignore values", "Declare variables", "Import packages", "Error handling"}, 0},
		{"Which of the following is a correct slice declaration?", []string{"arr := []int{}", "arr := [5]int{}", "arr := int[]{}", "arr := slice<int>()"}, 0},
		{"How do you handle errors in Go?", []string{"try-catch", "if err != nil", "panic-recover", "None"}, 1},
	}
}
