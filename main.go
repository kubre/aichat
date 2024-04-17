package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

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
	chat list.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.chat.SetSize(msg.Width, msg.Height)
	}

	var cmd tea.Cmd
	m.chat, cmd = m.chat.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.chat.View()
}

func main() {
	chat := []list.Item{
		Msg{username: "kubre.in", body: "Hi, Hello!"},
		Msg{username: "Bot", body: "Hi, kubre.in"},
		Msg{username: "kubre.in", body: "Are you really a bot?"},
		Msg{username: "Bot", body: "Yes, I'm a LLM AI bot"},
		Msg{username: "kubre.in", body: "That is cool"},
	}
	m := Model{chat: list.New(chat, list.NewDefaultDelegate(), 0, 0)}
	m.chat.Title = "AI Chat"

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
