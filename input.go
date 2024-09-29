package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Define an Input
type Input interface {
	Value() string
	Blur() tea.Msg
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}

// Define a short answer field
type ShortAnswerField struct {
	textInput textinput.Model
}

// Define a long answer field
type LongAnswerField struct {
	textArea textarea.Model
}

// Get the value of a short answer field
func (sa *ShortAnswerField) Value() string {
	return sa.textInput.Value()
}

// Get the value of a long answer field
func (la *LongAnswerField) Value() string {
	return la.textArea.Value()
}

// View a short answer fields text input
func (sa *ShortAnswerField) View() string {
	return sa.textInput.View()
}

// View a long answer fields text area
func (la *LongAnswerField) View() string {
	return la.textArea.View()
}

// Blur a short answer fields text input
func (sa *ShortAnswerField) Blur() tea.Msg {
	return sa.textInput.Blur
}

// Blur a long answer fields text area
func (la *LongAnswerField) Blur() tea.Msg {
	return la.textArea.Blur
}

// Update a short answer fields text input
func (sa *ShortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	sa.textInput, cmd = sa.textInput.Update(msg)
	return sa, cmd
}

// Update a long answer fields text area
func (la *LongAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	la.textArea, cmd = la.textArea.Update(msg)
	return la, cmd
}

// Create new short answer fields
func NewShortAnswerField() *ShortAnswerField {
	ti := textinput.New()
	ti.Placeholder = "Write here..."
	ti.Focus()
	return &ShortAnswerField{ti}
}

// Create new long answer field
func NewLongAnswerField() *LongAnswerField {
	ta := textarea.New()
	ta.Placeholder = "Write here..."
	ta.Focus()
	return &LongAnswerField{ta}
}
