package main

// Define and return the questions for the app
func GetQuestions() []Question {
	return []Question{
		newShortQuestion("What is your name?"),
		newShortQuestion("What is your favourite editor?"),
		newLongQuestion("What is your favourite quote?"),
		newLongQuestion("Why do you exist?"),
	}
}
