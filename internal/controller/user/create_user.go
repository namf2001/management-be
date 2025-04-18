package user

import "context"

func (i impl) CreateUser(ctx context.Context, username, email, password string) (int, error) {
	userId, err := i.repo.User().CreateUser(ctx, username, email, password)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
