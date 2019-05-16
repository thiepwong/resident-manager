package models

type Contact struct {
	tableName    struct{} `sql:"resident.contact,alias:contact"`
	Id           string   `json:"Id" sql:"id"`
	Name         string   `json:"Name" sql:"name"`
	DepartmentId string   `json:"DepartmentId" sql:"department_id"`
	PhoneNumber  string   `json:"PhoneNumber" sql:"phone_no"`
}

type ContactModel struct {
	tableName    struct{} `sql:"resident.contact,alias:contact"`
	Id           string   `json:"Id" sql:"id"`
	Name         string   `json:"Name" sql:"name"`
	DepartmentId string   `json:"DepartmentId" sql:"department_id"`
	Department   Department
	PhoneNumber  string `json:"PhoneNumber" sql:"phone_no"`
}
