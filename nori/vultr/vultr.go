package vultr

import (
  "context"
  "os"

  "github.com/mrusme/planor/nori/adapter"

  "github.com/vultr/govultr/v2"
  "golang.org/x/oauth2"
)

type Vultr struct {
  cfg           oauth2.Config
  apiKey        string
  vultr         *govultr.Client
}

func (cloud *Vultr) LoadProfile(profile *string) (error)  {
  cloud.apiKey = os.Getenv(*profile)
  cloud.cfg = oauth2.Config{}

  return nil
}

func (cloud *Vultr) LoadClients() (error) {
  ctx := context.Background()
  ts := cloud.cfg.TokenSource(ctx, &oauth2.Token{AccessToken: cloud.apiKey})
  cloud.vultr = govultr.NewClient(oauth2.NewClient(ctx, ts))
  cloud.vultr.SetUserAgent("github.com/mrusme/planor")

  return nil
}

func (cloud *Vultr) GetCapabilities() ([]adapter.Capability) {
  var caps []adapter.Capability

  caps = append(caps, adapter.Capability{
    ID: "instances",
    Name: "Cloud Instances",
  })

  return caps
}


