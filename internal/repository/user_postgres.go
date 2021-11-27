package repository

type UserPostgres struct {
}

func NewUserPostgres() *UserPostgres {
	return &UserPostgres{}
}

func (u UserPostgres) Create(id int) int {
	return 1
}

func (h *UserPostgres) GetAll() int {
	return 1
}

func (h *UserPostgres) GetOne(id int) int {
	return 1
}

func (h *UserPostgres) Delete(id int) int {
	return 1
}

func (h *UserPostgres) Update(id int) int {
	return 1
}
