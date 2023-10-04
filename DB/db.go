package DB

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var (
	MysqlUrl = os.Getenv("URL_MYSQL")
)

var Db *sqlx.DB //global variable

func InitDB(url string) (err error) {
	Db, err = sqlx.Connect("mysql", url)
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
	fmt.Println(result.RowsAffected())
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
	fmt.Println(result.RowsAffected())
}
