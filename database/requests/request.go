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

type UserUploadImage struct {
	Token string `json:"jwt" form:"jwt" validate:"required"`
	Image string `json:"image" form:"image" `
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
	Design_id      string `json:"design_id" `
	Fabric_id      string `json:"fabric_id" `
	Detail         string `json:"detail" `
	Size           string `json:"size" validate:"required"`
	Total_quantity int    `json:"total_quantity " validate:"required"`
	Create_by      string `json:"create_by" validate:"required"`
}

// DESIGN REQUEST
type DesignID struct {
	Design_id string `json:"design_id" validate:"required"`
}

type AddDesign struct {
	Image string `json:"image"`
	Type  string `json:"type" validate:"required"`
}

type UpdateDesign struct {
	Design_ID string `json:"design_id" db:"design_id"`
	Image     string `json:"image"`
	Type      string `json:"type" db:"type"`
}

type DeleteDesign struct {
	Design_ID string `json:"design_id" db:"design_id"`
}

// FABRIC REQUEST
type FabricID struct {
	Fabric_id string `json:"fabric_id" validate:"required"`
}

type AddFabric struct {
	Fabric_url string `json:"fabric_url" `
	Quantity   int    `json:"quantity" validate:"required"`
}

type UpdateFabric struct {
	Fabric_id  string `json:"fabric_id" validate:"required"`
	Fabric_url string `json:"fabric_url" `
	Quantity   int    `json:"quantity" validate:"required"`
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
