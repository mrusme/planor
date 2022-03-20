package amazon

import (
  "context"

  "github.com/mrusme/planor/nori/models"

  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/ec2"
  // "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func (cloud *Amazon) ListInstances() ([]models.Instance, error) {
  input := ec2.DescribeInstancesInput{
    MaxResults: aws.Int32(1000),
    NextToken: nil,
  }

  var instances []models.Instance

  for true {
    ret, err := cloud.ec2.DescribeInstances(context.Background(), &input)
    if err != nil {
      return nil, err
    }

    for _, reservation := range ret.Reservations {
      for _, instance := range reservation.Instances {
        newInstance := models.Instance{
          ID: *instance.InstanceId,
          Name: *instance.InstanceId,

          Type: string(instance.InstanceType),
          Architecture: string(instance.Architecture),
          CPUCores: int(*instance.CpuOptions.CoreCount),
          CPUThreads: int(*instance.CpuOptions.ThreadsPerCore),

          Image: *instance.ImageId,
          Status: string(instance.State.Name),
        }

        if instance.PublicIpAddress != nil {
          newInstance.IPv4 = *instance.PublicIpAddress
        }

        if instance.Ipv6Address != nil {
          newInstance.IPv6 = *instance.Ipv6Address
        }

        instances = append(instances, newInstance)
      }
    }

    input.NextToken = ret.NextToken
    if input.NextToken == nil {
      break
    }
  }


  return instances, nil
}

