package subject

import (
  //"fmt"
)

type subject struct {
  name string
  schedule [7] bool
  start_hour [7] int
  start_min [7] int
  end_hour [7] int
  end_min [7] int
}

func NewSubject (name string, schedule [7] bool, start_hour, start_min, end_hour, end_min [7] int) subject {
  s := subject {name, schedule, start_hour, start_min, end_hour, end_min}
  return s
}
