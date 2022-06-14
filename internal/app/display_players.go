package app

import "fmt"

func (i *Implementation) DisplayPlayers() {
	players := i.db.GetPlayers()

	for _, v := range players {
		fmt.Printf("%#v\n", v)
		//fmt.Println(v.FirstName, " ", v.LastName, "\n")

	}
}
