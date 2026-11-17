package seeds

import (
	"log"

	"github.com/terenjit/ecommerce-veggie/user-service/internal/core/domain/model"
	"github.com/terenjit/ecommerce-veggie/user-service/utils/conv"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatalf("error while seeding user : %v", err)
	}

	role := model.Role{Name: "Super Admin"}
	if err := db.Where(model.Role{Name: role.Name}).FirstOrCreate(&role).Error; err != nil {
		log.Fatalf("error while seeding role : %v", err)
	}

	admin := model.User{
		Name:       "super_admin",
		Email:      "superadmin@mail.com",
		Password:   string(bytes),
		RoleID:     role.ID,
		IsVerified: true,
	}

	if err := db.Where(model.User{Email: admin.Email}).FirstOrCreate(&admin).Error; err != nil {
		log.Fatalf("error while seeding user : %v", err)
	} else {
		log.Printf("Admin user %s already seeded", admin.Name)
	}

	userRole := model.UserRole{
		UserID: admin.ID,
		RoleID: role.ID,
	}

	if err := db.Where(model.UserRole{UserID: admin.ID, RoleID: role.ID}).FirstOrCreate(&userRole).Error; err != nil {
		log.Fatalf("error while seeding user role : %v", err)
	}
}
