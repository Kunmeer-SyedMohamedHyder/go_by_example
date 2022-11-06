package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user
type ByName []user

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Less(i, j int) bool {
	return a[i].Last < a[j].Last
}
func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func pretty_print(users []user) {
	for ind, user := range users {
		fmt.Println("User #", ind)
		fmt.Println("\t", user.First, user.Last, user.Age)
		fmt.Println("\t", "Sorted sayings are:")
		sort.Strings(user.Sayings)
		for _, saying := range user.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	ba := ByAge(users)
	bn := ByName(users)

	sort.Sort(ba)
	fmt.Println("-----Sorted by AGE------")
	pretty_print(ba)

	sort.Sort(bn)
	fmt.Println("-----Sorted by NAME------")
	pretty_print(bn)

}
