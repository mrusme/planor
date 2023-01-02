package fleek

import (
	"errors"

	"github.com/mrusme/planor/nori/models"
)

func (cloud *Fleek) ListLogGroups(updateStreams bool, updateEvents bool) ([]models.LogGroup, error) {
	return nil, errors.New("Unsupported")
}

func (cloud *Fleek) UpdateLogStreams(logGroup *models.LogGroup, updateEvents bool) error {
	return errors.New("Unsupported")
}

func (cloud *Fleek) UpdateLogEvents(logStream *models.LogStream) error {
	return errors.New("Unsupported")
}
