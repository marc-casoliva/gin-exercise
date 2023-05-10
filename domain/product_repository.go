package domain

type ProductRepository interface {
	Save(p Product) error
	RetrieveById(id string) (Product, error)
}
