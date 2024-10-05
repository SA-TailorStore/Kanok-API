package requests

type UserRegisterRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Phone_number string `json:"phone_number" validate:"required"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UsernameRequest struct {
	Username string `json:"username" validate:"required"`
}

type UserJWTRequest struct {
	Token string `json:"token" validate:"required"`
}

type UserIDRequest struct {
	User_id string `json:"user_id" validate:"required"`
}

type OrderIDRequest struct {
	Order_id string `json:"order_id" validate:"required"`
}

type ProductIDRequest struct {
	Product_id string `json:"product_id" validate:"required"`
}
type DesignIDRequest struct {
	Design_id string `json:"design_id" validate:"required"`
}
type FabricIDRequest struct {
	Fabric_id string `json:"fabric_id" validate:"required"`
}
type MaterialIDRequest struct {
	Material_id string `json:"material_id" validate:"required"`
}

type CreateOrderRequest struct {
	Store_phone   string `json:"store_phone" validate:"required"`
	Store_address string `json:"store_address" validate:"required"`
	User_phone    string `json:"user_phone" validate:"required"`
	User_address  string `json:"user_address" validate:"required"`
	Create_by     string `json:"create_by" validate:"required"`
}

type CreateProductRequest struct {
	Detail       string `json:"detail" db:"detail"`
	Size         string `json:"size" db:"size"`
	Max_quantity int    `json:"max_quantity " db:"max_quantity "`
	Create_by    string `json:"create_by" db:"create_by"`
}

type CreateDesignRequest struct {
	Design_url string `json:"design_url" db:"design_url"`
	Type       string `json:"type" db:"type"`
}

type CreateFabricRequest struct {
	Fabric_url string `json:"fabric_url" db:"fabric_url"`
	Quantity   int    `json:"quantity" db:"quantity"`
}

type CreateMaterialRequest struct {
	Material_name string `json:"material_name" db:"material_name"`
	Product_id    string `json:"product_id" db:"product_id"`
	Category      string `json:"category" db:"category"`
	Quantity      int    `json:"quantity" db:"quantity"`
}
