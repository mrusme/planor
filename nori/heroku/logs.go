package heroku

import (
  "errors"

  "github.com/mrusme/planor/nori/models"
)


func (cloud *Heroku) ListLogGroups(updateStreams bool, updateEvents bool) ([]models.LogGroup, error) {
  return nil, errors.New("Unsupported")
}

func (cloud *Heroku) UpdateLogStreams(logGroup *models.LogGroup, updateEvents bool) (error) {
  return errors.New("Unsupported")
}

func (cloud *Heroku) UpdateLogEvents(logStream *models.LogStream) (error) {
  return errors.New("Unsupported")
}

