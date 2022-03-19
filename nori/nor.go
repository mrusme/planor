package nori

import (
  "errors"
  "github.com/mrusme/planor/nori/models"
  "github.com/mrusme/planor/nori/amazon"
)

type Nor interface {
  GetCapabilities() (map[string]string)

  LoadProfile(profile *string) (error)
  ListPipelines() ([]models.Pipeline, error)

  ListLogGroups(updateStreams bool, updateEvents bool) ([]models.LogGroup, error)
  UpdateLogStreams(logGroup *models.LogGroup, updateEvents bool) (error)
  UpdateLogEvents(logStream *models.LogStream) (error)
}

func New(cloud *string, profile *string) (Nor, error) {
  if *cloud == "aws" {
    cloud := new(amazon.Amazon)
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

  return nil, errors.New("No such cloud")
}

