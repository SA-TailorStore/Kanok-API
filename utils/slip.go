package utils

import (
	"strings"
)

var (
	BBL   = "002" // ธนาคารกรุงเทพ
	KBANK = "004" // ธนาคารกสิกรไทย
	KTB   = "006" // ธนาคารกรุงไทย
	TTB   = "011" // ธนาคารทหารไทยธนชาต
	SCB   = "014" // ธนาคารไทยพาณิชย์
	BAY   = "025" // ธนาคารกรุงศรีอยุธยา
	KKP   = "069" // ธนาคารเกียรตินาคินภัทร
	CIMBT = "022" // ธนาคารซีไอเอ็มบีไทย
	TISCO = "067" // ธนาคารทิสโก้
	UOBT  = "024" // ธนาคารยูโอบี
	TCD   = "071" // ธนาคารไทยเครดิตเพื่อรายย่อย
	LHFG  = "073" // ธนาคารแลนด์ แอนด์ เฮ้าส์
	ICBCT = "070" // ธนาคารไอซีบีซี (ไทย)
	SME   = "098" // ธนาคารพัฒนาวิสาหกิจขนาดกลางและขนาดย่อมแห่งประเทศไทย
	BAAC  = "034" // ธนาคารเพื่อการเกษตรและสหกรณ์การเกษตร
	EXIM  = "035" // ธนาคารเพื่อการส่งออกและนำเข้าแห่งประเทศไทย
	GSB   = "030" // ธนาคารออมสิน
	GHB   = "033" // ธนาคารอาคารสงเคราะห์
)

type QRCode struct {
	SendingBankCode   string
	ReceivingBankCode string
	IDK               string
	AccountNumber     string
	TransactionREF    string
	Type              string
	Language          string
	AdditionalInfo    string
}

func ParseCode(code string) QRCode {
	tranref := GetTransactionREF(code, code[18:21])
	lang := GetLanguage(code, tranref)
	addition := GetAdditionalInfo(code, lang)
	return QRCode{
		ReceivingBankCode: code[:3],
		IDK:               code[3:18],
		SendingBankCode:   code[18:21],
		AccountNumber:     code[21:25],
		TransactionREF:    tranref,
		Type:              lang[0:4],
		Language:          lang[4:],
		AdditionalInfo:    addition,
	}
}

func GetTransactionREF(str string, bankcode string) string {
	switch bankcode {
	case KBANK:
		return str[25:45]
	case SCB:
		return str[25:50]
	default:
		return str
	}
}

func GetLanguage(code string, transref string) string {
	index := strings.Index(code, transref)
	index = index + len(transref)
	return code[index : index+6]
}

func GetAdditionalInfo(code string, lang string) string {
	index := strings.Index(code, lang)
	index = index + len(lang)
	return code[index:]
}

func CalculateCRC16(data string) uint16 {
	var crc uint16 = 0xFFFF // ค่าเริ่มต้น

	for i := 0; i < len(data); i++ {
		crc ^= uint16(data[i]) << 8
		for j := 0; j < 8; j++ {
			if (crc & 0x8000) != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}

	return crc & 0xFFFF
}
