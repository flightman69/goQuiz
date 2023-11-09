package main

import (
	"quiz/cmd/quiz"
)

func main() {
<<<<<<< HEAD
	/* This opens the Questions file */

	problemFile, testTime, shuffle := utils.GetArguments()
	file, err := os.Open(*problemFile)
	if err != nil {
		utils.ExitWithMessage(fmt.Sprintf("Failed to open the CSV file: %s\n", *problemFile), 1)
	} // Could use json file instead of CSV file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		utils.ExitWithMessage("Could not parse provided csv", 1)
	}
	problems := parsers.ParseLines(lines)
	if *shuffle {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	// This Prompt the user with questions

	marks, err := utils.QuestionUser(problems, *testTime, utils.ConsoleReader, utils.QuizTimer)
	fmt.Printf("You got %d/%d correct!\n", marks, len(problems))
	if err != nil {
		userErrs, _ := err.(utils.QuizEvaluation)
		userErrs.PrintErrors()

		if userErrs.Unattempted {
			userErrs.PrintUnattempted()
		}
	}
=======
	quiz.Execute()
>>>>>>> 87151503e130aed105c6e0e134a4c34dceacd7a9
}
