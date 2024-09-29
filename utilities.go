package main

import "fmt"

func (m Model) BuildOutputString() string {
	var output string

	// Build up an output string
	for _, q := range m.questions {
		output += fmt.Sprintf("%s: %s\n", q.question, q.answer)
	}

	return output
}

// Set done if our question is the last
func (m *Model) CheckIfDone() {
	if m.index == len(m.questions)-1 {
		m.done = true
	}
}

// Bring us to the next question
func (m *Model) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	}
}
