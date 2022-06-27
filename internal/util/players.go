package util

import (
	"github.com/rmv0x11/op-backgammon/internal/storage"
	"strconv"
)

func PlayersIntoIDs(p []*storage.Player) string {
	var res string

	for _, v := range p {
		res = res + "," + strconv.FormatInt(v.ID.Int64, 10)
	}
	return res
}
