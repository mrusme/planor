package amazon

import (
  "context"

  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/config"
  "github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

type Amazon struct {
  cfg           aws.Config
  cpc           *codepipeline.Client
}

func (cloud *Amazon) LoadProfile(profile *string) (error)  {
  var err error

  cloud.cfg, err = config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(*profile))
  if err != nil {
    return err
  }

  return nil
}

func (cloud *Amazon) LoadClients() (error) {
  cloud.cpc = codepipeline.NewFromConfig(cloud.cfg)
  return nil
}

