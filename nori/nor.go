package nori

import (
	"errors"

	"github.com/mrusme/planor/nori/adapter"
	"github.com/mrusme/planor/nori/models"

	"github.com/mrusme/planor/nori/amazon"
	"github.com/mrusme/planor/nori/fleek"
	"github.com/mrusme/planor/nori/heroku"
	"github.com/mrusme/planor/nori/render"
	"github.com/mrusme/planor/nori/vultr"
)

type Nor interface {
	GetCapabilities() []adapter.Capability

	LoadProfile(profile *string) error
	LoadClients() error

	ListInstances() ([]models.Instance, error)

	ListPipelines() ([]models.Pipeline, error)

	ListLogGroups(updateStreams bool, updateEvents bool) ([]models.LogGroup, error)
	UpdateLogStreams(logGroup *models.LogGroup, updateEvents bool) error
	UpdateLogEvents(logStream *models.LogStream) error
}

func New(cloudType *string, profile *string) (Nor, error) {
	var cloud Nor

	switch *cloudType {
	case "aws":
		cloud = new(amazon.Amazon)
	case "vultr":
		cloud = new(vultr.Vultr)
	case "heroku":
		cloud = new(heroku.Heroku)
	case "render":
		cloud = new(render.Render)
	case "fleek":
		cloud = new(fleek.Fleek)
	default:
		return nil, errors.New("No such cloud")
	}

	err := cloud.LoadProfile(profile)
	if err != nil {
		return nil, err
	}

	err = cloud.LoadClients()
	if err != nil {
		return nil, err
	}

	return cloud, nil
}
