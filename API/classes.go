package main

import (
    "fmt"
    "time"
)

type clock struct {
    hour int
    min int
}

type lesson struct {
    subject         string
    start_time    clock
    end_time        clock
}

func min_lesson (l1, l2) lesson {
    if l1.start_time <= l2.start_time {
        return l1
    }
    return l2
}

func NewLesson (subject string, start_time, end_time clock) lesson {
    l := lesson {subject, start_time, end_time}
    return l
}

func NewClock(hour, min int) clock {
    t := clock {hour, min}
    return t
}

func less(a, b clock) bool {
    return (a.hour < b.hour) || (a.hour == b.hour && a.min < b.min)
}

func greater(a, b clock) bool {
    return less(b, a)
}

func leq(a, b clock) bool {
    return less(a, b) || a == b
}

func geq(a, b clock) bool {
    return leq(b, a)
}

type subject struct {
    // dni for student 
    name string
    teacher string
    subject_lessons [7] lesson
}

func NewSubject (name, teacher string, subject_lessons [7] lesson) subject {
    s := subject {name, teacher, subject_lessons}
    return s
}

var map_subjects map[string]subject


type user struct {
    dni int
    subjects [] string
    past_lessons [] lesson
}

func NewUser (dni int, subjects []string, past_lessons []lesson) user {
    u := user {dni, subjects, past_lessons}
    return u
}

func nextLesson(u user) lesson {
    var today = int(time.Now().Weekday()) - 1
    next_lesson := NewLesson("fail", 99, 99)

    hour, min, _ := time.Now().Clock()
    current_time := NewClock(hour, min)

    for _, temp_subject_string := range u.subject{
        temp_lesson = map_subjects[temp_subject_string].subject_lessons[day]

        if geq(temp_lesson.end_time, current_time) {
            next_lesson = min_lesson(next_lesson, temp_lesson)
        }
    }

    return next_lesson
}


// updates past lessons of the current day
func update_past_lessons(u user) user {
    var today = int(time.Now().Weekday())-1

    hour, min, _ := time.Now().Clock()
    current_time := NewClock(hour, min)

    u.past_lessons = nil
    for _, name := range u.subjects {
        subject, ok := map_subjects[name]
        if ok && subject.schedule[today] {
            if leq(subject.end_time[today], current_time) {
                l := NewLesson(subject.name, subject.start_time[today], subject.end_time[today])
                u.past_lessons = append(u.past_lessons, l)
            }
        }
    }
    return u
}

func main() {
    map_subjects = make(map[string]subject)
    var list_subjects []subject
    algebra := NewSubject("algebra", "Casanellas",
                            [7]bool{true, true, true, true, true, true, false},
                            [7]clock{{8, 0}, {8, 0}, {8, 0}, {8, 0}, {8, 0}, {11, 0}, {-1, -1}},
                            [7]clock{{9, 0}, {9, 0}, {9, 0}, {9, 0}, {9, 0}, {11, 30}, {-1, -1}})
    list_subjects = append(list_subjects, algebra)
    map_subjects["algebra"] = algebra
    calcul := NewSubject("calcul", "Noy",
                            [7]bool{true, false, true, false, true, true, false},
                            [7]clock{{10, 0}, {10, 0}, {10, 0}, {10, 0}, {10, 0}, {11, 30}, {-1, -1}},
                            [7]clock{{11, 0}, {11, 0}, {11, 0}, {11, 0}, {11, 0}, {12, 30}, {-1, -1}})
    list_subjects = append(list_subjects, calcul)
    map_subjects["calcul"] = calcul
    u1 := NewUser("Max", []string{"algebra", "calcul", "io"}, []lesson{})
    fmt.Println(u1.dni)
    fmt.Println(nextLesson(u1))
}
