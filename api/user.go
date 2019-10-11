package user

/*import (
  "fmt"
)*/

type user struct {
  name string
  subjects []string
}

func NewUser (name string, subjects []string) user {
  u := user {name, subjects}
  return u
}
