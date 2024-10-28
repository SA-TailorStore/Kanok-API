package responses

// USER
type UserLogin struct {
	User_id  string `json:"user_id"`
	Password string `json:"password"`
}

type User struct {
	User_id          string `json:"user_id"`
	Username         string `json:"username"`
	Display_name     string `json:"display_name"`
	User_profile_url string `json:"user_profile_url"`
	Role             string `json:"role"`
	Phone_number     string `json:"phone_number"`
	Address          string `json:"address"`
	Timestamp        string `json:"timestamp"`
}

type UserTailor struct {
	User_id          string `json:"user_id"`
	Username         string `json:"username"`
	Display_name     string `json:"display_name"`
	User_profile_url string `json:"user_profile_url"`
	Role             string `json:"role"`
	Phone_number     string `json:"phone_number"`
	Address          string `json:"address"`
	Timestamp        string `json:"timestamp"`
	Order_id         string `json:"order_id"`
	Product_Process  int    `json:"product_process"`
	Product_Total    int    `json:"product_total"`
}

type Username struct {
	Username string `json:"username"`
}

type UserID struct {
	User_id string `json:"user_id"`
}

type UserDisplayName struct {
	Display_name string `json:"display_name"`
}

type UserProUrl struct {
	User_id          string `json:"user_id"`
	User_profile_url string `json:"user_profile_url"`
}

type UserRole struct {
	Role string `json:"role"`
}

type UserJWT struct {
	Token string `json:"token"`
}

type UserCreateOrder struct {
	Display_name string `json:"display_name"`
	Phone_number string `json:"phone_number"`
	Address      string `json:"address"`
}

// ORDER
type Order struct {
	Order_id        string `json:"order_id"`
	Is_payment      int    `json:"is_payment"`
	Status          string `json:"status"`
	Store_phone     string `json:"store_phone"`
	Store_address   string `json:"store_address"`
	Price           int    `json:"price"`
	Due_date        string `json:"due_date"`
	Tracking_number string `json:"tracking_number"`
	Created_by      string `json:"created_by"`
	User_phone      string `json:"user_phone"`
	User_address    string `json:"user_address"`
	Tailor_id       string `json:"tailor_id"`
	Tailor_phone    string `json:"tailor_phone"`
	Tailor_address  string `json:"tailor_address"`
	Timestamp       string `json:"timestamp"`
}

type ShowOrder struct {
	Order_id  string `json:"order_id"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type OrderID struct {
	Order_id string `json:"order_id"`
}

type CheckProcess struct {
	Is_ready bool `json:"is_ready"`
}

type CreateOrder struct {
	Order_id string          `json:"order_id"`
	Products map[string]bool `json:"product"`
}

// PRODUCT
type Product struct {
	Product_id       string `json:"product_id"`
	Design_id        int    `json:"design_id"`
	Fabric_id        int    `json:"fabric_id"`
	Detail           string `json:"detail"`
	Size             string `json:"size"`
	Process_quantity int    `json:"process_quantity"`
	Total_quantity   int    `json:"total_quantity"`
	Created_by       string `json:"created_by"`
	Timestamp        string `json:"timestamp"`
	Design_url       string `json:"design_url"`
}

type ProductID struct {
	Product_id string `json:"product_id"`
}

type ProductProcess struct {
	Is_ready         bool `json:"is_ready"`
	Process_quantity int  `json:"process_quantity"`
	Total_quantity   int  `json:"total_quantity"`
}

// DESIGN
type Design struct {
	Design_id  int    `json:"design_id" db:"design_id"`
	Design_url string `json:"design_url" db:"design_url"`
	Type       string `json:"type" db:"type"`
}

type DesignID struct {
	Design_id int `json:"design_id" db:"design_id"`
}

type DesignURL struct {
	Design_url string `json:"design_url" db:"design_url"`
}

type DesignType struct {
	Type string `json:"type" db:"type"`
}

// FABRIC
type Fabric struct {
	Fabric_id  int    `json:"fabric_id" db:"fabric_id"`
	Fabric_url string `json:"fabric_url" db:"fabric_url"`
	Quantity   int    `json:"quantity" db:"quantity"`
}

type FabricID struct {
	Fabric_id int `json:"fabric_id" db:"fabric_id"`
}

type FabricURL struct {
	Fabric_url string `json:"fabric_url" db:"fabric_url"`
}

type FabricQuantity struct {
	Quantity int `json:"quantity" db:"quantity"`
}

type CheckFabric struct {
	Product_index string `json:"product_index" db:"fabric_id"`
	Quantity      bool   `json:"quantity" db:"quantity"`
}

// MATERIAL
type Material struct {
	Material_id   int    `json:"material_id"`
	Material_name string `json:"material_name"`
	Amount        int    `json:"amount"`
}

type MaterialName struct {
	Material_name string `json:"material_name" db:"material_name"`
}

type MaterialAmount struct {
	Amount int `json:"amount" db:"amount"`
}

type MaterialID struct {
	Material_id int `json:"material_id"`
}
