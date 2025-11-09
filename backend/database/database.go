package database

import (
	"log"
	"time"

	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/config"
	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/entity"
	"github.com/theeraphat-mag/SUT-Club-Event-Center/backend/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.DBURL()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. ", err)
	}
	log.Println("Connected to database")

	DB = db

	// AutoMigrate entities to create/update tables in the database
	log.Println("Running database migrations...")
	db.AutoMigrate(&entity.Genders{}, &entity.Users{})
	log.Println("Database migrations completed.")

	// Data seeding
	DataSeeding(db)
}

func DataSeeding(db *gorm.DB) {
	// Seed Genders
	var genderCount int64
	db.Model(&entity.Genders{}).Count(&genderCount)
	if genderCount == 0 {
		log.Println("Seeding initial gender data...")
		genders := []entity.Genders{
			{Gender: "Male"},
			{Gender: "Female"},
			{Gender: "Other"},
		}
		for _, gender := range genders {
			db.Create(&gender)
		}
		log.Println("Gender data seeding completed.")
	} else {
		log.Println("Gender data already exists, skipping seeding.")
	}

	// Seed Users
	var userCount int64
	db.Model(&entity.Users{}).Count(&userCount)
	if userCount == 0 {
		log.Println("Seeding initial user data...")
		hashedPassword, _ := services.HashPassword("123456")
		birthDay, _ := time.Parse("2006-01-02", "1988-11-12")
		user := &entity.Users{
			FirstName: "Software",
			LastName:  "Analysis",
			Email:     "sa@gmail.com",
			Address:   "sut",
			Age:       80,
			Password:  hashedPassword,
			BirthDay:  birthDay,
			GenderID:  1,
		}
		db.Create(user)
		log.Println("User data seeding completed.")
	} else {
		log.Println("User data already exists, skipping seeding.")
	}
}
