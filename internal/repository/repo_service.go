package repository

type Service interface {
	Users() *UserRepository
}

type RepositoryService struct {
	UsersRepo UserRepository
}

func NewRepositoryService(users UserRepository) Service {
	return &RepositoryService{
		UsersRepo: users,
	}
}

func (r RepositoryService) Users() *UserRepository {
	return &r.UsersRepo
}
