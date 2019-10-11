package lesson

/*import (
  "fmt"
)*/

type lesson struct {
  subject     string
  teacher     string
  start_hour  int
  start_min   int
  end_hour    int
  end_min     int
}

func NewLesson (subject, teacher string, start_hour, start_min, end_hour, end_min int) lesson {
  l := lesson {subject, teacher, start_hour, start_min, end_hour, end_min}
  return l
}
