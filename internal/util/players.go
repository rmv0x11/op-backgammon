package util

import (
	"github.com/rmv0x11/op-backgammon/internal/model"
	"strconv"
)

func PlayersIntoIDs(p []*model.Player) string {
	var res string

	for _, v := range p {
		res = res + "," + strconv.FormatInt(v.ID, 10)
	}
	return res
}
