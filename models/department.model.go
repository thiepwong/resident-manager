package models

type Department struct {
	tableName struct{} `sql:"resident.department,alias:department"`
	ID        string   `json:"Id" sql:"id"`
	Name      string   `json:"Name" sql:"name"`
	SideId    string   `json:"SideId" sql:"side_id"`
}
