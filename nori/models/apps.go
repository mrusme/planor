package models

// "fmt"

type App struct {
	ID   string
	Name string
}

func (app App) FilterValue() string {
	return app.Name
}

func (app App) Title() string {
	return app.Name
}

func (app App) Description() string {
	return app.ID
}
