package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

// type of /login HTTP request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// type of /register HTTP request
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// given a file (named fileName) and a text, inserts the text in the file with end of line
func add(fileName, text string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	if _, err := file.WriteString(text + "\n"); err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func performRegister(name, email, username, password string) (bool, error) {
	fileUsernames, err := os.OpenFile("data/usernames.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false, err
	}
	defer fileUsernames.Close()

	scanner := bufio.NewScanner(fileUsernames)
	for scanner.Scan() { // cheks that the username is not taken
		if scanner.Text() == username {
			return false, errors.New("username is already being used")
		}
	}

	//add all the information
	err = add("data/usernames.txt", username)
	if err != nil {
		return false, err
	}
	err = add("data/names.txt", name)
	if err != nil {
		return false, err
	}
	err = add("data/passwords.txt", password)
	if err != nil {
		return false, err
	}
	err = add("data/emails.txt", email)
	if err != nil {
		return false, err
	}

	return true, nil
}

func performLogin(username, password string) (bool, error) {
	fileUsernames, err := os.OpenFile("data/usernames.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false, err
	}
	defer fileUsernames.Close()

	filePasswords, err := os.OpenFile("data/passwords.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false, err
	}
	defer filePasswords.Close()

	// cheks that the username exists and that its password is correct
	scannerUsernames := bufio.NewScanner(fileUsernames)
	scannerPasswords := bufio.NewScanner(filePasswords)

	for scannerUsernames.Scan() {
		password := scannerPasswords.Text()

		if scannerUsernames.Text() == username {
			return password == password, nil
		}
	}

	return false, errors.New("user not registered")
}

func login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	succ, err := performLogin(req.Form["username"], req.Form["password"])
	value, _ := json.Marshal(map[string]interface{}{
		"success": succ,
		"error":   err.Error(),
	})

	w.Header().Set("content-type", "application/json")
	w.Write(value)
}

func register(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	succ, err := performLogin(req.Form["name"], req.Form["email"], req.Form["username"], req.Form["password"])
	value, _ := json.Marshal(map[string]interface{}{
		"success": succ,
		"error":   err.Error(),
	})

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
