package main

import (
    //"fmt"
    "os"
    "bufio"
    "encoding/json"
	"net/http"
	"io/ioutil"
)

type Answer struct {
    word string
    value int
}

// given a file (named file_name) and a text, inserts the text in the file with end of line
func add (file_name, text string) {
	file, _ := os.OpenFile(file_name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if _, err := file.WriteString(text + "\n"); err != nil {
	    panic(err)
	}
	defer file.Close()
}

func perform_register (name, username, password, email string) bool {
	file_usernames, _ := os.OpenFile("usernames.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer file_usernames.Close()

    scanner := bufio.NewScanner(file_usernames)
    for scanner.Scan() { // cheks that the username is not taken       
        if scanner.Text() == username {
        	return false
        }
    }

    //add all the information
    add("usernames.txt", username)
    add("names.txt", name)
    add("passwords.txt", password)
    add("emails.txt", email)

    return true
}

func perform_login (username, password string) bool {
	file_usernames, _ := os.OpenFile("usernames.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer file_usernames.Close()

	file_passwords, _ := os.OpenFile("passwords.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer file_passwords.Close()  

    // cheks that the username exists and that its password is correct      
    scanner_usernames := bufio.NewScanner(file_usernames)
    scanner_passwords := bufio.NewScanner(file_passwords)

    for scanner_usernames.Scan() {
    	password := scanner_passwords.Text()

        if scanner_usernames.Text() == username {
        	return password == password
        }
    }
    
    return false
}

type RequestData1 struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type RequestData2 struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func login(w http.ResponseWriter, req *http.Request) {
	body1, _ := ioutil.ReadAll(req.Body)
	
	var msg RequestData1
	_ = json.Unmarshal(body1, &msg)

    value, _ := json.Marshal(perform_login(msg.Username, msg.Password))

	w.Header().Set("content-type", "application/json")
	w.Write(value)
}


func register(w http.ResponseWriter, req *http.Request) {
    body1, _ := ioutil.ReadAll(req.Body)
    
    var msg RequestData2
    _ = json.Unmarshal(body1, &msg)

    value, _ := json.Marshal(msg.Name)

    w.Header().Set("content-type", "application/json")
    w.Write(value)
}

// SetupHandlers initiates the servers HTTP endpoints
func SetupHandlers() {
	http.HandleFunc("/login", login)
    http.HandleFunc("/register", register)
	http.ListenAndServe(":8080", nil)
}

func main() {
	SetupHandlers()
}
