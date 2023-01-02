package heroku

import (
	"context"
	// "errors"

	"github.com/mrusme/planor/nori/models"

	herokugo "github.com/heroku/heroku-go/v5"
)

func (cloud *Heroku) ListInstances() ([]models.Instance, error) {
	apps, err := cloud.ListApps()
	if err != nil {
		return nil, err
	}

	var dynos []models.Instance
	input := herokugo.ListRange{
		Field:      "id",
		Max:        100,
		Descending: false,
	}
	for _, app := range apps {
		ret, err := cloud.heroku.DynoList(context.Background(), app.ID, &input)
		if err != nil {
			return dynos, err
		}

		for _, dyno := range ret {
			newDyno := models.Instance{
				ID:   dyno.ID,
				Name: dyno.Name,

				Type:  dyno.Size,
				Image: dyno.Command,

				Status: dyno.State,
			}

			dynos = append(dynos, newDyno)
		}
	}

	return dynos, nil
}
