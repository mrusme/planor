package logs

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
  viewportStyle = lipgloss.NewStyle().
    Margin(0, 0, 0, 0).
    Padding(1, 1).
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("#874BFD")).
    BorderTop(true).
    BorderLeft(true).
    BorderRight(true).
    BorderBottom(true)
)

type KeyMap struct {
    Refresh       key.Binding
    Select        key.Binding
    SwitchFocus   key.Binding
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
  SwitchFocus: key.NewBinding(
    key.WithKeys("tab"),
    key.WithHelp("tab", "switch focus"),
  ),
}

type Model struct {
  keymap          KeyMap
  listGroups      list.Model
  items           []list.Item
  viewport        viewport.Model
  ctx             *uictx.Ctx

  focused         int
  focusables      [2]tea.Model
}

func (m Model) Init() tea.Cmd {
  return nil
}

func NewModel(ctx *uictx.Ctx) (Model) {
  m := Model{
    keymap:        DefaultKeyMap,
    focused:       0,
  }

  m.listGroups = list.New(m.items, list.NewDefaultDelegate(), 0, 0)
  m.listGroups.Title = "Groups"
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

    case key.Matches(msg, m.keymap.SwitchFocus):
      m.focused++
      if m.focused >= len(m.focusables) {
        m.focused = 0
      }
      // return m, nil

    case key.Matches(msg, m.keymap.Select):
      group, ok := m.listGroups.SelectedItem().(models.LogGroup)
      if ok {
        var items []list.Item
        for _, stream := range group.Streams {
          items = append(items, stream)
        }
        m.listGroups.SetItems(items)
        m.listGroups.Select(0)
        return m, nil
      }

      stream, ok := m.listGroups.SelectedItem().(models.LogStream)
      if ok {
        m.viewport.SetContent(m.renderViewport(&stream))
        return m, nil
      }
    }

  case tea.WindowSizeMsg:
    listGroupsWidth := int(math.Floor(float64(m.ctx.Content[0]) / 4.0))
    viewportWidth := m.ctx.Content[0] - listGroupsWidth
    viewportHeight := m.ctx.Content[1] - 6

    m.listGroups.SetSize(
      listGroupsWidth,
      m.ctx.Content[1],
    )

    viewportStyle.Width(viewportWidth)
    viewportStyle.Height(viewportHeight)
    m.viewport.Width =  viewportWidth - 4
    m.viewport.Height = viewportHeight
    // cmds = append(cmds, viewport.Sync(m.viewport))

  case []list.Item:
    m.items = msg
    m.listGroups.SetItems(m.items)
    m.ctx.Loading = false
  }

  var cmd tea.Cmd

  if m.focused == 0 {
    viewportStyle.BorderForeground(lipgloss.Color("#874BFD"))
    m.listGroups, cmd = m.listGroups.Update(msg)
    cmds = append(cmds, cmd)
  } else if m.focused == 1 {
    viewportStyle.BorderForeground(lipgloss.Color("#FFFFFF"))
    m.viewport, cmd = m.viewport.Update(msg)
    cmds = append(cmds, cmd)
  }

  return m, tea.Batch(cmds...)
}

func (m Model) View() (string) {
  var view string

  view = lipgloss.JoinHorizontal(
    lipgloss.Top,
    m.listGroups.View(),
    viewportStyle.Render(m.viewport.View()),
  )

  return view
}

func (m *Model) refresh() (tea.Cmd) {
  return func () (tea.Msg) {
    var items []list.Item

    logGroups, err :=  (*m.ctx.Cloud).ListLogGroups(true)
    if err != nil {
      fmt.Printf("%s", err) // TODO: Implement error message
      return items
    }
    for _, logGroup := range logGroups {
      items = append(items, logGroup)
    }

    return items
  }
}

func (m *Model) renderViewport(logStream *models.LogStream) (string) {
  var vp string = ""

  for _, event := range logStream.LogEvents {
    vp = fmt.Sprintf(
      "%s%s\n",
      vp,
      event.Message,
    )
  }

  return vp
}

