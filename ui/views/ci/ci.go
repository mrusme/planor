package ci

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
  "github.com/charmbracelet/bubbles/spinner"
	"github.com/mrusme/planor/ui/uictx"
)

type KeyMap struct {
    Refresh       key.Binding
}

var DefaultKeyMap = KeyMap{
  Refresh: key.NewBinding(
    key.WithKeys("r", "R"),
    key.WithHelp("r/R", "refresh"),
  ),
}

type Model struct {
  keymap          KeyMap
  list            list.Model
  items           []list.Item
  spinner         spinner.Model
  ctx             *uictx.Ctx

  refreshing      bool
}

func (m Model) Init() tea.Cmd {
  return m.spinner.Tick
}

func NewModel(ctx *uictx.Ctx) (Model) {
  m := Model{
    keymap:        DefaultKeyMap,
    refreshing:    false,
  }

  m.spinner = spinner.New()
  m.spinner.Spinner = spinner.Dot
  m.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

  m.list = list.New(m.items, list.NewDefaultDelegate(), 0, 0)
  m.list.Title = "Pipelines"
  m.ctx = ctx

  return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmds []tea.Cmd

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch {
    case key.Matches(msg, m.keymap.Refresh):
      m.refreshing = true
      cmds = append(cmds, m.refresh())
    }

  case tea.WindowSizeMsg:
    m.list.SetSize(
      m.ctx.Content[0],
      m.ctx.Content[1],
    )

  case spinner.TickMsg:
    var cmd tea.Cmd
    m.spinner, cmd = m.spinner.Update(msg)
    cmds = append(cmds, cmd)
    // return m, cmd

  case []list.Item:
    m.items = msg
    m.list.SetItems(m.items)
    m.refreshing = false
  }

  var cmd tea.Cmd
  m.list, cmd = m.list.Update(msg)
  cmds = append(cmds, cmd)

  if m.refreshing == true {
    cmds = append(cmds, m.spinner.Tick)
  }

  return m, tea.Batch(cmds...)
}

func (m Model) View() (string) {
  var view string

  if m.refreshing == true {
    view = lipgloss.JoinHorizontal(
      lipgloss.Top,
      m.list.View(),
      m.spinner.View(),
    )
  } else {
    view = lipgloss.JoinHorizontal(
      lipgloss.Top,
      m.list.View(),
    )
  }

  return view
}

func (m *Model) refresh() (tea.Cmd) {
  return func () (tea.Msg) {
    var items []list.Item

    pipelines, err :=  (*m.ctx.Cloud).ListPipelines()
    if err != nil {
      fmt.Printf("%s", err) // TODO: Implement error message
    }
    for _, pipeline := range pipelines {
      items = append(items, pipeline)
    }

    return items
  }
}

