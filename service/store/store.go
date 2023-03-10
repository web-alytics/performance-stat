//go:generate mockgen --destination=./mocks/${GOFILE} --package=mocks --source=${GOFILE}
package store

var defaultUser string = "defaultUser"

type Repository interface {
	GetByID(id int) ([]byte, error)
	PutData() error
	GetAll() error
}

type storeService struct {
	store Repository
}

func (s *storeService) RetreiveNameByID(id int) (string, error) {
	dataBytes, err := s.store.GetByID(id)
	if err != nil {
		return defaultUser, err
	}
	return string(dataBytes), nil

}
