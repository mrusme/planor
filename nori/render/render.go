package render

import (
	// "context"
	"os"

	"github.com/mrusme/planor/nori/adapter"

	gorender "github.com/mrusme/go-render"
)

type Render struct {
	apiKey string
	render *gorender.RenderClient
}

func (cloud *Render) LoadProfile(profile *string) error {
	cloud.apiKey = os.Getenv(*profile)

	return nil
}

func (cloud *Render) LoadClients() error {
	var err error
	cloud.render, err = gorender.New(cloud.apiKey)

	return err
}

func (cloud *Render) GetCapabilities() []adapter.Capability {
	var caps []adapter.Capability

	caps = append(caps, adapter.Capability{
		ID:   "instances",
		Name: "Services",
	})

	return caps
}
