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

type UserRole struct {
	Role string `json:"role"`
}

type UserUpdate struct {
	Token        string `json:"token" validate:"required"`
	Display_name string `json:"display_name"`
	Phone_number string `json:"phone_number"`
	Address      string `json:"address"`
}

type UserUploadImage struct {
	Token string `json:"jwt" form:"jwt" validate:"required"`
	Image string `json:"image" form:"image" `
}

// ORDER REQUEST
type OrderID struct {
	Order_id string `json:"order_id" validate:"required"`
}

type UpdateOrder struct {
	Order_id      string `json:"order_id" validate:"required"`
	Store_phone   string `json:"store_phone" validate:"required"`
	Store_address string `json:"store_address" validate:"required"`
	User_phone    string `json:"user_phone" validate:"required"`
	User_address  string `json:"user_address" validate:"required"`
}
type CreateOrder struct {
	Token    string    `json:"token" validate:"required"`
	Products []Product `json:"products" validate:"required"`
}

type UpdateStatus struct {
	Order_id string `json:"order_id" validate:"required"`
	Status   string `json:"status" validate:"required"`
	Price    int    `json:"price"`
}

type UpdatePayment struct {
	Order_id   string `json:"order_id" validate:"required"`
	Image      string `json:"image"`
	Is_payment int    `json:"is_payment"`
}

type UpdateTracking struct {
	Order_id        string `json:"order_id" validate:"required"`
	Tracking_number string `json:"tracking_number" validate:"required"`
}

type UpdatePrice struct {
	Order_id string  `json:"order_id" validate:"required"`
	Price    float32 `json:"price" validate:"required"`
}

type UpdateTailor struct {
	Order_id  string `json:"order_id" validate:"required"`
	Tailor_id string `json:"tailor_id" validate:"required"`
	Due_date  string `json:"due_date" validate:"required"`
}

// PRODUCT REQUEST
type ProductID struct {
	Product_id string `json:"product_id" validate:"required"`
}

type ProductOrderID struct {
	Created_by string `json:"create_by" validate:"required"`
}

type Product struct {
	Design_id      int    `json:"design_id" validate:"required"`
	Fabric_id      int    `json:"fabric_id" validate:"required"`
	Detail         string `json:"detail" validate:"required"`
	Size           string `json:"size" validate:"required"`
	Total_quantity int    `json:"total_quantity" validate:"required"`
}

type CreateProduct struct {
	Order_id string    `json:"order_id" validate:"required"`
	Products []Product `json:"products" validate:"required"`
}

type UpdateProduct struct {
	Product_id     string `json:"product_id" validate:"required"`
	Design_id      string `json:"design_id" validate:"required"`
	Fabric_id      string `json:"fabric_id" validate:"required"`
	Detail         string `json:"detail" validate:"required"`
	Size           string `json:"size" validate:"required"`
	Total_quantity int    `json:"total_quantity" validate:"required"`
}

type UpdateProcessQuantity struct {
	Product_id        string `json:"product_id" validate:"required"`
	Increase_quantity int    `json:"increase_quantity"`
	Decrease_quantity int    `json:"decrease_quantity"`
}

// DESIGN REQUEST
type DesignID struct {
	Design_id int `json:"design_id" validate:"required"`
}

type AddDesign struct {
	Image string `json:"image"`
	Type  string `json:"type" validate:"required"`
}

type UpdateDesign struct {
	Design_id int    `json:"design_id" validate:"required"`
	Image     string `json:"image"`
	Type      string `json:"type"`
}

type DeleteDesign struct {
	Design_ID int `json:"design_id" validate:"required"`
}

// FABRIC REQUEST
type FabricID struct {
	Fabric_id int `json:"fabric_id" validate:"required"`
}

type AddFabric struct {
	Image    string `json:"image" `
	Quantity int    `json:"quantity" validate:"required"`
}

type UpdateFabric struct {
	Fabric_id int    `json:"fabric_id" validate:"required"`
	Image     string `json:"image"`
	Quantity  int    `json:"quantity"`
}
type UpdateFabrics struct {
	Fabric_id int `json:"fabric_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

// METERIAL REQUEST
type MaterialID struct {
	Material_id int `json:"material_id" validate:"required"`
}

type AddMaterial struct {
	Material_name string `json:"material_name" validate:"required"`
	Amount        int    `json:"amount"`
}

type UpdateMaterial struct {
	Material_id   int    `json:"material_id" validate:"required"`
	Material_name string `json:"material_name"`
	Amount        int    `json:"amount"`
}
