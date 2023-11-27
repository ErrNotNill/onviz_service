package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"onviz/DB"
	"onviz/service/bitrix/models"
	"os"
	"strconv"
	"strings"
)

func LeadTestAdd(Id string) {
	result, err := DB.Db.Exec(`insert into Leads (ID) values (?)`,
		Id)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}

func AddLeadToDB(Id, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead, FormName string) {
	fmt.Println("db connected")

	result, err := DB.Db.Exec(`insert into Leads (ID, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead, FormName ) values (?, ?, ?,?, ?, ?, ?,?,?,?)`,
		Id, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead, FormName)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}

func GetLeadsAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	rows, err := DB.Db.Query(`select ID, COALESCE(ResponsibleID,0), COALESCE(Title, ''), 
       COALESCE(Name,''), COALESCE(Phone,''), COALESCE(DateCreate,''), 
       COALESCE(SourceId,''), COALESCE(SourceDescription,''), COALESCE(AssignedByLead,''), COALESCE(Email,''),
       COALESCE(FormName,'')
from Leads`)
	if err != nil {
		fmt.Println("cant get rows")
	}
	defer rows.Close()
	products := []models.Leads{}

	for rows.Next() {
		p := models.Leads{}
		err := rows.Scan(&p.ID, &p.ResponsibleID, &p.Title, &p.Name, &p.Phone, &p.DateCreate, &p.SourceId,
			&p.SourceDescription, &p.AssignedByLead, &p.Email, &p.FormName)
		if err != nil {
			fmt.Println("i cant scan this")
			continue
		}
		products = append(products, p)
	}
	err = json.NewEncoder(w).Encode(&products)
	if err != nil {
		log.Println("i cant encode")
	}

	/*var count int
	qry, err := DB.Db.Query(`SELECT COUNT(*) FROM Leads WHERE ID != ' '`)
	if err != nil {
		fmt.Println(`error query`)
	}
	for qry.Next() {
		if err := qry.Scan(&count); err != nil {
			fmt.Println(`error scan`)
		}
	}
	fmt.Fprintf(w, "COUNT is: %v\n", count)

	data, err := json.MarshalIndent(&products, "", "    ")
	if err != nil {
		fmt.Println("i cant convert to json")
	}
	if r.Method == "GET" {
		w.Write(data)
	}*/

	/*for _, v := range products {
		fmt.Fprintf(w, "%v\n", v)
	}*/

	/*tmpl, _ := template.ParseFiles("templates/index.html")
	err = tmpl.Execute(w, products)
	if err != nil {
		fmt.Println("i cant parse template")
		return
	}*/
}

func LeadsAdd(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parse form")
	}
	leadId := r.Form.Get("lead_id")
	name := r.Form.Get("name") // x will be "" if parameter is not set
	phone := r.Form.Get("phone")
	dateCreate := r.Form.Get("date_create")
	sourceId := r.Form.Get("source_id")
	sourceDescription := r.Form.Get("source_description")
	title := r.Form.Get("title")
	email := r.Form.Get("email")
	formName := r.Form.Get("formname")

	assignedNotFormatted := r.Form.Get("assigned")
	assignedFormatted := strings.Split(assignedNotFormatted, "_")
	assignedByLead := assignedFormatted[1]

	fmt.Println("id: ", leadId, "title: ", title, "name: ", name, "phone: ", phone, "email: ", email, "date_create: ", dateCreate, "source_id: ", sourceId, "source_desc: ", sourceDescription, "assigned: ", assignedByLead, "formName: ", formName)

	AddLeadToDB(leadId, title, name, phone, email, dateCreate, sourceId, sourceDescription, assignedByLead, formName)

	formattedLeadId, _ := strconv.Atoi(leadId)
	fmt.Println("FORMATTEDLeadID", formattedLeadId)
	//tasks.AddTaskToLead("✅ НОВАЯ ЗАЯВКА!", assignedByLead, formattedLeadId)
	fmt.Println("lead added")
}

func GetAllFromLead(db *sql.DB) {
	rows, err := db.Query("select * from Leads")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	products := []models.Lead{}

	for rows.Next() {
		p := models.Lead{}
		err := rows.Scan(&p.Id, &p.Title, &p.Link, &p.Status, &p.Assigned)
		if err != nil {
			fmt.Println("i cant scan this")
			fmt.Println(err.Error())
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.Id, p.Title, p.Link, p.Status, p.Assigned)
	}
}

var WebHookLeads = os.Getenv("WEBHOOK_LEADS")

func GetLeads(w http.ResponseWriter, r *http.Request) {
	req, err := http.Get(WebHookLeads)
	if err != nil {
		log.Println("Error http:post request to Bitrix24")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("error body close")
		}
	}(r.Body)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error read body")
	}
	fmt.Println("BODY: ", string(body))
}
