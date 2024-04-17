package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var appLayout = lipgloss.NewStyle()

type Msg struct {
	username string
	body     string
}

func (m Msg) FilterValue() string {
	return m.body
}

func (m Msg) Title() string {
	return m.body
}

func (m Msg) Description() string {
	return m.username
}

type Model struct {
	chat  list.Model
	input textinput.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.chat.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		if key.Matches(msg, key.NewBinding(key.WithKeys("enter"))) {
			items := m.chat.Items()
			items = append(items, Msg{
				username: "kubre.in",
				body:     m.input.Value(),
			})
			m.chat.SetItems(items)
			m.input.SetValue("")
		}
	}

	var cmd tea.Cmd
	m.input.Focus()
	m.input, cmd = m.input.Update(msg)
	cmd = textinput.Blink
	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf("%s\n%s", m.chat.View(), m.input.View())
}

func main() {
	chat := []list.Item{
		Msg{username: "kubre.in", body: "Hi, Hello!"},
		Msg{username: "Bot", body: "Hi, kubre.in"},
		Msg{username: "kubre.in", body: "Are you really a bot?"},
		Msg{username: "Bot", body: "Yes, I'm a LLM AI bot"},
		Msg{username: "kubre.in", body: "That is cool"},
	}
	m := Model{
		chat:  list.New(chat, list.NewDefaultDelegate(), 0, 0),
		input: textinput.New(),
	}
	m.chat.Title = "AI Chat"
	m.chat.SetFilteringEnabled(false)
	m.chat.SetShowHelp(false)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
