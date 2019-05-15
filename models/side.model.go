package models

type Side struct {
	tableName struct{} `sql:"resident.side,alias:side"`
	Id        string   `json:"Id" sql:"id"`
	Name      string   `json:"Name" sql:"name"`
	Address   string   `json:"Address" sql:"address"`
	Ip        string   `json:"Ip" sql:"ip_address"`
	Cover     string   `json:"Cover" sql:"cover_photos"`
	Hotline   string   `json:"Hotline" sql:"hotline"`
}
