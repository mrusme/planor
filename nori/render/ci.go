package render

import (
  "errors"

  "github.com/mrusme/planor/nori/models"
)

func (cloud *Render) ListPipelines() ([]models.Pipeline, error) {
  return nil, errors.New("Unsupported")
}
