package Models

type Transaction struct {
	Id       int32
	BuyerId  int32
	IP       string
	Device   string
	Producto []int32
}
