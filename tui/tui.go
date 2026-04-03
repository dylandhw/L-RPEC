package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
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

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", "space":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		}
	}

	return m, nil
}

var (
	footer = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#b4b6b8")).
		Faint(true)
	header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1d7fdb")).
		Bold(true)
	check = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1ddb26")).
		Bold(true)
	boldWhite = lipgloss.NewStyle().
			Foreground(lipgloss.Color("white")).
			Bold(true)
	bannerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1ddb26")).
			Bold(true)
)

const banner = `
 ██╗      			██████╗  ██████╗  ███████╗  ██████╗
 ██║     			██╔══██╗ ██╔══██╗  ██╔════╝ ██╔════╝
 ██║      ██████	 ██████╔╝ ███████╔╝ █████╗   ██║
 ██║     			██╔══██╗ ██╔═══╝   ██╔══╝   ██║
 ███████╗			██║  ██║ ██║       ███████╗ ╚██████╗
 ╚══════╝			╚═╝  ╚═╝ ╚═╝       ╚══════╝  ╚═════╝`

func (m model) View() tea.View {
	b := bannerStyle.Render(banner)

	s := fmt.Sprintf("\n")
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "✗"
		}
		s += fmt.Sprintf("%s %s%s%s %s\n",
			check.Render(cursor),
			boldWhite.Render("["),
			check.Render(checked),
			boldWhite.Render("]"),
			boldWhite.Render(choice),
		)
	}
	s += footer.Render("\nCommands: ↑/↓ to navigate, [enter] to toggle, ctrl+c/q to quit \n")

	content := lipgloss.JoinVertical(lipgloss.Center, b, s)

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
