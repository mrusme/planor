package ci

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mrusme/planor/ui/uictx"
)

var (
  appStyle = lipgloss.NewStyle().Padding(1, 2)

  titleStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#FFFDF5")).
    Background(lipgloss.Color("#25A065")).
    Padding(0, 1)

  statusMessageStyle = lipgloss.NewStyle().
    Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
    Render
)

type Model struct {
  list            list.Model
  items           []list.Item
  ctx             *uictx.Ctx
}

func (m Model) Init() tea.Cmd {
  return nil
}

func NewModel(ctx *uictx.Ctx) (Model) {
  m := Model{}

  pipelines, err :=  (*ctx.Cloud).ListPipelines()
  if err == nil {
    for _, pipeline := range pipelines {
      m.items = append(m.items, pipeline)
    }
  }

  m.list = list.New(m.items, list.NewDefaultDelegate(), 0, 0)
  m.ctx = ctx

  return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.WindowSizeMsg:
    topGap, rightGap, bottomGap, leftGap := appStyle.GetPadding()
    m.list.SetSize(msg.Width-leftGap-rightGap, msg.Height-topGap-bottomGap)

  }


  var cmd tea.Cmd
  m.list, cmd = m.list.Update(msg)

  return m, cmd
}

func (m Model) View() (string) {
  return appStyle.Render(m.list.View())
}

