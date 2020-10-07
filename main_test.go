package main

import (
	"database/sql"
	"github.com/lib/pq/hstore"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type TestModel struct {
	gorm.Model
	SomeMap hstore.Hstore
}

func TestGORM(t *testing.T) {
	if os.Getenv("GORM_DIALECT") != "postgres" {
		return
	}

	DB.AutoMigrate(new(TestModel))

	test := TestModel{
		SomeMap: hstore.Hstore{
			Map: map[string]sql.NullString{
				"a": sql.NullString{
					String: "",
					Valid:  false,
				},
			},
		},
	}

	DB.Create(&test)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
