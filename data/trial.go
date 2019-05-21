package data

import (
	"time"
)

type Timestamp time.Time

type Trial struct {
  ID: string
  ExerciseID: string
  SectionID: string
  Date: Timestamp
}
