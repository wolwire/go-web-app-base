package pkg

import (
	"github.com/flowista2/pkg/database"
	"gorm.io/gorm"
)

type Repository struct {
    db *database.Db
}

func (r *Repository) Create(value interface{}) error {
    return r.db.Create(value).Error
}

func (r *Repository) Find(id uint, value interface{}) error {
    return r.db.First(value, id).Error
}

func (r *Repository) Update(value interface{}) error {
    return r.db.Save(value).Error
}

func (r *Repository) Delete(value interface{}) error {
    return r.db.Delete(value).Error
}

func (r *Repository) FindBy(value interface{}, query interface{}, args ...interface{}) error {
    return r.db.Where(query, args...).Find(value).Error
}

func (r *Repository) Where(query interface{}, args ...interface{}) *gorm.DB {
    return r.db.Where(query, args...)
}

func (r *Repository) Preload(query string, args ...interface{}) *gorm.DB {
    return r.db.Preload(query, args...)
}
