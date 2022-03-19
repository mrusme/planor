package vultr

import (
  "errors"

  "github.com/mrusme/planor/nori/models"
)


func (cloud *Vultr) ListLogGroups(updateStreams bool, updateEvents bool) ([]models.LogGroup, error) {
  return nil, errors.New("Unsupported")
}

func (cloud *Vultr) UpdateLogStreams(logGroup *models.LogGroup, updateEvents bool) (error) {
  return errors.New("Unsupported")
}

func (cloud *Vultr) UpdateLogEvents(logStream *models.LogStream) (error) {
  return errors.New("Unsupported")
}

