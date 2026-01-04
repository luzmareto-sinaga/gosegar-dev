package seeds

import (
	"log"
	"user-service/internal/core/domain/model"

	"gorm.io/gorm"
)

// Membuat roles super admin dan customer
func SeedRole(db *gorm.DB) {
	roles := []model.Role{
		{
			Name: "Super Admin",
		},
		{
			Name: "Customer",
		},
	}

	// Create otomatis jika tidak ada value role
	for _, role := range roles {
		if err := db.FirstOrCreate(&role, model.Role{Name: role.Name}).Error; err != nil {
			log.Fatal("%s: %v", err.Error(), err)
		} else {
			log.Printf("Role %s created", role.Name)
		}
	}
}
