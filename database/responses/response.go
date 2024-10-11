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
	User_profile_url string `json:"user_profile_url"`
}

type UserRole struct {
	Role string `json:"role"`
}

type UserJWT struct {
	Token string `json:"token"`
}

type UserAddressPhone struct {
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
	User_phone      string `json:"user_phone"`
	User_address    string `json:"user_address"`
	Price           int    `json:"price"`
	Due_date        string `json:"due_date"`
	Tracking_number string `json:"tracking_number"`
	Tailor_id       string `json:"tailor_id"`
	Created_by      string `json:"created_by"`
	Timestamp       string `json:"timestamp"`
}

type OrderID struct {
	Order_id string `json:"order_id"`
}

// PRODUCT
type Product struct {
	Product_id       string `json:"product_id"`
	Design_id        string `json:"design_id"`
	Fabric_id        string `json:"fabric_id"`
	Detail           string `json:"detail"`
	Size             string `json:"size"`
	Process_quantity int    `json:"process_quantity"`
	Total_quantity   int    `json:"total_quantity"`
	Created_by       string `json:"created_by"`
	Timestamp        string `json:"timestamp"`
}

type ProductID struct {
	Product_id string `json:"product_id"`
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

type FabricType struct {
	Quantity int `json:"quantity" db:"quantity"`
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
