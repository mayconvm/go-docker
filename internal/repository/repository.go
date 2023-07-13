package repository

type Repository interface {
	Save()
	GetOne()
	Search()
}
