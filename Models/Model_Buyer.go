package Models

type Buyer struct {
	Id   int
	Name string
	Age  int
}

func New_Buyer() *Buyer {
	return &Buyer{}
}
