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
	Email        string   `json:"Email" sql:"email"`
}

type EmployeeModel struct {
	tableName    struct{}         `sql:"resident.employee,alias:employee"`
	ID           string           `json:"Id" sql:"id"`
	DepartmentId string           `json:"DepartmentId" sql:"department_id"`
	Department   *DepartmentModel `json:"Department"`
	Name         string           `json:"Name" sql:"name"`
	Mobile       string           `json:"Mobile" sql:"mobile"`
	AccountId    string           `json:"AccountId" sql:"account_id"`
	Address      string           `json:"Address" sql:"address"`
	CreatedBy    string           `json:"CreatedBy" sql:"created_by"`
	Role         int              `json:"Role" sql:"role"`
	Status       int              `json:"Status" sql:"status"`
	Email        string   `json:"Email" sql:"email"`
}

type SignUpModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	SendOTP  bool 	`json:"notSendOTP"`
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

type ChangePassword struct {
	Id          string `json:"Id"`
	OldPassword string `json:"OldPassword"`
	NewPassword string `json:"NewPassword"`
}

type ResetPassword struct {
	OTP         string `json:"otpCode"`
	Mobile      string `json:"mobile"`
	NewPassword string `json:"password"`
}
