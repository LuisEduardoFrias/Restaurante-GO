package Models

type Transaction struct {
	Id       int
	BuyerId  int
	IP       string
	Device   string
	Producto []int
}

func New_Transaction() *Transaction {
	return &Transaction{}
}
