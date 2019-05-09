package models

type Employee struct {
	tableName    struct{} `sql:"resident.employees,alias:employees"`
	ID           string   `json:"Id" sql:"id"`
	DepartmentId string   `json:"DepartmentId" sql:"department_id"`
	Name         string   `json:"Name" sql:"name"`
	Mobile       string   `json:"Mobile" sql:"mobile"`
	AccountId    string   `json:"AccountId" sql:"account_id"`
	Address      string   `json:"Address" sql:"address"`
	CreatedBy    string   `json:"Created_by" sql:"created_by"`
	Role         int      `json:"Role" sql:"role"`
	Status       int      `json:"Status" sql:"status"`
}
