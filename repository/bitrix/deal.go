package bitrix

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"onviz/models"
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

func DealerDealAdded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parse form")
	}
	dealId := r.Form.Get("deal_id")
	name := r.Form.Get("name") // x will be "" if parameter is not set
	phone := r.Form.Get("phone")
	email := r.Form.Get("email")
	dateCreate := r.Form.Get("date_create")
	sourceId := r.Form.Get("source_id")
	sourceDescription := r.Form.Get("source_description")
	title := r.Form.Get("title")
	fmt.Println("lead added")

	assignedByLead := r.Form.Get("assigned")

	fmt.Println("id: ", dealId, "title: ", title, "name: ", name, "phone: ", phone, "email: ", email, "date_create: ", dateCreate, "source_id: ", sourceId, "source_desc: ", sourceDescription, "assigned: ", assignedByLead)

	AddedDealToDB(dealId, title, name, phone, email, dateCreate, sourceId, sourceDescription, assignedByLead)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("error body close")
		}
	}(r.Body)
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error read body")
	}
	fmt.Println("BODY: ", string(body))
	var lead models.Lead

	err = json.Unmarshal(body, &lead)
	if err != nil {
		fmt.Println("error unmarshall")
		return
	}
	newData, err := json.Marshal(lead)
	if err != nil {
		fmt.Println("i can't marshal lead")
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(newData))
	}
	fmt.Println("LEAD :", lead)
}
