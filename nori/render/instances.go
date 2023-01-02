package render

import (
	"github.com/mrusme/planor/nori/models"
	// gorender "github.com/mrusme/go-render"
)

func (cloud *Render) ListInstances() ([]models.Instance, error) {
	var instances []models.Instance
	ret, err := cloud.render.ListServices()
	if err != nil {
		return nil, err
	}

	for _, instance := range ret {
		newInstance := models.Instance{
			ID:   instance.ID,
			Name: instance.Name,

			Type: instance.Type,

			Status: instance.Deploys[0].Status,
		}

		instances = append(instances, newInstance)
	}

	return instances, nil
}
