package deals

import (
	"database/sql"
	"fmt"
	"log"
)

func AddedDealToDB(Id, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead string) {
	db, err := sql.Open("mysql", "mysql:mysql@/Onviz")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("err db close")
		}
	}(db)

	fmt.Println("db connected")

	err = db.Ping()
	if err != nil {
		fmt.Println("db not pinged")
		return
	}
	result, err := db.Exec(`insert into Deals (ID, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead) values (?, ?, ?,?, ?, ?, ?,?,?)`,
		Id, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}

	//tasks.AddTaskToDeal()

	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}
