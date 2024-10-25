package walletdto

// type WalletDTOResponse struct {
// 	Username string    `json:"username"`
// 	Wallet  int64    `json:"wallet"`
// }

type WalletDTOResponse struct {
	json     struct{} `naming:"snake_case"`
	Username string
	Wallet   int64
}
