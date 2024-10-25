package productdto

// type ProductDTOResponse struct {
// 	ProductName string    `json:"productName"`
// 	Price  int64    `json:"price"`
// }

type ProductDTOResponse struct {
	json        struct{} `naming:"snake_case"`
	ProductName string
	Price       int64
}
