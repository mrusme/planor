package logs

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mrusme/planor/nori/models"
	"github.com/mrusme/planor/ui/uictx"
)

var (
	listStyle = lipgloss.NewStyle().
			Margin(0, 0, 0, 0).
			Padding(1, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

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
	Refresh     key.Binding
	Select      key.Binding
	GoBack      key.Binding
	SwitchFocus key.Binding
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
	GoBack: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "go back"),
	),
	SwitchFocus: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch focus"),
	),
}

type Model struct {
	keymap     KeyMap
	listGroups list.Model
	items      []list.Item
	viewport   viewport.Model
	ctx        *uictx.Ctx

	focused    int
	focusables [2]tea.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func NewModel(ctx *uictx.Ctx) Model {
	m := Model{
		keymap:  DefaultKeyMap,
		focused: 0,
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
			m.listGroups.Title = "Groups"
			cmds = append(cmds, m.refresh())

		case key.Matches(msg, m.keymap.SwitchFocus):
			m.focused++
			if m.focused >= len(m.focusables) {
				m.focused = 0
			}
			// return m, nil

		case key.Matches(msg, m.keymap.GoBack):
			m.listGroups.SetItems(m.items)
			return m, nil

		case key.Matches(msg, m.keymap.Select):
			group, ok := m.listGroups.SelectedItem().(models.LogGroup)
			if ok {
				m.listGroups.Title = "Streams"
				(*m.ctx.Cloud).UpdateLogStreams(&group, false)
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
				(*m.ctx.Cloud).UpdateLogEvents(&stream)
				m.viewport.SetContent(m.renderViewport(&stream))
				return m, nil
			}
		}

	case tea.WindowSizeMsg:
		listWidth := int(math.Floor(float64(m.ctx.Content[0]) / 4.0))
		listHeight := m.ctx.Content[1] - 1
		viewportWidth := m.ctx.Content[0] - listWidth - 4
		viewportHeight := m.ctx.Content[1] - 1

		listStyle.Width(listWidth)
		listStyle.Height(listHeight)
		m.listGroups.SetSize(
			listWidth-2,
			listHeight-2,
		)

		viewportStyle.Width(viewportWidth)
		viewportStyle.Height(viewportHeight)
		m.viewport = viewport.New(viewportWidth-4, viewportHeight-4)
		m.viewport.Width = viewportWidth - 4
		m.viewport.Height = viewportHeight - 4
		// cmds = append(cmds, viewport.Sync(m.viewport))

	case []list.Item:
		m.items = msg
		m.listGroups.SetItems(m.items)
		m.ctx.Loading = false
	}

	var cmd tea.Cmd

	if m.focused == 0 {
		listStyle.BorderForeground(lipgloss.Color("#FFFFFF"))
		viewportStyle.BorderForeground(lipgloss.Color("#874BFD"))
		m.listGroups, cmd = m.listGroups.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.focused == 1 {
		listStyle.BorderForeground(lipgloss.Color("#874BFD"))
		viewportStyle.BorderForeground(lipgloss.Color("#FFFFFF"))
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var view string

	view = lipgloss.JoinHorizontal(
		lipgloss.Top,
		listStyle.Render(m.listGroups.View()),
		viewportStyle.Render(m.viewport.View()),
	)

	return view
}

func (m *Model) refresh() tea.Cmd {
	return func() tea.Msg {
		var items []list.Item

		logGroups, err := (*m.ctx.Cloud).ListLogGroups(false, false)
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

func (m *Model) renderViewport(logStream *models.LogStream) string {
	var vp string = ""

	for _, event := range logStream.LogEvents {
		cleaned := strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) && r != '\\' {
				return r
			}
			return -1
		}, event.Message)

		vp = fmt.Sprintf(
			"%s%s\n",
			vp,
			cleaned,
		)
	}

	return vp
}
