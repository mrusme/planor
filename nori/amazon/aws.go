package amazon

import (
	"context"

	"github.com/mrusme/planor/nori/adapter"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Amazon struct {
	cfg aws.Config
	ec2 *ec2.Client
	cpc *codepipeline.Client
	cwl *cloudwatchlogs.Client
}

func (cloud *Amazon) LoadProfile(profile *string) error {
	var err error

	cloud.cfg, err = config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(*profile))
	if err != nil {
		return err
	}

	return nil
}

func (cloud *Amazon) LoadClients() error {
	cloud.ec2 = ec2.NewFromConfig(cloud.cfg)
	cloud.cpc = codepipeline.NewFromConfig(cloud.cfg)
	cloud.cwl = cloudwatchlogs.NewFromConfig(cloud.cfg)
	return nil
}

func (cloud *Amazon) GetCapabilities() []adapter.Capability {
	var caps []adapter.Capability

	caps = append(caps, adapter.Capability{
		ID:   "instances",
		Name: "Elastic Cloud Compute",
	})
	caps = append(caps, adapter.Capability{
		ID:   "ci",
		Name: "CodePipeline",
	})
	caps = append(caps, adapter.Capability{
		ID:   "logs",
		Name: "CloudWatch Logs",
	})

	return caps
}
