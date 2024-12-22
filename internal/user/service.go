package user

import (
	"context"
)

type Repository interface {
	GetMany(ctx context.Context, filter Filter) ([]User, error)
	WriteMany(ctx context.Context, users []User) error
}

type Service struct {
	Repo Repository
}

func (s *Service) GetUsers(ctx context.Context, filter Filter) ([]User, error) {
	return s.Repo.GetMany(ctx, filter)
}

func (s *Service) WriteUsers(ctx context.Context, users []User) error {
	return s.Repo.WriteMany(ctx, users)
}
