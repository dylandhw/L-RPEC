package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	menuState state = iota
	howItWorksState
)

type model struct {
	state    state
	choices  []string
	cursor   int
	selected map[int]struct{}
	width    int
	height   int
}

func initialModel() model {
	return model{
		choices: []string{"Run Vegeta Stress Testing", "Grafana Dashboard        ", "Rate Limiting            ", "How It Works             "},

		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter":
			_, ok := m.selected[m.cursor]
			choice := m.choices[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
			if choice == "How It Works             " {
				m.state = howItWorksState
				return m, nil
			}
		case "b":
			if m.state != menuState {
				m.state = menuState
				return m, nil
			}
		}
	}

	return m, nil
}

const banner = `
 ██╗      		   ██████╗  ██████  ███████╗  ██████╗
 ██║     		   ██╔══██╗ ██╔══██  ██╔════╝ ██╔════╝
 ██║      ██████    ██████╔╝ ██████║  █████╗   ██║
 ██║     		   ██╔══██╗ ██╔═══╝  ██╔══╝   ██║
 ███████╗		   ██║  ██║ ██║      ███████╗ ╚██████╗
 ╚══════╝		   ╚═╝  ╚═╝ ╚═╝      ╚══════╝  ╚═════╝`

var (
	/* MENU STATE */
	footer = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#b4b6b8")).
		Faint(true)
	header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1d7fdb")).
		Bold(true)
	check = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fcba03")).
		Bold(true)
	boldWhite = lipgloss.NewStyle().
			Foreground(lipgloss.Color("white")).
			Bold(true)
	bannerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1ddb26")).
			Bold(true)
)

func (m model) View() tea.View {
	switch m.state {
	case menuState:
		return m.renderMenu()
	case howItWorksState:
		return m.renderHowItWorks()
	}
	return m.renderMenu()
}

func (m model) renderMenu() tea.View {
	b := bannerStyle.Render(banner)
	var s strings.Builder
	s.WriteString("\n")
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "»"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "✗"
		}
		s.WriteString(fmt.Sprintf("%s %s%s%s %s\n",
			check.Render(cursor),
			boldWhite.Render("["),
			check.Render(checked),
			boldWhite.Render("]"),
			boldWhite.Render(choice),
		))
	}
	s.WriteString(footer.Render("\nCommands: ↑/↓ to navigate, [enter] to toggle, b to go back, ctrl+c/q to quit \n"))

	content := lipgloss.JoinVertical(lipgloss.Center, b, s.String())

	return tea.NewView(lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	))
}

var (
	/* HOW IT WORKS */
	paragraph = lipgloss.NewStyle().
		Width(60).
		Foreground(lipgloss.Color("#b4b6b8")).
		Bold(true)
)

func (m model) renderHowItWorks() tea.View {
	b := bannerStyle.Render(banner)
	var s strings.Builder
	s.WriteString("\n")

	text1 := paragraph.Render("\nL-RPEC is a lightweight reverse proxy that routes incoming requests based on configuration, caches responses in memory, and signs outbound requests using HMAC. It acts as a simplified edge layer, similar to a CDN, allowing you to experiment with caching strategies, routing logic, and request security.")
	text2 := paragraph.Render("\nEssentially a toy Cloudlfare CDN. Largely made to satisfy my curiosity about infrastructure. The core functionality is built entirely with the Go stdlib, you can find a more technical insight below.")
	s.WriteString(footer.Render("\nCommands: ↑/↓ to navigate, [enter] to toggle, b to go back, ctrl+c/q to quit \n"))
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		b,
		text1,
		text2,
		s.String(),
	)

	return tea.NewView(lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	))
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
