package uictx

import (
  "github.com/mrusme/planor/nori"
)

type Ctx struct {
  Screen          [2]int
  Content         [2]int
  Cloud           *nori.Nor
}

func New(cloud *nori.Nor) (Ctx) {
  return Ctx{
    Screen: [2]int{0, 0},
    Content: [2]int{0, 0},
    Cloud: cloud,
  }
}

