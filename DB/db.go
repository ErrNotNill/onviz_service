package DB

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB //global variable

//mysql:mysql@/Onviz //on vps
//mysql:mysql@tcp(45.141.79.120:3306)/Onviz //locally

func InitDB() (err error) {
	var dataSourceName = "mysqld:mysql@tcp(45.141.79.120:3306)/Onviz"
	Db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		fmt.Println("not connected")
		return
	}
	err = Db.Ping()
	return
}

func LeadTestAdd(Id string) {
	result, err := Db.Exec(`insert into Leads (ID) values (?)`,
		Id)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected()) // количество доба
}

func AddLeadToDB(Id, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead, FormName string) {
	fmt.Println("db connected")

	result, err := Db.Exec(`insert into Leads (ID, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead, FormName ) values (?, ?, ?,?, ?, ?, ?,?,?,?)`,
		Id, Title, Name, Phone, Email, DateCreate, SourceId, SourceDescription, AssignedByLead, FormName)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}

/*func LeadCollectToDb(id, title, link, status, assigned string) {
	db, err := ConnectToDb()

	result, err := db.Exec("insert into Leads (id, title, link, status, assigned) values ($1, $2, $3, $4, $5)",
		id, title, link, status, assigned)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected()) // количество добавленных строк
}*/
