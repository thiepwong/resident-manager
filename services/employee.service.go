package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type EmployeeService interface {
	Register(string, string, string, string, string, int, string) *models.Employee
	GetById(string) *models.Employee
	GetAll() *[]models.EmployeeModel
	GetList(int, int, string) (*[]models.EmployeeModel, error)
	Update(string, string, string, string) *models.Employee
	Signin(string, string, string) (map[string]interface{}, error)
}

type employeeServiceImp struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeServiceImp{

		employeeRepo: repo,
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

func (s *employeeServiceImp) GetById(id string) *models.Employee {
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

func (s *employeeServiceImp) GetList(pageIndex int, pageSize int, orderBy string) (*[]models.EmployeeModel, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	rs, e := s.employeeRepo.GetPagination(offset, pageSize, orderBy)
	if e != nil {
		return nil, e
	}
	return rs, nil
}

func (s *employeeServiceImp) Update(departmentId string, name string, mobile string, accountId string) *models.Employee {
	var _emp = &models.Employee{}
	return _emp
}

func (s *employeeServiceImp) Signin(username string, password string, system string) (map[string]interface{}, error) {

	message := map[string]interface{}{
		"username": username,
		"password": password,
		"system":   system}

	bytesRepresentation, err := json.Marshal(message)
	url := "http://171.244.49.164:3333/api/v1/accounts/sign-in?api_token=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJTUElOIFYxIiwiaWF0IjoxNTQ1MjEwNDkwNzQxLCJleHAiOjE1NDUyNDY0OTA3NDEsInN5cyI6IlBBUktJTkcifQ.MnSeQKn34b8x-yoXnnndxw1FaQf2f1Z2XDsYgYqvmOUmF0rMVsK1lWDsDSGaVelZQ7lYW3o4aFvI7MrdFGTlxj0g333z_lHYoR2YapvZyAPseLfF7NHthE72JbcAd9L6ynyjGP5sBpjQGkt5o45dppnWZQj4_5GvetsrUeSZhMQ"
	req, e := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if e != nil {
		log.Fatal("Loi roi")
		return nil, e
	}

	req.Header.Set("Content-Type", "Application/json")
	req.Header.Set("Authorization", "key=AAAAtc-5Fto:APA91bFxm1mLGKf9rGaCDu-f6K8cWOqWEO8qR9XYdkwsi4Bng75y9XxeCY6rySPIzpY1EfveXlgWIzTfpnn49TNmjj2pzq7TlcVOuNVB5fu96cDtN59RSXHvEaqIyXHEOfiYHtaSoogm")

	// Do the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	return result, nil
}
