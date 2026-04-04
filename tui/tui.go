package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	menuState state = iota
	howItWorksState
	vegetaStressTestState
)

type model struct {
	// generic frame
	state    state
	choices  []string
	cursor   int
	selected map[int]struct{}
	width    int
	height   int

	// progress bar
	progress  float64
	duration  time.Duration
	testing   bool
	completed bool
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

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(16*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		if m.testing {
			totalTicks := m.duration.Seconds() * (1000.0 / 16.0) // ticks per second
			m.progress += 1.0 / totalTicks

			if m.progress >= 1.0 {
				m.progress = 0.0
				m.testing = false
				m.completed = true
				return m, nil
			}
			return m, tickCmd()
		}

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
			} else if choice == "Run Vegeta Stress Testing" {
				m.state = vegetaStressTestState
				m.testing = true
				m.completed = false
				m.duration = 10 * time.Second
				return m, tickCmd()
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
		Foreground(lipgloss.Color("#fcba03")). // #1d7fdb
		Bold(true)
	check = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fcba03")).
		Bold(true)
	boldWhite = lipgloss.NewStyle().
			Foreground(lipgloss.Color("white")).
			Bold(true)
	bannerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#fcba03")).
			Bold(true)
)

func (m model) View() tea.View {
	switch m.state {
	case menuState:
		return m.renderMenu()
	case howItWorksState:
		return m.renderHowItWorks()
	case vegetaStressTestState:
		return m.renderVegetaStressTest()
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

	link = lipgloss.NewStyle().
		Width(60).
		Underline(false).
		Bold(true).
		Foreground(lipgloss.Color("#2754e8")).
		Faint(true)
)

func (m model) renderHowItWorks() tea.View {
	b := bannerStyle.Render(banner)
	var s strings.Builder
	s.WriteString("\n")

	text1 := paragraph.Render("\nL-RPEC is a lightweight reverse proxy that routes incoming requests based on configuration, caches responses in memory, and signs outbound requests using HMAC. It acts as a simplified edge layer, similar to a CDN, allowing you to experiment with caching strategies, routing logic, and request security.")
	text2 := paragraph.Render("\nEssentially a toy Cloudlfare CDN. Largely made to satisfy my curiosity about infrastructure. The core functionality is built entirely with the Go stdlib, you can find a more technical insight below.")
	text3 := paragraph.Render("\nBelow are some links that I found helpful when building this.")
	link1 := link.Render("\n- What is a reverse proxy?")
	link2 := link.Render("- HTTP Caching")
	link3 := link.Render("- How CDNs Work")
	link4 := link.Render("- Signing Requests")

	s.WriteString(footer.Render("\nCommands: ↑/↓ to navigate, [enter] to toggle, b to go back, ctrl+c/q to quit \n"))
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		b,
		text1,
		text2,
		text3,
		link1,
		link2,
		link3,
		link4,
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

var (
	progressBarFill  = lipgloss.NewStyle().Foreground(lipgloss.Color("#fcba03")).Bold(true)
	progressBarEmpty = lipgloss.NewStyle().Foreground(lipgloss.Color("#333333"))
	progressLabel    = lipgloss.NewStyle().Foreground(lipgloss.Color("#b4b6b8")).Bold(true)
)

func (m model) renderVegetaStressTest() tea.View {
	b := bannerStyle.Render(banner)
	var s strings.Builder

	barWidth := 50
	filled := int(m.progress * float64(barWidth))
	empty := barWidth - filled

	if !m.completed {
		bar := progressBarFill.Render(strings.Repeat("█", filled)) +
			progressBarEmpty.Render(strings.Repeat("░", empty))

		pct := int(m.progress * 100)
		label := progressLabel.Render(fmt.Sprintf("\nRunning stress test... %d%%\n", pct))

		status := ""
		if !m.testing && m.progress == 0.0 {
			status = progressLabel.Render("Stress test completed\n")
		}

		s.WriteString(label)
		s.WriteString("\n  " + bar + "\n\n")
		s.WriteString(status)
		s.WriteString(footer.Render("\nCommands: b to go back, ctrl+c/q to quit\n"))

	}
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		b,
		s.String(),
	)
	/*
	 SO WE HAVE A ROUGH IDEA OF WHERE WE ARE AT.... IT DISABLES THE BAR UPON COMPLETION (YAY!)
	 BUT IT ALSO DISABLES THE STATUS MESSAGES AND THE CONTROLS FOOTER...SO LET'S FIX THIS TOMORROW
	*/
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
