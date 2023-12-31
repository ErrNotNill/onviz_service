package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type AutoGenerated struct {
	Result struct {
		Task struct {
			ID    string `json:"id"`
			Title string `json:"title"`
		}
	}
}

var BitrixWebHookTask = os.Getenv("WEBHOOK_TASK")

func AddTaskToDeal(title string, responsibleId int, DealId int) {
	newReq := fmt.Sprintf(`{"fields":{"TITLE":"%v", "RESPONSIBLE_ID":%v, "UF_CRM_TASK":{"UF_CRM_TASK":"L_%v"}}}`, title, responsibleId, DealId)
	tr := bytes.NewReader([]byte(newReq))
	_, err := http.Post(BitrixWebHookTask, "application/json", tr) //nolint
	if err != nil {
		log.Println("Error http:post request to Bitrix24")
	}
}

//tasks.task.add
//{fields:{TITLE:'task for test', RESPONSIBLE_ID:13938, UF_CRM_TASK:{UF_CRM_TASK:'L_105528'}}}

func AddTaskToLead(title string, responsibleId string, leadId int) {
	newReq := fmt.Sprintf(`{"fields":{"TITLE":"%v", "RESPONSIBLE_ID":%v, "UF_CRM_TASK":{"UF_CRM_TASK":"L_%v"}}}`, title, responsibleId, leadId)
	fmt.Println("URL:>", BitrixWebHookTask)

	var jsonStr = []byte(newReq)
	req, err := http.NewRequest("POST", BitrixWebHookTask, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := io.ReadAll(resp.Body)

	var task AutoGenerated
	err = json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println("error unmarshall")
		return
	}
	newData, err := json.Marshal(task)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("newData: ", string(newData))
	}
	fmt.Println("LEAD :", task)

	fmt.Println("error lead collect")
	//DB.LeadCollectToDb(lead.Id, lead.Title, lead.Link, lead.Status, lead.Assigned)
	fmt.Println("lead added")

}
