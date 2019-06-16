package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/thiepwong/resident-manager/common"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type EmployeeService interface {
	Register(string, string, string, string, string, int, string) *models.Employee
	GetById(string) *models.EmployeeModel
	GetAll() *[]models.EmployeeModel
	GetList(bool, string, int, int, int, string) (*[]models.EmployeeModel, error)
	Update(string, string, string, string) *models.Employee
	Signin(string, string, string) (interface{}, error)
	SignUp(*models.SignUpModel) (interface{}, error)
	Activate(*models.Activate) (interface{}, error)
	SendOTP(string) (interface{}, error)
	GetRole(string) (*models.EmployeeModel, error)
}

type employeeServiceImp struct {
	employeeRepo repositories.EmployeeRepository
	config       *common.Config
}

func NewEmployeeService(repo repositories.EmployeeRepository, cfg *common.Config) EmployeeService {
	return &employeeServiceImp{

		employeeRepo: repo,
		config:       cfg,
	}
}

func (s *employeeServiceImp) Register(departmentId string, name string, mobile string, address string, accountId string, role int, createdBy string) *models.Employee {
	_id := uuid.Must(uuid.NewV4())
	var _emp = &models.Employee{
		ID:           _id.String(),
		AccountId:    accountId,
		Name:         name,
		Mobile:       mobile,
		DepartmentId: departmentId,
		Address:      address,
		Role:         role,
		CreatedBy:    createdBy,
		Status:       1,
	}
	rs, e := s.employeeRepo.Register(_emp)
	if e != nil {
		fmt.Print(e.Error())
		return nil
	}
	return rs
}

func (s *employeeServiceImp) GetById(id string) *models.EmployeeModel {
	rs, e := s.employeeRepo.GetById(id)
	if e != nil {
		return nil
	}
	return rs
}

func (s *employeeServiceImp) GetAll() *[]models.EmployeeModel {

	rs, e := s.employeeRepo.GetAll()
	if e != nil {
		return nil
	}
	return rs

}

func (s *employeeServiceImp) GetList(isDept bool, requestId string, role int, pageIndex int, pageSize int, orderBy string) (*[]models.EmployeeModel, error) {

	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.employeeRepo.GetPagination(isDept, requestId, role, offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *employeeServiceImp) Update(departmentId string, name string, mobile string, accountId string) *models.Employee {
	var _emp = &models.Employee{}
	return _emp
}

func (s *employeeServiceImp) Signin(username string, password string, system string) (interface{}, error) {

	message := map[string]interface{}{
		"username": username,
		"password": password,
		"system":   system,
	}

	bytesRepresentation, err := json.Marshal(message)
	url := s.config.Option.SmsUrl + "accounts/sign-in?api_token=" + s.config.Option.SmsApiToken
	req, e := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if e != nil {
		return nil, e
	}

	req.Header.Set("Content-Type", "Application/json")
	//req.Header.Set("Authorization", "key=AAAAtc-5Fto:APA91bFxm1mLGKf9rGaCDu-f6K8cWOqWEO8qR9XYdkwsi4Bng75y9XxeCY6rySPIzpY1EfveXlgWIzTfpnn49TNmjj2pzq7TlcVOuNVB5fu96cDtN59RSXHvEaqIyXHEOfiYHtaSoogm")

	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var res models.Response
	json.NewDecoder(response.Body).Decode(&res)

	if res.Errors != nil {

		_err := res.Errors.(map[string]interface{})
		if _err["code"] != "200" {
			_str := fmt.Sprintf("%s", _err["message"])
			e = errors.New(_str)
		}
	}

	return res.Data, e
}

func (s *employeeServiceImp) SignUp(m *models.SignUpModel) (interface{}, error) {
	bytesRepresentation, err := json.Marshal(m)
	url := s.config.Option.SmsUrl + "accounts/sign-up?api_token=" + s.config.Option.SmsApiToken
	req, e := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if e != nil {
		log.Fatal("Loi roi")
		return nil, e
	}
	req.Header.Set("Content-Type", "Application/json")
	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var res models.Response
	json.NewDecoder(response.Body).Decode(&res)
	if res.Errors != nil {

		_err := res.Errors.(map[string]interface{})
		if _err["code"] != "200" {
			_str := fmt.Sprintf("%s", _err["message"])
			e = errors.New(_str)
		}
	}

	return res.Data, e
}

func (s *employeeServiceImp) Activate(m *models.Activate) (interface{}, error) {
	bytesRepresentation, err := json.Marshal(m)
	url := s.config.Option.SmsUrl + "accounts/verify-account?api_token=" + s.config.Option.SmsApiToken
	req, e := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if e != nil {
		return nil, e
	}

	req.Header.Set("Content-Type", "Application/json")
	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var res models.Response
	json.NewDecoder(response.Body).Decode(&res)

	if res.Errors != nil {

		_err := res.Errors.(map[string]interface{})
		if _err["code"] != "200" {
			_str := fmt.Sprintf("%s", _err["message"])
			e = errors.New(_str)
		}
	}

	return res.Data, e
}

func (s *employeeServiceImp) SendOTP(mobile string) (interface{}, error) {
	args := map[string]interface{}{
		"mobile": mobile,
	}

	bytesRepresentation, err := json.Marshal(args)
	url := s.config.Option.SmsUrl + "accounts/send-otp?api_token=" + s.config.Option.SmsApiToken
	req, e := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if e != nil {
		return nil, e
	}

	req.Header.Set("Content-Type", "Application/json")
	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var res models.Response
	json.NewDecoder(response.Body).Decode(&res)
	if res.Errors != nil {

		_err := res.Errors.(map[string]interface{})
		if _err["code"] != "200" {
			_str := fmt.Sprintf("%s", _err["message"])
			e = errors.New(_str)
		}
	}

	return res.Data, e
}

func (s *employeeServiceImp) GetRole(accountId string) (*models.EmployeeModel, error) {
	return s.employeeRepo.GetRole(accountId)
}
