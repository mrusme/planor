package ui

import (
  "strings"
  "github.com/mrusme/planor/ui/navigation"
  "github.com/mrusme/planor/ui/uictx"

  "github.com/mrusme/planor/ui/views"
  "github.com/mrusme/planor/ui/views/ci"

  "github.com/charmbracelet/bubbles/key"
  tea "github.com/charmbracelet/bubbletea"
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
  views         []views.View
  ctx           *uictx.Ctx
}

func NewModel(ctx *uictx.Ctx) Model {
  m := Model{
    keymap:        DefaultKeyMap,
    ctx:           ctx,
  }

  m.nav = navigation.NewModel(m.ctx)
  m.views = append(m.views, ci.NewModel(m.ctx))

  return m
}

func (m Model) Init() tea.Cmd {
  return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  cmds := make([]tea.Cmd, 0)

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch {
    case key.Matches(msg, m.keymap.Quit):
      return m, tea.Quit
    }

  case tea.WindowSizeMsg:
    m.setSizes(msg.Width, msg.Height)
  }

  v, cmd := m.views[0].Update(msg)
  m.views[0] = v
  cmds = append(cmds, cmd)

  nav, cmd := m.nav.Update(msg)
  m.nav = nav
  cmds = append(cmds, cmd)

  return m, tea.Batch(cmds...)
}

func (m Model) View() (string) {
  s := strings.Builder{}
  s.WriteString(m.nav.View() + "\n\n")
  s.WriteString(m.views[0].View())
  return s.String()
}

func (m Model) setSizes(winWidth int, winHeight int) {
  (*m.ctx).Screen[0] = winWidth
  (*m.ctx).Screen[1] = winHeight
  m.ctx.Content[0] = m.ctx.Screen[0]
  m.ctx.Content[1] = m.ctx.Screen[1] - 4
}
