package migrations

import "gorm.io/gorm"

var migrations = []func(*gorm.DB) error{
	userMigrateUp,
}

func Migrate(db *gorm.DB) error {
	for _, migration := range migrations {
		err := migration(db)
		if err != nil {
			return err
		}
	}
	return nil
}