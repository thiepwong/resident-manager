package services

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/thiepwong/resident-manager/models"
	"github.com/thiepwong/resident-manager/repositories"
)

type ResidentService interface {
	GetById(string) (*models.ResidentRoom, error)
	GetList(string, string, int, int, string) (*[]models.ResidentRoom, error)
	Add(*models.ResidentModel) (*models.ResidentRoom, error)
	Update(*models.ResidentModel) (*models.ResidentRoom, error)
	Delete(residentRoomId string) (bool, error)
}

type residentServiceImp struct {
	residentRepo repositories.ResidentRepository
}

func NewResidentService(repo repositories.ResidentRepository) ResidentService {
	return &residentServiceImp{
		residentRepo: repo,
	}
}

func (s *residentServiceImp) GetById(id string) (*models.ResidentRoom, error) {
	return s.residentRepo.GetById(id)
}

func (s *residentServiceImp) GetList(sideId string, search string, pageIndex int, pageSize int, pageOrder string) (*[]models.ResidentRoom, error) {
	if pageIndex < 1 || pageSize < 1 {
		return nil, errors.New("Page index or page Size is invalid! Please check!")
	}
	var offset int
	offset = (pageIndex - 1) * pageSize
	return s.residentRepo.GetPagination(sideId, search, offset, pageSize, pageOrder)
}

func (s *residentServiceImp) Add(res *models.ResidentModel) (*models.ResidentRoom, error) {

	if res == nil {
		return nil, errors.New("Invalid infomation for resident!")
	}

	var _resident = models.Resident{
		Id:          uuid.Must(uuid.NewV4()).String(),
		AccountId:   res.Resident.AccountId,
		Email:       res.Resident.Email,
		PhoneNumber: res.Resident.PhoneNumber,
		FullName:    res.Resident.FullName,
	}

	var _residentRoom = models.ResidentRoom{
		Id:         uuid.Must(uuid.NewV4()).String(),
		ResidentId: _resident.Id,
		RoomId:     res.RoomId,
		Active:     true,
	}

	return s.residentRepo.Add(&_resident, &_residentRoom)
}

func (s *residentServiceImp) Update(res *models.ResidentModel) (*models.ResidentRoom, error) {

	if res == nil {
		return nil, errors.New("Invalid infomation for resident!")
	}

	if res.Resident.Id == "" {
		return nil, errors.New("Invalid infomation for resident, no ID info!")
	}

	var _resident = models.Resident{
		Id:          res.Resident.Id,
		AccountId:   res.Resident.AccountId,
		Email:       res.Resident.Email,
		PhoneNumber: res.Resident.PhoneNumber,
		FullName:    res.Resident.FullName,
	}

	var _residentRoom = models.ResidentRoom{
		Id:     res.ResidentRoomId,
		RoomId: res.RoomId,
	}

	return s.residentRepo.Update(&_resident, &_residentRoom)
}

func (s *residentServiceImp) Delete(resRId string) (bool, error) {
	return s.residentRepo.Delete(resRId)
}
