package models

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
	Email             string `json:"email"`
	FormName          string `json:"formname"`
}
