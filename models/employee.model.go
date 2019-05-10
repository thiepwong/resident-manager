package models

type Employee struct {
	tableName    struct{} `sql:"resident.employee,alias:employee"`
	ID           string   `json:"Id" sql:"id"`
	DepartmentId string   `json:"DepartmentId" sql:"department_id"`
	Name         string   `json:"Name" sql:"name"`
	Mobile       string   `json:"Mobile" sql:"mobile"`
	AccountId    string   `json:"AccountId" sql:"account_id"`
	Address      string   `json:"Address" sql:"address"`
	CreatedBy    string   `json:"CreatedBy" sql:"created_by"`
	Role         int      `json:"Role" sql:"role"`
	Status       int      `json:"Status" sql:"status"`
}

type EmployeeModel struct {
	tableName    struct{}    `sql:"resident.employee,alias:employee"`
	ID           string      `json:"Id" sql:"id"`
	DepartmentId string      `json:"DepartmentId" sql:"department_id"`
	Department   *Department `json:"Department"`
	Name         string      `json:"Name" sql:"name"`
	Mobile       string      `json:"Mobile" sql:"mobile"`
	AccountId    string      `json:"AccountId" sql:"account_id"`
	Address      string      `json:"Address" sql:"address"`
	CreatedBy    string      `json:"CreatedBy" sql:"created_by"`
	Role         int         `json:"Role" sql:"role"`
	Status       int         `json:"Status" sql:"status"`
}
