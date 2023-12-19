package pkg

import (
	"github.com/flowista2/pkg/database"
	"gorm.io/gorm"
)

type Repository struct {
	Db *database.Db
}

func (r *Repository) Create(value interface{}) error {
	return r.Db.Create(value).Error
}

func (r *Repository) Find(id int, value interface{}) error {
	return r.Db.First(value, id).Error
}

func (r *Repository) Update(value interface{}) error {
	return r.Db.Save(value).Error
}

func (r *Repository) Delete(value interface{}) error {
	return r.Db.Delete(value).Error
}

func (r *Repository) FindBy(value interface{}, query interface{}, args ...interface{}) error {
	return r.Db.Where(query, args...).Find(value).Error
}

func (r *Repository) Where(query interface{}, args ...interface{}) *gorm.DB {
	return r.Db.Where(query, args...)
}

func (r *Repository) Preload(query string, args ...interface{}) *gorm.DB {
	return r.Db.Preload(query, args...)
}
