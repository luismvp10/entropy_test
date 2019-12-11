package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/luismvp10/entropy_test/api/models"
	"log"
)

var users = []models.User{
	models.User{
		Nombre:   "Luis Valencia",
		Edad:     20,
		Foto:     "image1.png",
		Email:    "luismvalenp@gmail.com",
		Password: "12345",
	},
	models.User{
		Nombre:   "Juan Olivares",
		Edad:     21,
		Foto:     "image2.png",
		Email:    "juanol@gmail.com",
		Password: "root",
	},
}

var contacts = []models.Contact{
	models.Contact{
		Nombre:    "Juan",
		Apodo:     "Juancho",
		Email:     "juanol@gmail.com",
		NumeroTel: "5540368914",
		Direccion: "Manzana 85",
	},

	models.Contact{
		Nombre:    "Juan",
		Apodo:     "Juancho",
		Email:     "juanol@gmail.com",
		NumeroTel: "5540368914",
		Direccion: "Manzana 86",
	},
}

var groups = []models.Group{
	models.Group{

		Nombre: "UnoPrueba",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Contact{}, &models.GroupUser{}, &models.Group{}, &models.Message{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Contact{}, &models.GroupUser{}, &models.Group{}, &models.Message{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Contact{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.GroupUser{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.GroupUser{}).AddForeignKey("group_id", "groups(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		contacts[i].UserID = users[i].ID
		err = db.Debug().Model(&models.Contact{}).Create(&contacts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

	}
}
