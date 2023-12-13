package repository

import (
	"github.com/flowista2/pkg"
)

type UserRepository struct {
	*pkg.Repository
}

func UserRep() *UserRepository {
    return &UserRepository{}
}
