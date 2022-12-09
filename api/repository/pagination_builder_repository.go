package repository

type PaginationBuilderRepository interface {
}

type paginationBuilderRepositoryImpl struct {
	productsRepository ProductsRepository
}

func NewPaginationBuilderRepository(productsRepository ProductsRepository) *paginationBuilderRepositoryImpl  {
	return &paginationBuilderRepositoryImpl{productsRepository}
}

