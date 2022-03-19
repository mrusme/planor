package amazon

import (
  "context"
  "time"

  "github.com/mrusme/planor/nori/models"

  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
  "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func (cloud *Amazon) ListLogGroups(updateEvents bool) ([]models.LogGroup, error) {
  input := cloudwatchlogs.DescribeLogGroupsInput{
    Limit: aws.Int32(50),
    NextToken: nil,
  }

  var logGroups []models.LogGroup

  for true {
    ret, err := cloud.cwl.DescribeLogGroups(context.Background(), &input)
    if err != nil {
      return nil, err
    }

    for _, logGroup := range ret.LogGroups {
      newLogGroup := models.LogGroup{
        ID: *logGroup.Arn,
        Name: *logGroup.LogGroupName,
        SizeBytes: *logGroup.StoredBytes,
      }

      err := cloud.UpdateLogStreams(&newLogGroup, updateEvents)
      if err != nil {
        return nil, err
      }

      logGroups = append(logGroups, newLogGroup)
    }

    input.NextToken = ret.NextToken
    if input.NextToken == nil {
      break
    }
  }


  return logGroups, nil
}

func (cloud *Amazon) UpdateLogStreams(logGroup *models.LogGroup, updateEvents bool) (error) {
  input := cloudwatchlogs.DescribeLogStreamsInput{
    LogGroupName: aws.String(logGroup.Name),
    Limit: aws.Int32(2),
    NextToken: nil,
    Descending: aws.Bool(true),
    // OrderBy: https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs@v1.15.1/types#OrderBy
    OrderBy: types.OrderByLastEventTime,
  }

  for true {
    ret, err := cloud.cwl.DescribeLogStreams(context.Background(), &input)
    if err != nil {
      return err
    }

    for _, logStream := range ret.LogStreams {
      newLogStream := models.LogStream{
        ID: *logStream.Arn,
        Name: *logStream.LogStreamName,
        GroupName: logGroup.Name,
        CreatedAt: time.UnixMilli(*logStream.CreationTime),
        FirstEventAt: time.UnixMilli(*logStream.FirstEventTimestamp),
        LastEventAt: time.UnixMilli(*logStream.LastEventTimestamp),
        LastIngestedAt: time.UnixMilli(*logStream.LastIngestionTime),
      }

      if updateEvents == true {
        err := cloud.UpdateLogEvents(&newLogStream)
        if err != nil {
          return err
        }
      }

      logGroup.Streams = append(logGroup.Streams, newLogStream)
    }

    // input.NextToken = ret.NextToken
    // if input.NextToken == nil {
    //   break
    // }
    break
  }


  return nil
}

func (cloud *Amazon) UpdateLogEvents(logStream *models.LogStream) (error) {
  input := cloudwatchlogs.GetLogEventsInput{
    LogGroupName: aws.String(logStream.GroupName),
    LogStreamName: aws.String(logStream.Name),
    Limit: aws.Int32(100),
    NextToken: nil,
    StartFromHead: aws.Bool(false),
  }

  for true {
    ret, err := cloud.cwl.GetLogEvents(context.Background(), &input)
    if err != nil {
      return err
    }

    for _, logEvent := range ret.Events {
      newLogEvent := models.LogEvent{
        Message: *logEvent.Message,
        Timestamp: time.UnixMilli(*logEvent.Timestamp),
        IngestedAt: time.UnixMilli(*logEvent.IngestionTime),
      }

      logStream.LogEvents = append(logStream.LogEvents, newLogEvent)
    }

    // input.NextToken = ret.NextForwardToken
    // if input.NextToken == nil {
    //   break
    // }
    break
  }


  return nil
}
