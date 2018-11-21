package locale

import "strings"

const (
	//CBUSEnglish locale for CB US English
	CBUSEnglish string = "cb-en-us"
	//CB2USEnglish locale for CB2 US English
	CB2USEnglish string = "c2-en-us"
	//LONUSEnglish locale for LON US English
	LONUSEnglish string = "nd-en-us"
	//CBCANEnglish locale for CB CAN English
	CBCANEnglish string = "cn-en-ca"
	//CB2CANEnglish locale for CB2 CAN English
	CB2CANEnglish string = "c2-en-ca"
	//CBCANFrench locale for CB CAN French
	CBCANFrench string = "cn-fr-ca"
	//CB2CANFrench locale for CB2 CAN French
	CB2CANFrench string = "c2-fr-ca"
)

//Get gets locale based on Brand and language ID (used in backend service) value
func Get(brand string, languageID string) string {
	var locale string

	if len(brand) == 0 {
		return "cb-en-us"
	}

	switch strings.ToUpper(brand) {
	case "CB":
		locale = "cb-en-us"
	case "CB2":
		locale = "c2-en-us"
	case "LON":
		locale = "nd-en-us"
	case "CANCB":
		locale = "cn-en-ca"
		if strings.ToUpper(languageID) == "FR_CA" {
			locale = "cn-fr-ca"
		}
	case "CANCB2":
		locale = "c2-en-ca"
		if strings.ToUpper(languageID) == "FR_CA" {
			locale = "c2-fr-ca"
		}
	}
	return locale
}

//GetCompany gets company based on locale
func GetCompany(locale string) string {
	var company string

	if len(locale) == 0 {
		return "CB"
	}

	switch strings.ToLower(locale) {
	case "cb-en-us":
		company = "CB"
	case "c2-en-us":
		company = "CB2"
	case "nd-en-us":
		company = "LON"
	case "cn-en-ca":
		company = "CANCB"
	case "c2-en-ca":
		company = "CANCB2"
	case "cn-fr-ca":
		company = "CANCB"
	case "c2-fr-ca":
		company = "CANCB2"
	}
	return company
}

//GetCompanyCode gets company code based on company name
func GetCompanyCode(company string) int {
	var companyCode int

	if len(company) == 0 {
		return 100
	}

	switch strings.ToUpper(company) {
	case "CB":
		companyCode = 100
	case "CB2":
		companyCode = 200
	case "LON":
		companyCode = 300
	case "CANCB":
		companyCode = 301
	case "CANCB2":
		companyCode = 201
	}

	return companyCode
}

//GetCompanyCodeOfLocale gets companyCode and language (used in back end) based on locale
func GetCompanyCodeOfLocale(locale string) (int, string) {
	var companyCode int
	var lang string

	if len(locale) == 0 {
		return 100, "en_US"
	}

	switch strings.ToLower(locale) {
	case "cb-en-us":
		companyCode = 100
		lang = "en_US"
	case "c2-en-us":
		companyCode = 200
		lang = "en_US"
	case "nd-en-us":
		companyCode = 300
		lang = "en_US"
	case "cn-en-ca":
		companyCode = 301
		lang = "en_US"
	case "c2-en-ca":
		companyCode = 201
		lang = "en_US"
	case "cn-fr-ca":
		companyCode = 301
		lang = "fr_CA"
	case "c2-fr-ca":
		companyCode = 201
		lang = "fr_CA"
	}
	return companyCode, lang
}

//IsMultiLang checks if the company-based locale is a multi lang locale and returns the base/parent locale. e.g. cn-fr-ca(multi lang), cn-en-ca(base)
func IsMultiLang(locale string) (multiLang bool, baseLocale string) {
	if strings.ToLower(locale) == "cn-fr-ca" {
		multiLang = true
		baseLocale = "cn-en-ca"
	}
	return
}

var webSvcLocale = map[string]string{
	"EN_US": "en_US",
	"FR_CA": "fr_CA",
}

//WebSvcLocale returns locale (language ID) understood by web service
func WebSvcLocale(locale string) string {
	return webSvcLocale[strings.ToUpper(locale)]
}

var locales = map[string]string{
	"cb-en-us": "Crate US English",
	"c2-en-us": "CB2 US English",
	"nd-en-us": "Land of Nod US English",
	"cn-en-ca": "Crate Canada English",
	"c2-en-ca": "CB2 Canada English",
	"cn-fr-ca": "Crate Canada French",
}

//ValidLocale validates the locale string
func ValidLocale(locale string) bool {
	_, exists := locales[strings.ToLower(locale)]
	return exists
}

var companyCodeLocales = map[int]string{
	100: CBUSEnglish,
	200: CB2USEnglish,
	300: LONUSEnglish,
	301: CBCANEnglish,
	201: CB2CANEnglish,
}

//TranslateCompCodeToLocale Returns corresponding company locale
func TranslateCompCodeToLocale(code int) string {
	return companyCodeLocales[code]
}
