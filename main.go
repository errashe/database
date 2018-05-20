package main

import (
	. "fmt"
	"math/rand"
	"time"

	. "./er"
)

var t time.Time

func main() {
	rand.Seed(time.Now().Unix())
	// s := Servers{}

	// t = time.Now()
	// for i := 0; i < 90000; i++ {
	// 	n := rand.Intn(9)
	// 	s = append(s, Server{
	// 		i,
	// 		Sprintf("%[1]d%[1]d.%[1]d%[1]d.%[1]d%[1]d.%[1]d%[1]d", n),
	// 		rand.Intn(2),
	// 		time.Now().Add(time.Duration(rand.Intn(500)-250) * time.Minute),
	// 	})
	// }
	// Println("FILLING: ", time.Since(t))

	// t = time.Now()
	// s.SortByPriority()
	// Println("SORTING: ", time.Since(t))

	// func() {
	// 	t := time.Now()
	// 	defer func() { Println("PRINTING: ", time.Since(t)) }()
	// 	Println(s)
	// }()

	// t = time.Now()
	// s.Save("data.txt")
	// Println("SAVING: ", time.Since(t), len(s))

	t = time.Now()
	var q Servers
	err := q.Load("data.txt")
	if err != nil {
		Println(err)
	}
	Println("LOADING: ", time.Since(t), len(q))

	// var q Servers
	// q.Load("data.er")
	// Println(q)
	// Println(time.Since(t))

	// t = time.Now()
	// q.Save("data.er")
	// Println(time.Since(t))
}
