package service

import "golang-shop/internal/repository"

type UserService struct {
	UserPostgres *repository.UserPostgres
}

func NewUserService() *UserService {
	return &UserService{
		UserPostgres: repository.NewUserPostgres(),
	}
}

func (h *UserService) Create(id int) int {
	_id := h.UserPostgres.Create(id)
	return _id
}

func (h *UserService) GetAll() int {
	_id := h.UserPostgres.GetAll()
	return _id
}

func (h *UserService) GetOne(id int) int {
	_id := h.UserPostgres.GetOne(id)
	return _id
}

func (h *UserService) Delete(id int) int {
	_id := h.UserPostgres.Delete(id)
	return _id
}

func (h *UserService) Update(id int) int {
	_id := h.UserPostgres.Update(id)
	return _id
}
