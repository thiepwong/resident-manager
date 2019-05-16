package models

type Block struct {
	tableName struct{} `sql:"resident.block,alias:block"`
	Id        string   `json:"Id" sql:"id"`
	Name      string   `json:"Name" sql:"name"`
	SideId    string   `json:"SideId" sql:"side_id"`
}

type BlockModel struct {
	tableName struct{} `sql:"resident.block,alias:block"`
	Id        string   `json:"Id" sql:"id"`
	Name      string   `json:"Name" sql:"name"`
	SideId    string   `json:"SideId" sql:"side_id"`
	Side      Side
}
