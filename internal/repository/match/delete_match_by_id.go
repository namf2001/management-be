package match

import "context"

// DeleteMatchByID deletes a match by its ID.
func (i impl) DeleteMatchByID(ctx context.Context, id int) error {
	return i.entClient.Match.DeleteOneID(id).Exec(ctx)
}
