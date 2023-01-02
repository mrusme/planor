package heroku

import (
	"context"
	// "errors"

	"github.com/mrusme/planor/nori/models"

	herokugo "github.com/heroku/heroku-go/v5"
)

func (cloud *Heroku) ListApps() ([]models.App, error) {
	input := herokugo.ListRange{
		Field:      "id",
		Max:        100,
		Descending: false,
	}

	ret, err := cloud.heroku.AppList(context.Background(), &input)
	if err != nil {
		return nil, err
	}

	var apps []models.App
	for _, app := range ret {
		newApp := models.App{
			ID:   app.ID,
			Name: app.Name,
		}

		apps = append(apps, newApp)
	}

	return apps, nil
}
