package repositories

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/thiepwong/resident-manager/models"
)

type EmployeeRepository interface {
	Register(*models.Employee) (*models.Employee, error)
	GetById(string) (*models.EmployeeModel, error)
	GetAll() (*[]models.EmployeeModel, error)
	GetPagination(bool, string, int, int, int, string) (*[]models.EmployeeModel, error)
	Update(*models.Employee) (*models.Employee, error)
	GetRole(string) (*models.EmployeeModel, error)
	Check(string) (bool, error)
	Delete(string) (bool, error)
}

type employeeRepositoryContext struct {
	db *pg.DB
}

func NewEmployeeRepository(db *pg.DB) *employeeRepositoryContext {
	return &employeeRepositoryContext{
		db: db,
	}
}

func (emp *employeeRepositoryContext) Register(employee *models.Employee) (*models.Employee, error) {

	return employee, emp.db.Insert(employee)
	//fmt.Printf("Du lieu da goi: %s", e.Error())
	//return employee, e
}

func (emp *employeeRepositoryContext) GetById(id string) (*models.EmployeeModel, error) {
	var _emp models.EmployeeModel
	emp.db.Model(&_emp).Column("employee.*", "Department", "Department.Side").Where("employee.id=?", id).Select()
	return &_emp, nil
	//.Join("JOIN department as d on d.id = department_id").
}

func (emp *employeeRepositoryContext) GetAll() (*[]models.EmployeeModel, error) {

	var _employee []models.EmployeeModel
	emp.db.Model(&_employee).Column("employee.*", "Department").Order("id ASC").Limit(20).Select()
	// emp.db.Model(&_employee).Join("JOIN resident.department as d ON d.id = employee.department_id").Order("id ASC").Limit(20).Select()
	//	emp.db.Query(pg.Scan(_employee), "select * from resident.employee")
	return &_employee, nil
}

func (emp *employeeRepositoryContext) GetPagination(isDepartment bool, requestId string, role int, offset int, limit int, orderBy string) (*[]models.EmployeeModel, error) {
	var _employee []models.EmployeeModel
	if orderBy == "" {
		orderBy = "id DESC"
	}

	if role != 0 {
		if isDepartment == false {
			emp.db.Model(&_employee).Column("employee.*", "Department").Where("department.side_id=? and employee.role=?", requestId, role).Order(orderBy).Limit(limit).Offset(offset).Select()
		} else {
			emp.db.Model(&_employee).Column("employee.*", "Department").Where("employee.department_id=? and employee.role=?", requestId, role).Order(orderBy).Limit(limit).Offset(offset).Select()
		}
	} else {

		if isDepartment == false {
			emp.db.Model(&_employee).Column("employee.*", "Department").Where("department.side_id=?", requestId).Order(orderBy).Limit(limit).Offset(offset).Select()
		} else {
			emp.db.Model(&_employee).Column("employee.*", "Department").Where("employee.department_id=?", requestId).Order(orderBy).Limit(limit).Offset(offset).Select()
		}
	}

	return &_employee, nil
}

func (emp *employeeRepositoryContext) Update(employee *models.Employee) (*models.Employee, error) {
	if employee.ID == "" {
		return nil, errors.New("Id is required!")
	}
	return employee, emp.db.Update(employee)
}

func (emp *employeeRepositoryContext) GetRole(accountId string) (*models.EmployeeModel, error) {
	var _employee models.EmployeeModel
	if accountId == "" {
		return nil, errors.New("Account Id is required!")
	}

	emp.db.Model(&_employee).Column("employee.*", "Department", "Department.Side").Where("employee.account_id =?", accountId).Select()

	return &_employee, nil

}

func (emp *employeeRepositoryContext) Check(mobile string) (bool, error) {
	var _emp models.Employee
	emp.db.Model(&_emp).Column("id").Where("mobile=?", mobile).Where("status>?", 0).Select()
	if _emp.ID == "" {
		return false, errors.New("This mobile number is not registered!")
	}
	return true, nil
}

func (emp *employeeRepositoryContext) Delete(id string) (bool, error) {
	var _emp models.Employee
	r, e := emp.db.Model(&_emp).Where("id= ?", id).Delete()
	if e != nil {
		return false, e
	}
	if r != nil {

	}

	return true, nil
}
