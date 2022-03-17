package ui

import (
  "strings"

  "github.com/mrusme/planor/ui/uictx"
  "github.com/mrusme/planor/ui/navigation"

  "github.com/charmbracelet/bubbles/key"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
    Up            key.Binding
    Down          key.Binding
    Quit          key.Binding
}

var DefaultKeyMap = KeyMap{
  Up: key.NewBinding(
    key.WithKeys("k", "up"),
    key.WithHelp("↑/k", "move up"),
  ),
  Down: key.NewBinding(
    key.WithKeys("j", "down"),
    key.WithHelp("↓/j", "move down"),
  ),
  Quit: key.NewBinding(
    key.WithKeys("q", "Q"),
    key.WithHelp("q/Q", "quit"),
  ),
}


type Model struct {
  keymap        KeyMap
  nav           navigation.Model
  ctx           uictx.Ctx
}

func NewModel() Model {
  return Model{
    keymap:        DefaultKeyMap,
    nav:           navigation.NewModel(),
  }
}

func (m Model) Init() tea.Cmd {
  return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var (
    cmd        tea.Cmd
    sidebarCmd tea.Cmd
    helpCmd    tea.Cmd
    cmds       []tea.Cmd
  )

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch {
    case key.Matches(msg, m.keymap.Quit):
      cmd = tea.Quit
    }

  case tea.WindowSizeMsg:
    m.onWindowSizeMsg(msg)
  }

  cmds = append(cmds, cmd, sidebarCmd, helpCmd)
  return m, tea.Batch(cmds...)
}

func (m Model) View() (string) {
  s := strings.Builder{}
  s.WriteString(m.nav.View(m.ctx))
  s.WriteString("\n")
  mainContent := lipgloss.JoinHorizontal(
    lipgloss.Top,
    // TODO: Content
  )
  s.WriteString(mainContent)
  return s.String()
}

func (m *Model) onWindowSizeMsg(msg tea.WindowSizeMsg) {
  m.ctx.Screen[0] = msg.Width
  m.ctx.Screen[1] = msg.Height
  m.ctx.Content[1] = msg.Height - 1
}

