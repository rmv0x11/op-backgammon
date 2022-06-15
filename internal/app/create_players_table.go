package app

import "context"

func (i *Implementation) CreatePlayersTable(ctx context.Context) error {
	return i.db.CreatePlayersTable(ctx)
}
