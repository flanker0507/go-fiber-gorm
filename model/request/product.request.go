package request

type ProductCreateRequest struct {
	Name      string `json:"name" validate:"required"`
	Deskripsi string `json:"deskripsi" validate:"required"`
	Harga     string `json:"harga" validate:"required"`
	Stok      string `json:"stok" validate:"required"`
}

type ProductUpdateRequest struct {
	Name      string `json:"name" validate:"required"`
	Deskripsi string `json:"deskripsi" validate:"required"`
	Harga     string `json:"harga" validate:"required"`
	Stok      string `json:"stok" validate:"required"`
}
