package requests

// USER REQUEST
type UserRegister struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	Phone_number string `json:"phone_number" validate:"required"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Username struct {
	Username string `json:"username" validate:"required"`
}

type UserJWT struct {
	Token string `json:"token" validate:"required"`
}

type UserID struct {
	User_id string `json:"user_id" validate:"required"`
}

type UserUpdate struct {
	Token        string `json:"token" validate:"required"`
	Display_name string `json:"display_name" validate:"required"`
	Phone_number string `json:"phone_number" validate:"required"`
	Address      string `json:"address" validate:"required"`
}

// ORDER REQUEST
type OrderID struct {
	Order_id string `json:"order_id" validate:"required"`
}

type CreateOrder struct {
	Store_phone   string `json:"store_phone" validate:"required"`
	Store_address string `json:"store_address" validate:"required"`
	User_phone    string `json:"user_phone" validate:"required"`
	User_address  string `json:"user_address" validate:"required"`
	Create_by     string `json:"create_by" validate:"required"`
}

// PRODUCT REQUEST
type ProductID struct {
	Product_id string `json:"product_id" validate:"required"`
}

type CreateProduct struct {
	Design_id    string `json:"design_id" db:"design_id"`
	Fabric_id    string `json:"fabric_id" db:"fabric_id"`
	Detail       string `json:"detail" db:"detail"`
	Size         string `json:"size" db:"size"`
	Max_quantity int    `json:"max_quantity " db:"max_quantity "`
	Create_by    string `json:"create_by" db:"create_by"`
}

// DESIGN REQUEST
type DesignID struct {
	Design_id string `json:"design_id" validate:"required"`
}

type CreateDesign struct {
	Design_url string `json:"design_url" db:"design_url"`
	Type       string `json:"type" db:"type"`
}

// FABRIC REQUEST
type FabricID struct {
	Fabric_id string `json:"fabric_id" validate:"required"`
}

type CreateFabric struct {
	Fabric_url string `json:"fabric_url" db:"fabric_url"`
	Quantity   int    `json:"quantity" db:"quantity"`
}

// METERIAL REQUEST
type MaterialID struct {
	Material_id string `json:"material_id" validate:"required"`
}

type CreateMaterial struct {
	Material_name string `json:"material_name" db:"material_name"`
	Product_id    string `json:"product_id" db:"product_id"`
	Category      string `json:"category" db:"category"`
	Quantity      int    `json:"quantity" db:"quantity"`
}
