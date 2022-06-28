package util

import (
	"github.com/rmv0x11/op-backgammon/internal/model"
	"log"
	"strconv"
	"strings"
)

func PlayersIntoIDs(p []*model.Player) string {
	var res string

	for _, v := range p {
		res = res + "," + strconv.FormatInt(v.ID, 10)
	}
	return res
}

func RoundsIntoIDs(r []*model.Round) string {
	var res string

	for _, v := range r {
		res = res + "," + strconv.FormatInt(v.ID, 10)
	}
	return res
}

func IDsIntoRounds(s string) []*model.Round {
	rounds := make([]*model.Round, 0)

	ssplit := strings.Split(s, ",")

	for _, v := range ssplit {
		r := new(model.Round)

		vparsed, err := strconv.ParseInt(v, 64, 10)
		if err != nil {
			log.Fatalln("unable parse int, err: ", err)
		}

		r.ID = vparsed

		rounds = append(rounds, r)
	}
	return rounds
}
