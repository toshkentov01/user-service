package storage

import (
	"github.com/toshkentov01/alif-tech-task/user-service/storage/repo"
	"github.com/toshkentov01/alif-tech-task/user-service/storage/postgres"
)
// I is an interface for storage
type I interface {
	User() repo.UserRepository
}

type storage struct {
	userRepo repo.UserRepository
}

// NewStorage ...
func NewStorage() I {
	return &storage{
		userRepo: postgres.NewUserRepo(),
	}
}

func (s storage) User() repo.UserRepository {
	return s.userRepo
}