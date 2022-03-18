package ci

import (
  "fmt"
  "math"

  "github.com/charmbracelet/bubbles/key"
  "github.com/charmbracelet/bubbles/list"
  "github.com/charmbracelet/bubbles/viewport"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
  "github.com/mrusme/planor/nori/models"
  "github.com/mrusme/planor/ui/uictx"
)

var (
  viewportStyle = lipgloss.NewStyle().Margin(0, 0, 0, 0)
)

type KeyMap struct {
    Refresh       key.Binding
    Select        key.Binding
}

var DefaultKeyMap = KeyMap{
  Refresh: key.NewBinding(
    key.WithKeys("r", "R"),
    key.WithHelp("r/R", "refresh"),
  ),
  Select: key.NewBinding(
    key.WithKeys("enter"),
    key.WithHelp("enter", "select"),
  ),
}

type Model struct {
  keymap          KeyMap
  list            list.Model
  items           []list.Item
  viewport        viewport.Model
  ctx             *uictx.Ctx

  refreshing      bool
}

func (m Model) Init() tea.Cmd {
  return nil
}

func NewModel(ctx *uictx.Ctx) (Model) {
  m := Model{
    keymap:        DefaultKeyMap,
  }

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
      m.ctx.Loading = true
      cmds = append(cmds, m.refresh())

    case key.Matches(msg, m.keymap.Select):
      i, ok := m.list.SelectedItem().(models.Pipeline)
      if ok {
        m.viewport.SetContent(m.renderViewport(&i))
        return m, nil
      }
    }

  case tea.WindowSizeMsg:
    m.list.SetSize(
      int(math.Floor(float64(m.ctx.Content[0]) / 3.0)),
      m.ctx.Content[1],
    )
    m.viewport.Width = int(math.Ceil(float64(m.ctx.Content[0]) / 3.0 * 2.0))
    m.viewport.Height = m.ctx.Content[1]
    cmds = append(cmds, viewport.Sync(m.viewport))

  case []list.Item:
    m.items = msg
    m.list.SetItems(m.items)
    m.ctx.Loading = false
  }

  var cmd tea.Cmd
  m.list, cmd = m.list.Update(msg)
  cmds = append(cmds, cmd)

  m.viewport, cmd = m.viewport.Update(msg)
  cmds = append(cmds, cmd)

  return m, tea.Batch(cmds...)
}

func (m Model) View() (string) {
  var view string

  view = lipgloss.JoinHorizontal(
    lipgloss.Top,
    m.list.View(),
    m.viewport.View(),
  )

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

func (m *Model) renderViewport(pipeline *models.Pipeline) (string) {
  var vp string = ""

  vp = fmt.Sprintf(
    "%s\n\nUpdated %s\n",
    pipeline.Name,
    pipeline.UpdatedAt.String(),
  )

  for _, stage := range pipeline.Stages {
    vp = fmt.Sprintf(
      "%s\n  %s\n  Status: %s\n",
      vp,
      stage.Name,
      stage.Status,
    )

    for _, action := range stage.Actions {
      prct := fmt.Sprintf("%d%%", action.PercentComplete)
      if prct == "0%" {
        prct = ""
      }

      vp = fmt.Sprintf(
        "%s\n    %s\n    Updated: %s\n    Status: %s %s\n    %s\n",
        vp,
        action.Name,
        action.UpdatedAt.String(),
        action.Status,
        prct,
        action.Summary,
      )
    }
  }

  return viewportStyle.Render(vp)
}
