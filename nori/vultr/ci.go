package vultr

import (
	"errors"

	"github.com/mrusme/planor/nori/models"
)

func (cloud *Vultr) ListPipelines() ([]models.Pipeline, error) {
	return nil, errors.New("Unsupported")
}
