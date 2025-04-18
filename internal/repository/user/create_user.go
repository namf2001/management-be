package user

import "context"

// CreateUser creates a new user in the database.
func (i impl) CreateUser(ctx context.Context, username, email, password string) (int, error) {
	newUser := i.entClient.User.Create().SetUsername(username).SetEmail(email).SetPasswordHash(password)
	createdUser, err := newUser.Save(ctx)
	if err != nil {
		return 0, err
	}

	return int(createdUser.ID), nil
}
