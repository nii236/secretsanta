package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Santas struct {
	santas []*Santa
}

type Santa struct {
	Person    string
	HasTarget bool
	Target    string
}

func main() {
	people := []string{"Guy", "Brush", "Threep", "Wood"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := r.Perm(len(people))
	shuffledPeople := []string{}
	santas := []*Santa{}
	for _, person := range people {
		santas = append(santas, &Santa{
			Person:    person,
			HasTarget: false,
			Target:    "",
		})
	}
	for _, i := range list {
		shuffledPeople = append(shuffledPeople, people[i])
	}

	for _, santa := range santas {
		for i, person := range shuffledPeople {
			if person == "" {
				continue
			}
			if person == santa.Person {
				continue
			}
			santa.Target = person
			santa.HasTarget = true
			shuffledPeople[i] = ""
			break
		}
	}

	for _, santa := range santas {
		fmt.Printf("%s -> %s\n", santa.Person, santa.Target)

	}

}
