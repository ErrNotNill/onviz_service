package LEADS

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"onviz/DB"
	"onviz/deals"
	"strconv"
	"strings"
	"time"
)

type Lead struct {
	Id       string `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Link     string `json:"link,omitempty"`
	Status   string `json:"status,omitempty"`
	Assigned string `json:"assigned,omitempty"`
}

type Leads struct {
	ID                string `json:"ID,omitempty"`
	ResponsibleID     int    `json:"responsibleID,omitempty"`
	Title             string `json:"title,omitempty"`
	Name              string `json:"name,omitempty"`
	Phone             string `json:"phone,omitempty"`
	DateCreate        string `json:"dateCreate,omitempty"`
	SourceId          string `json:"sourceId,omitempty"`
	SourceDescription string `json:"sourceDescription,omitempty"`
	AssignedByLead    string `json:"assignedByLead"`
}

func GetLeadsAll(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Db.Query(`select ID, COALESCE(ResponsibleID,0), COALESCE(Title, ''), 
       COALESCE(Name,''), COALESCE(Phone,''), COALESCE(DateCreate,''), 
       COALESCE(SourceId,''), COALESCE(SourceDescription,''), COALESCE(AssignedByLead,'')
from Leads`)
	if err != nil {
		fmt.Println("cant get rows")
		log.Println(err)
	}
	defer rows.Close()
	products := []Leads{}

	for rows.Next() {
		//fmt.Println("started cycle from table")
		p := Leads{}
		err := rows.Scan(&p.ID, &p.ResponsibleID, &p.Title, &p.Name, &p.Phone, &p.DateCreate, &p.SourceId,
			&p.SourceDescription, &p.AssignedByLead)
		if err != nil {
			fmt.Println("i cant scan this")
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	//fmt.Fprintf(w, "%v\n", products)
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
	} // Parses the request body
	leadId := r.Form.Get("lead_id")
	name := r.Form.Get("name") // x will be "" if parameter is not set
	phone := r.Form.Get("phone")
	dateCreate := r.Form.Get("date_create")
	sourceId := r.Form.Get("source_id")
	sourceDescription := r.Form.Get("source_description")
	title := r.Form.Get("title")
	email := r.Form.Get("email")
	formName := r.Form.Get("formName")
	//DB.LeadTestAdd(leadId)
	fmt.Println("lead added")

	//responsibleID := 13983 //Admin
	assignedNotFormatted := r.Form.Get("assigned")
	assignedFormatted := strings.Split(assignedNotFormatted, "_")
	assignedByLead := assignedFormatted[1]

	fmt.Println("id: ", leadId, "title: ", title, "name: ", name, "phone: ", phone, "email: ", email, "date_create: ", dateCreate, "source_id: ", sourceId, "source_desc: ", sourceDescription, "assigned: ", assignedByLead, "formName: ", formName)

	DB.AddLeadToDB(leadId, title, name, phone, email, dateCreate, sourceId, sourceDescription, assignedByLead, formName)

	formattedLeadId, _ := strconv.Atoi(leadId)
	fmt.Println("FORMATTEDLeadID", formattedLeadId)
	//tasks.AddTaskToLead("✅ НОВАЯ ЗАЯВКА!", assignedByLead, formattedLeadId)

	/*jsnm, err := json.Marshal(newreq)
	if err != nil {
		log.Println("Error to convert json fields from struct")
	}*/

	fmt.Println("lead added")
	/*defer func(Body io.ReadCloser) {
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
	var lead Lead

	err = json.Unmarshal(body, &lead)
	if err != nil {
		fmt.Println("error unmarshall")
		return
	}
	newData, err := json.Marshal(lead)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}
	fmt.Println("LEAD :", lead)

	//DB.LeadCollectToDb(lead.Id, lead.Title, lead.Link, lead.Status, lead.Assigned)
	fmt.Println("lead added")*/

}

func GetAllFromLead(db *sql.DB) {
	rows, err := db.Query("select * from Leads")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	products := []Lead{}

	for rows.Next() {
		p := Lead{}
		err := rows.Scan(&p.Id, &p.Title, &p.Link, &p.Status, &p.Assigned)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.Id, p.Title, p.Link, p.Status, p.Assigned)
	}
}

//maybe need to use fields in crm.deal.list parameters

func DealerDealAdded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parse form")
	} // Parses the request body
	dealId := r.Form.Get("deal_id")
	name := r.Form.Get("name") // x will be "" if parameter is not set
	phone := r.Form.Get("phone")
	email := r.Form.Get("email")
	dateCreate := r.Form.Get("date_create")
	sourceId := r.Form.Get("source_id")
	sourceDescription := r.Form.Get("source_description")
	title := r.Form.Get("title")

	//DB.LeadTestAdd(dealId)
	fmt.Println("lead added")

	//responsibleID := 13983 //Admin
	assignedByLead := r.Form.Get("assigned")

	fmt.Println("id: ", dealId, "title: ", title, "name: ", name, "phone: ", phone, "email: ", email, "date_create: ", dateCreate, "source_id: ", sourceId, "source_desc: ", sourceDescription, "assigned: ", assignedByLead)

	time.Now()

	deals.AddedDealToDB(dealId, title, name, phone, email, dateCreate, sourceId, sourceDescription, assignedByLead)

	/*jsnm, err := json.Marshal(newreq)
	if err != nil {
		log.Println("Error to convert json fields from struct")
	}*/

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
	var lead Lead

	err = json.Unmarshal(body, &lead)
	if err != nil {
		fmt.Println("error unmarshall")
		return
	}
	newData, err := json.Marshal(lead)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}
	fmt.Println("LEAD :", lead)

	fmt.Println("error lead collect")
	//DB.LeadCollectToDb(lead.Id, lead.Title, lead.Link, lead.Status, lead.Assigned)
	fmt.Println("lead added")
}

func GetLeads(w http.ResponseWriter, r *http.Request) {

	req, err := http.Get("https://onviz.bitrix24.ru/rest/13938/xldrq7vlqx5suw13/crm.lead.list")
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

	//var lead Lead

	/*err = json.Unmarshal(body, &lead)
	if err != nil {
		fmt.Println("error unmarshall")
		return
	}
	newData, err := json.Marshal(lead)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}*/
}
