package repo

import (
	"context"
	"fmt"
	"wfm_api/internal/model"
)

func (r * Repo) GetUserByAccount(ctx context.Context, account string) (model.User, error) {
	return model.User{}, fmt.Errorf("Not implemented")
}

func (r *  Repo) CreateUser(ctx context.Context, account, password, role string ) (model.User, error) {
	return model.User{}, fmt.Errorf("Not implemented")
}

func (r * Repo) DeleteUser(ctx context.Context, user model.User) error {
	return fmt.Errorf("Not implemented")
}

func (r * Repo) UpdateUser(ctx context.Context, user * model.User) (*model.User, error) {
	return user, fmt.Errorf("Not implemented")
}
