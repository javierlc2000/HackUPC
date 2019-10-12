package main

//"log"

/*

func register (name, username, key, email string) bool {
	file_users, err := os.OpenFile("users.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer file_users.Close()

    scanner := bufio.NewScanner(file_users)
    for scanner.Scan() { // cheks that the username is not taken
        if scanner.Text() == username {
        	return false
        }
    }

    //add all the information
	if _, err = file_users.WriteString(username + "\n"); err != nil {
	    panic(err)
	}

	file_names, _ := os.OpenFile("names.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if _, err = file_names.WriteString(name + "\n"); err != nil {
	    panic(err)
	}
	defer file_names.Close()

	file_keys, _ := os.OpenFile("keys.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if _, err = file_keys.WriteString(key + "\n"); err != nil {
	    panic(err)
	}
	defer file_keys.Close()

	file_emails, _ := os.OpenFile("emails.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if _, err = file_emails.WriteString(email + "\n"); err != nil {
	    panic(err)
	}
	defer file_emails.Close()


    return true
}

func login (username, key string) bool {
	file_users, _ := os.OpenFile("users.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer file_users.Close()

    position := 0
    found := false

    // cheks that the username exists
    scanner := bufio.NewScanner(file_users)
    for scanner.Scan() {
        if scanner.Text() == username {
        	found = true
        }
        if !found {
        	position++
        }
    }
    if !found {
    	return false
    }

    check_position := 0
    file_keys, _ := os.OpenFile("keys.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

    // cheks that the username has the same key
    scanner = bufio.NewScanner(file_keys)
    for scanner.Scan() {
    	password := scanner.Text()

        if check_position == position{
        	return password == key
        }
        check_position++
    }
    defer file_keys.Close()

    return false
}

func main() {
    register("Maria", "maria", "pasnataga", "maria@gmail.com")
    fmt.Println(login("alekhine", "haskhell2"))
}*/
