package exceptions

import "errors"

var (
	ErrMaterialNotFound   = errors.New("ไม่พบวัสดุ")
	ErrBadRequestMaterial = errors.New("bad request data type wrong")
	ErrDupicatedName      = errors.New("ชื่อวัสดุซ้ำกรุณาเพิ่มวัสดุใหม่อีกครั้ง")
)
