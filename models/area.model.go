package models

type Area struct {
	tableName struct{} `sql:"resident.area,alias:area"`
	ID        string   `json:"Id" sql:"id"`
	Name      string   `json:"Name" sql:"name"`
	SideId    string   `json:"SideId" sql:"side_id"`
}

type AreaModel struct {
	tableName struct{} `sql:"resident.area,alias:area"`
	ID        string   `json:"Id" sql:"id"`
	Name      string   `json:"Name" sql:"name"`
	SideId    string   `json:"SideId" sql:"side_id"`
	Side      Side
}
