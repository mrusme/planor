package navigation

import (
  "strings"

  "github.com/mrusme/planor/ui/uictx"

  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

var (
  highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

  tabsBorderHeight  = 1
  tabsContentHeight = 2
  TabsHeight        = tabsBorderHeight + tabsContentHeight

  activeTabBorder = lipgloss.Border{
    Top:         "─",
    Bottom:      " ",
    Left:        "│",
    Right:       "│",
    TopLeft:     "╭",
    TopRight:    "╮",
    BottomLeft:  "┘",
    BottomRight: "└",
  }

  tabBorder = lipgloss.Border{
    Top:         "─",
    Bottom:      "─",
    Left:        "│",
    Right:       "│",
    TopLeft:     "╭",
    TopRight:    "╮",
    BottomLeft:  "┴",
    BottomRight: "┴",
  }

  tab = lipgloss.NewStyle().
    Border(tabBorder, true).
    BorderForeground(highlight).
    Padding(0, 1)

  activeTab = tab.Copy().Border(activeTabBorder, true)

  tabGap = tab.Copy().
    BorderTop(false).
    BorderLeft(false).
    BorderRight(false)
)

var Navigation = []string{
  "CI",
}

type Model struct {
  CurrentId             int
}

func NewModel() Model {
  return Model{
    CurrentId: 0,
  }
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
  return m, nil
}

func (m Model) View(ctx uictx.Ctx) string {
  var items []string

  for i, nav := range Navigation {
    if m.CurrentId == i {
      items = append(items, activeTab.Render(nav))
    } else {
      items = append(items, tab.Render(nav))
    }
  }

  row := lipgloss.JoinHorizontal(lipgloss.Top, items...)
  gap := tabGap.Render(strings.Repeat(" ", max(0, ctx.Screen[0] - lipgloss.Width(row) - 2)))

  return lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}
