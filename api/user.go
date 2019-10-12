package main

import (
  //"fmt"
  "time"
)

type lesson struct {
  subject     string
  start_time  clock
  end_time    clock
}

func NewLesson (subject string, start_time, end_time clock) lesson {
  l := lesson {subject, start_time, end_time}
  return l
}

type clock struct {
  hour int
  min int
}

func NewClock(hour, min int) clock {
  t := clock {hour, min}
  return t
}

type subject struct {
  name string
  teacher string
  schedule [7] bool
  start_time [7] clock
  end_time [7] clock
}

func NewSubject (name, teacher string, schedule [7] bool, start_time, end_time [7] clock) subject {
  s := subject {name, teacher, schedule, start_time, end_time}
  return s
}

var list_subjects map[string]subject

type user struct {
  name string
  subjects [] string
  past_lessons [] lesson
}

func NewUser (name string, subjects []string, past_lessons []lesson) user {
  u := user {name, subjects, past_lessons}
  return u
}

func nextLesson(u user) lesson {
  var today = int(time.Now().Weekday())
  hour, min, _ := time.Now().Clock()
  var next_subject string
  for _, name := range u.subjects {
    subject := list_subjects[name]
    if subject.schedule[today] {
      if (subject.end_time[today].hour > hour) || (subject.end_time[today].hour == hour && subject.end_time[today].min >= min) {
           hour = subject.end_time[today].hour
           min = subject.end_time[today].min
           next_subject = subject.name
      }
    }
  }
  subject := list_subjects[next_subject]
  var ans = NewLesson(subject.name, subject.start_time[today], subject.end_time[today])
  return ans
}

// updates past lessons of the current day
func update_past_lessons(u user) user {
  var today = int(time.Now().Weekday())
  hour, min, _ := time.Now().Clock()
  var last_time clock
  if len(u.past_lessons) == 0 {
    last_time = {0, 0}
  }
  else {
    last_lesson = u.past_lessons[len(past_lessons) - 1]
    last_time = last_lesson.start_time
  }
  for _, name := range u.subjects {
    subject = list_subjects[name]
    if subject.schedule[today] {
      if subject.start_time[today] < last_time
    }
  }
}

func main() {

}
