package player

import (
	"context"
)

// DeletePlayer deletes a player by ID from the database.
func (i impl) DeletePlayer(ctx context.Context, id int) error {
	return i.repo.Player().DeletePlayer(ctx, id)
}
