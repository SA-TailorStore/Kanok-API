package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/go-resty/resty/v2"
	"github.com/liyue201/goqr"
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
	Checksum          string
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
		AdditionalInfo:    addition[0:4],
		Checksum:          addition[4:],
	}
}

func GetTransactionREF(str string, bankcode string) string {
	switch bankcode {
	case KBANK:
		return str[25:45]
	case SCB:
		return str[25:50]
	default:
		return str[25:25]
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

func HexToDecimal(hexString string) (int32, error) {
	decimalValue, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		return 0, err
	}
	return int32(decimalValue), nil
}

func GetStringQR(qrcode []*goqr.QRData) string {
	var result string
	for _, qrData := range qrcode {
		result += string(qrData.Payload)
	}
	return result
}

func SendString(str string) (map[string]interface{}, error) {
	var result map[string]interface{}
	body := map[string]interface{}{
		"data": str,
		"log":  true,
	}

	req := resty.New()
	resp, err := req.R().
		SetHeader("x-authorization", configs.NewConfig().SlipOk_Secret). // ตั้ง Header อื่นๆ
		SetBody(body).
		Post(configs.NewConfig().SlipOk_url)

	if err != nil {
		return result, err
	}

	json.Unmarshal([]byte(resp.String()), &result)
	/*
		fmt.Println("Url:", configs.NewConfig().SlipOk_url)
		fmt.Println("Response Status:", resp.Status())
		fmt.Println("Response Code:", resp.StatusCode())
	*/
	fmt.Println("Body:", resp.String())
	return result, err
}

func Bypass(str string) (map[string]interface{}, error) {
	result := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"success":       true,
			"message":       "✅",
			"rqUID":         "783_20191108_v4UIS1K2Mobile",
			"language":      "TH",
			"transRef":      "010092101507665143",
			"sendingBank":   "004",
			"receivingBank": "004",
			"transDate":     "20200401",
			"transTime":     "10:15:07",
			"sender": map[string]interface{}{
				"displayName": "นาย ธนาคาร ก",
				"name":        "Mr. Thanakarn K",
				"proxy": map[string]interface{}{
					"type":  nil,
					"value": nil,
				},
				"account": map[string]interface{}{
					"type":  "BANKAC",
					"value": "xxx-x-x0209-x",
				},
			},
			"receiver": map[string]interface{}{
				"displayName": "กสิกร ร",
				"name":        "KASIKORN R",
				"proxy": map[string]interface{}{
					"type":  "",
					"value": "",
				},
				"account": map[string]interface{}{
					"type":  "BANKAC",
					"value": "xxx-x-x3109-x",
				},
			},
			"amount":            1,
			"paidLocalAmount":   1,
			"paidLocalCurrency": "764",
			"countryCode":       "TH",
			"transFeeAmount":    0,
			"ref1":              "",
			"ref2":              "",
			"ref3":              "",
			"toMerchantId":      "",
		},
	}
	return result, nil
}
