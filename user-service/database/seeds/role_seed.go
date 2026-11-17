package seeds

import (
	"log"

	"github.com/terenjit/ecommerce-veggie/user-service/internal/core/domain/model"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []model.Role{
		{
			Name: "Super Admin",
		},
		{
			Name: "Customer",
		},
	}

	for _, role := range roles {
		if err := db.Where(model.Role{Name: role.Name}).FirstOrCreate(&role).Error; err != nil {
			log.Fatalf("Failed to seed role %s: %v", role.Name, err)
		}
	}
}
