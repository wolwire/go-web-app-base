package repository

import (
	"github.com/flowista2/pkg"
	"github.com/flowista2/pkg/database"
)

type UserRepository struct {
	*pkg.Repository
}

func UserRep(db *database.Db) *UserRepository {
    return &UserRepository{Repository: &pkg.Repository{Db: db}}
}
