package heroku

import (
  "os"

  "github.com/mrusme/planor/nori/adapter"

  herokugo "github.com/heroku/heroku-go/v5"
)

type Heroku struct {
  apiKey        string
  heroku        *herokugo.Service
}

func (cloud *Heroku) LoadProfile(profile *string) (error)  {
  cloud.apiKey = os.Getenv(*profile)

  return nil
}

func (cloud *Heroku) LoadClients() (error) {
  herokugo.DefaultTransport.BearerToken = cloud.apiKey

  cloud.heroku = herokugo.NewService(herokugo.DefaultClient)

  return nil
}

func (cloud *Heroku) GetCapabilities() ([]adapter.Capability) {
  var caps []adapter.Capability

  caps = append(caps, adapter.Capability{
    ID: "instances",
    Name: "Dynos",
  })
  caps = append(caps, adapter.Capability{
    ID: "ci",
    Name: "Builds",
  })

  return caps
}


