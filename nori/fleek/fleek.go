package fleek

import (
  // "context"
  "os"

  "github.com/mrusme/planor/nori/adapter"

  gofleek "github.com/mrusme/go-fleek"
)

type Fleek struct {
  apiKey        string
  teamId        string
  fleek         *gofleek.FleekClient
}

func (cloud *Fleek) LoadProfile(profile *string) (error)  {
  cloud.apiKey = os.Getenv(*profile)
  cloud.teamId = os.Getenv("FLEEK_TEAM_ID")

  return nil
}

func (cloud *Fleek) LoadClients() (error) {
  var err error
  cloud.fleek, err = gofleek.New(cloud.apiKey)

  return err
}

func (cloud *Fleek) GetCapabilities() ([]adapter.Capability) {
  var caps []adapter.Capability

  caps = append(caps, adapter.Capability{
    ID: "instances",
    Name: "Sites",
  })

  return caps
}


