package main

import (
    "fmt"
    "time"
    "sort"
    "bufio"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

)

type clock struct {
    hour int
    min int
}

type lesson struct {
    subject         string
    start_time      clock
    end_time        clock
    feedback        [11]int
}

func min_lesson (l1, l2 lesson) lesson {
    if leq(l1.start_time, l2.start_time) {
        return l1
    }
    return l2
}

func NewLesson (subject string, start_time, end_time clock) lesson {
	var _empty [11]int
    l := lesson {subject, start_time, end_time, _empty}
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
    name string
    teacher string
    schedule [7]lesson
}

func NewSubject (name, teacher string, schedule [7] lesson) subject {
    s := subject {name, teacher, schedule}
    return s
}

type user struct {
    username string
    email string
    name string
    subjects [] string
    feedback [] int
}

func NewUser (username, name, email string, subjects []string) user {
	var _empty []int
    u := user {username, name, email, subjects, _empty}
    return u
}


var map_subjects map[string]subject
var map_users map[string]user


func performset_feedback(u user, pointsstring string, subject_name string) {
	points, _ := strconv.Atoi(pointsstring)


	today := int(time.Now().Weekday()) - 1
	for i, x := range u.subjects {
	    if x == subject_name {
	        u.feedback[i] = points
	        for j, y := range map_subjects {
	            if y.name == subject_name {
	            	temp := map_subjects[j]
	            	temp.schedule[today].feedback[points]++
	                map_subjects[j] = temp
	            }
	        }
	    }
	}
}

func getInfoUser (u string) (map[string] string, bool) {
	_user, err := map_users[u]

	var m map[string]string

	m["username"] = _user.username
	m["name"] = _user.name
	m["email"] = _user.email
	
	for i, v := range _user.subjects {
		m["subject" + string(i)] = v
	}

	for i, v := range _user.feedback {
		m["feedback" + string(i)] = string(v)
	}

	return m, err
}


func performSchedule(u user) []lesson {
    var ans []lesson
    today := int(time.Now().Weekday()) - 1
    for _, name := range u.subjects {
        subject := map_subjects[name]
        x := subject.schedule[today]
        if x.start_time.hour != -1 {
            ans = append(ans, x)
        }
    }
    sort.Slice(ans, func(i, j int) bool {
        x := ans[i].start_time
        y := ans[j].start_time
        return less(x, y)
        })
    return ans
}


func set_feedback(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	
	performset_feedback(map_users[req.Form["username"][0]], req.Form["points"][0], req.Form["subject_name"][0])

	var m map[string]string

	m["result"] = "true"
	value, _ := json.Marshal(m)

	w.Header().Set("content-type", "application/json")
	w.Write(value)
}

func schedule(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	result := performSchedule(map_users[req.Form["username"][0]])

	var m map[string]lesson
	for i, v := range result {
		m[string(i)] = v
	}

	value, _ := json.Marshal(m)

	w.Header().Set("content-type", "application/json")
	w.Write(value)
}


func infouser(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	result, _ := getInfoUser(req.Form["username"][0])

	value, _ := json.Marshal(result)

	w.Header().Set("content-type", "application/json")
	w.Write(value)
}


// SetupHandlers initiates the servers HTTP endpoints
func SetupHandlers() {
	http.HandleFunc("/set_feedback", set_feedback)
	http.HandleFunc("/schedule", schedule)
	http.HandleFunc("/infouser", infouser)
	http.ListenAndServe(":8080", nil)
}


func init_everything () error {
	map_subjects = make (map[string]subject)
	map_users = make (map[string]user)

	filesUsernames, err := os.OpenFile("data/usernames.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesUsernames.Close()

	filesNames, err := os.OpenFile("data/names.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesNames.Close()

	filesEmails, err := os.OpenFile("data/emails.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesEmails.Close()

	ScannerUsernames := bufio.NewScanner(filesUsernames)
	ScannerNames := bufio.NewScanner(filesNames)
	ScannerEmails := bufio.NewScanner(filesEmails)

	for ScannerUsernames.Scan() { // cheks that the username is not taken
		ScannerNames.Scan()
		ScannerEmails.Scan()

		_username := ScannerUsernames.Text()
		_name := ScannerNames.Text()
		_email := ScannerEmails.Text()
		

		var _empty []string
		map_users[_username] = NewUser(_username, _name, _email, _empty)
	}

	
	filesAssignments, err := os.OpenFile("data/Assignments.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesAssignments.Close()

	ScannerAssignments := bufio.NewScanner(filesAssignments)
	
	for ScannerAssignments.Scan() { // cheks that the username is not taken
		_username := ScannerAssignments.Text()
		ScannerAssignments.Scan()
		_assignment := ScannerAssignments.Text()

		temp_user := map_users[_username] //EOOO
		temp_user.subjects = append(temp_user.subjects, _assignment)
		map_users[_username] = temp_user
	}

	filesSubject, err := os.OpenFile("data/subject.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesSubject.Close()

	ScannerSubject := bufio.NewScanner(filesSubject)
	
	for ScannerSubject.Scan() { // cheks that the username is not taken
		_assignment := ScannerSubject.Text()
		ScannerSubject.Scan()
		_teacher := ScannerSubject.Text()
		
		var _empty [7]lesson
		map_subjects[_assignment] = NewSubject(_assignment, _teacher, _empty)
	}

	
	filesSchedules, err := os.OpenFile("data/schedule.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer filesSchedules.Close()

	ScannerSchedules := bufio.NewScanner(filesSchedules)
	
	for ScannerSchedules.Scan() { // cheks that the username is not taken
		_assignment := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_day := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_starth := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_startm := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_endh := ScannerSchedules.Text()
		ScannerSchedules.Scan()
		_endm := ScannerSchedules.Text()

		temp_subject := map_subjects[_assignment]
		
		int_starth, _ := strconv.Atoi(_starth) 
		int_startm, _ := strconv.Atoi(_startm) 
		int_endh, _ := strconv.Atoi(_endh) 
		int_endm, _ := strconv.Atoi(_endm) 
		int_day, _ := strconv.Atoi(_day) 

		tmp_less := NewLesson(_assignment, NewClock(int_starth, int_startm), NewClock(int_endh, int_endm))

		temp_subject.schedule[int_day] = tmp_less
		map_subjects[_assignment] = temp_subject		
	}
	

	return errors.New("Everything okay :)")
}


func main() {
	init_everything()

	//fmt.Println(map_users["Javier"])

	SetupHandlers()
}

