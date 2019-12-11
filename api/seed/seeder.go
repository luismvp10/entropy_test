package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/luismvp10/entropy_test/api/models"
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

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
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

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}, &models.Contact{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Contact{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Contact{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

		contacts[i].UserID = users[i].ID
		err = db.Debug().Model(&models.Contact{}).Create(&contacts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

	}
}
