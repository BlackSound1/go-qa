package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define basic styles
type Styles struct {
	BorderColor       lipgloss.Color
	OutputBorderColor lipgloss.Color
	InputField        lipgloss.Style
}

// Define the data model
type Model struct {
	questions []Question
	width     int
	height    int
	index     int
	styles    *Styles
	done      bool
}

// Define a question
type Question struct {
	question  string
	answer    string
	inputType Input
}

// Creates and returns a default *Styles struct
func DefaultStyles() *Styles {
	s := new(Styles) // Create pointer to new Styles object

	// Add styles
	s.BorderColor = lipgloss.Color("36")
	s.OutputBorderColor = lipgloss.Color("25")
	s.InputField = lipgloss.
		NewStyle().
		BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).
		Width(80)

	return s
}

// Create new data *Model object
func New(questions []Question) *Model {
	styles := DefaultStyles() // Get the default styles
	return &Model{questions: questions, styles: styles}
}

// Define a new Question
func NewQuestion(question string) Question {
	return Question{question: question}
}

// Define a new short question
func newShortQuestion(q string) Question {
	question := NewQuestion(q)
	field := NewShortAnswerField()
	question.inputType = field
	return question
}

// Define a new long question
func newLongQuestion(q string) Question {
	question := NewQuestion(q)
	field := NewLongAnswerField()
	question.inputType = field
	return question
}

// Define what happens when model initially loads.
// In this case, nothing
func (m Model) Init() tea.Cmd {
	return nil
}

// Update Model info
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Create command to be passed along
	var cmd tea.Cmd

	// Get current question
	current := &m.questions[m.index]

	// Handle various updates depending on type
	switch msg := msg.(type) {

	// In case of a window resizing
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// In case of a keypress
	case tea.KeyMsg:

		// Decide what kind of keypress
		switch msg.String() {

		// Quit if user hit ctrl+c
		case tea.KeyCtrlC.String():
			return m, tea.Quit

		// If user hits enter, increment model index, and set answer field
		case tea.KeyEnter.String():
			m.CheckIfDone()

			// Set answer for current question
			current.answer = current.inputType.Value()

			// Go to next question
			m.Next()

			return m, current.inputType.Blur
		}
	}

	// Update the input type and command
	current.inputType, cmd = current.inputType.Update(msg)

	return m, cmd
}

// Change rendering of Model info
func (m Model) View() string {
	// Get current question
	current := m.questions[m.index]

	// If we're done asking question, show results
	if m.done {
		output := m.BuildOutputString()

		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			output,
		)
	}

	// Model width determined by question. If no width, no question
	if m.width == 0 {
		return "loading..."
	}

	// Render the app with a layout
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.questions[m.index].question,
			m.styles.InputField.Render(current.inputType.View()),
		),
	)
}

func main() {
	// Set up questions list
	questions := GetQuestions()

	// Create a new Model object
	m := New(questions)

	// Set up logging to a file
	f, err := tea.LogToFile("debug.log", "debug")

	// Check for errors
	if err != nil {
		log.Fatal(err)
	}

	// Close file whenever done
	defer f.Close()

	// Create new tea program using our Model m. App is fullscreen in terminal
	p := tea.NewProgram(m, tea.WithAltScreen())

	// Run app and check for errors
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
