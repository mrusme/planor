package render

import (
  "errors"

  "github.com/mrusme/planor/nori/models"
)


func (cloud *Render) ListLogGroups(updateStreams bool, updateEvents bool) ([]models.LogGroup, error) {
  return nil, errors.New("Unsupported")
}

func (cloud *Render) UpdateLogStreams(logGroup *models.LogGroup, updateEvents bool) (error) {
  return errors.New("Unsupported")
}

func (cloud *Render) UpdateLogEvents(logStream *models.LogStream) (error) {
  return errors.New("Unsupported")
}

