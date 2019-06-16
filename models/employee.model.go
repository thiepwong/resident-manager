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

type SignUpModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	ApiVersion    string      `json:"api_version"`
	Errors        interface{} `json:"errors"`
	StatusCode    int         `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	SystemTime    string      `json:"system_time"`
	Data          interface{} `json:"success"`
}

type Activate struct {
	Mobile string `json:"mobile"`
	Code   string `json:"otpCode"`
}
