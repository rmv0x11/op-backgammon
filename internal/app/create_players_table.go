package app

import "context"

func (i *Implementation) CreatePlayersTable(ctx context.Context) {
	i.db.CreatePlayersTable(ctx)
}
