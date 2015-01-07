package main

import (
	"fmt"
	"io"
	"time"
)

type Greet struct {
	User    string    `param:"user"`
	Message string    `param:"message"`
	Time    time.Time `param:"time"`
}

var Greets = []Greet{
	{"michelangelo", "Welcome to Gritter!", time.Now()},
	{"raphael", "Wanna know a secret?", time.Now()},
	{"botticelli", "Okay!", time.Now()},
	{"davinci", "I'm listening...", time.Now()},
}

func (g Greet) Write(w io.Writer) {
	fmt.Fprintf(w, "%s\n@%s at %s\n---\n", g.Message, g.User, g.Time.Format(time.UnixDate))
}

type User struct {
	Name, Bio string
}

var Users = map[string]User{
	"michelangelo": {"Michelangelo", "Bologna"},
	"raphael":      {"Raphael", "Naples"},
	"botticelli":   {"Botticelli", "Turin"},
	"davinci":      {"Leonardo Da Vinci", "Palermo"},
}

func (u User) Write(w io.Writer, handle string) {
	fmt.Fprintf(w, "%s (@%s)\n%s\n", u.Name, handle, u.Bio)
}
