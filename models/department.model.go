package models

type Department struct {
	tableName struct{} `sql:"resident.department,alias:department"`
	Id        string   `json:"Id" sql:"id"`
	Name      string   `json:"Name" sql:"name"`
	SideId    string   `json:"SideId" sql:"side_id"`
}

type DepartmentModel struct {
	tableName struct{} `sql:"resident.department,alias:department"`
	Id        string   `json:"Id" sql:"id"`
	Name      string   `json:"Name" sql:"name"`
	SideId    string   `json:"SideId" sql:"side_id"`
	Side      Side
}
