package product

import "api/src/product/domain"

var ProductMemRepo = ProductInMemoryRepo{ProductList: make([]domain.Product, 0), IsErr: false}
