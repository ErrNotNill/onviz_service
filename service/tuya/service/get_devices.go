package service

import (
	"fmt"
	"onviz/DB"
)

func GetUidFromTuyaUsersByEmail(userEmail string) string {
	rows, err := DB.Db.Query(`select uid from TuyaUsers where email = ?;`, userEmail)
	if err != nil {
		fmt.Println("cant get rows")
	}
	defer rows.Close()
	var uid string

	for rows.Next() {
		err := rows.Scan(&uid)
		if err != nil {
			fmt.Println("i cant scan this")
			continue
		}
	}
	return uid
}
