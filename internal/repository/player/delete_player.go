package player

import (
	"context"
)

// DeletePlayer deletes a player by ID
func (i impl) DeletePlayer(ctx context.Context, id int) error {
	return i.entClient.Player.DeleteOneID(id).Exec(ctx)
}
