package service

import (
	"fmt"
	"log"
	"onviz/DB"
)

func TransactionDeviceToDb(devices ResponseDevices) {

	tx, err := DB.Db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into ResultItems table
	resultItemStmt, err := tx.Prepare(`INSERT INTO UserDevice (
		UID, CreateTime, UpdateTime,
                        Name, ActiveTime, BizType,
                        Category, Icon, IP, LocalKey,
                        Online, OwnerID, ProductID, ProductName, Sub, TimeZone, UUID
                        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,?)
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer resultItemStmt.Close()

	for _, val := range devices.Result {
		_, err := resultItemStmt.Exec(
			val.UID, val.CreateTime, val.UpdateTime, val.Name, val.ActiveTime,
			val.BizType, val.Category, val.Icon, val.IP, val.LocalKey, val.Online, val.OwnerID, val.ProductID, val.ProductName,
			val.Sub, val.TimeZone, val.UUID,
		)
		if err != nil {
			log.Fatal(err)
		}
		for _, status := range val.Status {
			_, err = resultItemStmt.Exec(
				status.Code, status.Value,
			)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func TransactionTuyaUsersToDb(users TuyaUsers) {
	fmt.Println("TRANSACTION WAS STARTED")
	if users.Result.List == nil {
		log.Println("List of users is nil")
		return
	}

	if DB.Db == nil {
		log.Println("Database connection is nil")
		return
	}

	tx, err := DB.Db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return
	}

	statusItemStmt, err := tx.Prepare(`INSERT IGNORE INTO TuyaUsers (
    create_time, email, mobile, uid, update_time, username
) VALUES (?, ?, ?, ?, ?, ?);`)
	if err != nil {
		log.Println("Failed to prepare transaction:", err)
		tx.Rollback()
		return
	}
	defer statusItemStmt.Close()

	// Loop through your data and insert into the database
	for _, val := range users.Result.List {
		_, err := statusItemStmt.Exec(
			val.CreateTime, val.Email, val.Mobile, val.UID, val.UpdateTime, val.Username,
		)
		if err != nil {
			log.Println("Failed to insert into TuyaUsers:", err)
			tx.Rollback()
			return
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction:", err)
		tx.Rollback()
		return
	}
	fmt.Println("TRANSACTION WAS ENDED")
}
