package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("project-sa65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema ------------------------------------------------------
	database.AutoMigrate(
		&Room_Price{},
		&Room{},
		&Student{},
		&Semester{},
		&Employee{},
		&Booking{},
		&Payment_Bill{},
	)

	db = database

	// Room_Price Data ------------------------------------------------------
	// ค่อยไป inner join ตาราง Room_Price ของเพื่อน
	db.Model(&Room_Price{}).Create(&Room_Price{
		Room_Price: 4500,
	})
	db.Model(&Room_Price{}).Create(&Room_Price{
		Room_Price: 5000,
	})

	// Room Data ------------------------------------------------------
	db.Model(&Room{}).Create(&Room{
		Room_PriceID: 1,
	})
	db.Model(&Room{}).Create(&Room{
		Room_PriceID: 2,
	})

	// Student Data ------------------------------------------------------
	db.Model(&Student{}).Create(&Student{
		Student_Number: "B6348799",
	})
	db.Model(&Student{}).Create(&Student{
		Student_Number: "B6345728",
	})

	// Employee Data ------------------------------------------------------
	Employee_Personal_ID, err := bcrypt.GenerateFromPassword([]byte("111"), 14)

	db.Model(&Employee{}).Create(&Employee{
		Employee_Personal_ID: string(Employee_Personal_ID),
		Employee_Name:        "Jacky",
		Employee_Email:       "jacky@gmail.com",
	})
	db.Model(&Employee{}).Create(&Employee{
		Employee_Personal_ID: string(Employee_Personal_ID),
		Employee_Name:        "Arisa",
		Employee_Email:       "arisa@gmail.com",
	})
	db.Model(&Employee{}).Create(&Employee{
		Employee_Personal_ID: string(Employee_Personal_ID),
		Employee_Name:        "Admin",
		Employee_Email:       "@min.com",
	})

	// Booking Data ------------------------------------------------------
	booking1 := Booking{
		Check_In_Date: time.Date(2020, time.May, 10, 00, 00, 00, 0, time.UTC),
		RoomID:        1,
		StudentID:     0001,
	}
	db.Model(&Booking{}).Create(&booking1)

	booking2 := Booking{
		Check_In_Date: time.Date(2020, time.May, 10, 00, 00, 00, 0, time.UTC),
		RoomID:        2,
		StudentID:     0002,
	}
	db.Model(&Booking{}).Create(&booking2)

	// Semester Data ------------------------------------------------------
	semester1 := Semester{
		Semester:		"1/2564",
	}
	db.Model(&Semester{}).Create(&semester1)

	semester2 := Semester{
		Semester:		"2/2564",
	}
	db.Model(&Semester{}).Create(&semester2)

	semester3 := Semester{
		Semester:		"3/2564",
	}
	db.Model(&Semester{}).Create(&semester3)

	semester4 := Semester{
		Semester:		"1/2565",
	}
	db.Model(&Semester{}).Create(&semester4)
	
	//
	// === Query
	//

	/* var target User
	db.Model(&User{}).Find(&target, db.Where("email = ?", "tanapon@gmail.com"))

	var watchedPlaylist Playlist
	db.Model(&Playlist{}).Find(&watchedPlaylist, db.Where("title = ? and owner_id = ?", "Watched", target.ID))

	var watchedList []*WatchVideo
	db.Model(&WatchVideo{}).
		Joins("Playlist").
		Joins("Resolution").
		Joins("Video").
		Find(&watchedList, db.Where("playlist_id = ?", watchedPlaylist.ID))

	for _, wl := range watchedList {
		fmt.Printf("Watch Video: %v\n", wl.ID)
		fmt.Printf("%v\n", wl.Playlist.Title)
		fmt.Printf("%v\n", wl.Resolution.Value)
		fmt.Printf("%v\n", wl.Video.Name)
		fmt.Println("====")
	} */
}
