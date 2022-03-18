package amazon

import (
  "context"

  "github.com/mrusme/planor/nori/models"

  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/codepipeline"
)

func (cloud *Amazon) ListPipelines() ([]models.Pipeline, error) {
  input := codepipeline.ListPipelinesInput{
    MaxResults: aws.Int32(1000),
    NextToken: nil,
  }

  var pipelines []models.Pipeline

  for true {
    ret, err := cloud.cpc.ListPipelines(context.Background(), &input)
    if err != nil {
      return nil, err
    }

    for _, pipeline := range ret.Pipelines {
      newPipeline := models.Pipeline{
        ID: *pipeline.Name,
        Name: *pipeline.Name,
        Version: string(*pipeline.Version),
        CreatedAt: *pipeline.Created,
        UpdatedAt: *pipeline.Updated,
      }

      cloud.UpdatePipelineStatus(&newPipeline)

      pipelines = append(pipelines, newPipeline)
    }

    input.NextToken = ret.NextToken
    if input.NextToken == nil {
      break
    }
  }


  return pipelines, nil
}

func (cloud *Amazon) UpdatePipelineStatus(pipeline *models.Pipeline) (error) {
  stateRet, err := cloud.cpc.GetPipelineState(
    context.Background(),
    &codepipeline.GetPipelineStateInput{
      Name: &pipeline.Name,
    },
  )

  var state = stateRet
  if err != nil {
    pipeline.Error = err
    return err
  }

  var stages []models.PipelineStage
  for _, stage := range state.StageStates {
    var actions []models.PipelineStageAction
    for _, action := range stage.ActionStates {
      newAction := models.PipelineStageAction {
        Name: *action.ActionName,
        Status: string((*action.LatestExecution).Status),
        UpdatedAt: *action.LatestExecution.LastStatusChange,
      }

      if action.LatestExecution.Summary != nil {
        newAction.Summary = *action.LatestExecution.Summary
      }

      if action.LatestExecution.PercentComplete != nil {
        newAction.PercentComplete = *action.LatestExecution.PercentComplete
      }

      actions = append(actions, newAction)
    }

    newStage := models.PipelineStage{
      ID: *stage.StageName,
      Name: *stage.StageName,
      Status: string((*stage.LatestExecution).Status),
      Actions: actions,
    }

    stages = append(stages, newStage)
  }

  pipeline.Version = string(*state.PipelineVersion)
  pipeline.Stages = stages

  return nil
}
