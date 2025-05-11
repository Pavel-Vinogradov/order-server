package repository

type Repository[T any] interface {
	Create(T) (T, error)
	Update(T) (T, error)
	Delete(T) (bool, error)
	FindAll() ([]T, error)
	FindByID(int64) (T, error)
}
