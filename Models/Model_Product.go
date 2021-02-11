package Models

type Product struct {
	Id    int
	Name  string
	Price float32
}

func New_Product() *Product {
	return &Product{}
}
