package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         //items on the list
	cursor   int              //which item our cursor is pointing at
	selected map[int]struct{} //which items are selected
}

func initialModel() model {
	return model{
		choices:  []string{"Buy carrots", "Buy celery", "Buy pumpkin"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
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
		case "enter", " ":
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

func (m model) View() string {
	s := "What should we buy at the market?\n\n"
	for i, choice := range m.choices {
		//Is the cursor pointing at this choice?
		cursor := " " //no cursor
		if m.cursor == i {
			cursor = ">"
		}
		checked := " " //not selected
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	s += "\nPress q to quit.\n"
	return s
}
func main() {
	p := tea.NewProgram(initialModel())
	if _,err:=p.Run();err!=nil{
		fmt.Printf("There's been an erro:%v",err)
		os.Exit(1)
	}
}
