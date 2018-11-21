package crateTransactionID

import (
	"fmt"
	"strconv"
	"time"

	_l "github.com/zalora_icecream/commonFramework/locale"
)

var rollingNumber = 0

//GetTransactionID Used to generate a Crate and Barrel transaction ID
func GetTransactionID() string {
	tranID := ""
	tranID = tranID + "1"
	tranID = tranID + "9"
	tranID = tranID + "01"
	rollingNumber = rollingNumber + 1
	if rollingNumber > 9999 {
		//rolling number has to 4 digits
		rollingNumber = 1
	}
	tranID = tranID + pad(rollingNumber, "4")
	tranID = tranID + pad(numOfDays(), "5")
	tranID = tranID + pad(numOfSeconds(), "5")
	cd, _ := checkDigit(tranID)
	tranID = tranID + strconv.Itoa(cd)

	return tranID

}

//GetTransactionIDV2 Used to generate a Crate and Barrel transaction ID
//Takes care of adding first digit based on company
func GetTransactionIDV2(locale string) string {
	tranID := ""
	tranID = tranID + firstDigit(locale)
	tranID = tranID + "9"
	tranID = tranID + "01"
	rollingNumber = rollingNumber + 1
	if rollingNumber > 9999 {
		//rolling number has to 4 digits
		rollingNumber = 1
	}
	tranID = tranID + pad(rollingNumber, "4")
	tranID = tranID + pad(numOfDays(), "5")
	tranID = tranID + pad(numOfSeconds(), "5")
	cd, _ := checkDigit(tranID)
	tranID = tranID + strconv.Itoa(cd)

	return tranID

}

func pad(number int, numberOfDigits string) string {
	numberString := fmt.Sprintf("%0"+numberOfDigits+"v", number)
	return numberString
}

func numOfDays() int {
	t2 := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)
	t1 := time.Now()
	delta := t1.Sub(t2)
	numOfDays := int(delta.Hours() / 24)
	return numOfDays
}

func numOfSeconds() int {
	now := time.Now()
	then := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	delta := now.Sub(then)
	numOfSeconds := int(delta.Seconds())
	return numOfSeconds

}

// Calculates the luhn on the number string
func checkDigit(numbers string) (int, error) {
	if _, err := strconv.ParseInt(numbers, 10, 64); err != nil {
		return 0, err
	}

	checksum := 0
	result := 0

	// Calculate the checksum, from right to left.
	for i := len(numbers) - 1; i >= 0; i-- {
		result = int(numbers[i] - '0')
		isOdd := (len(numbers)-i+1)%2 == 0

		if isOdd {
			result *= 2

			if 10 <= result {
				result -= 9
			}
		}

		checksum += result
	}

	checksum = (checksum % 10)

	var value int

	if checksum == 0 {
		value = 0
	} else {
		value = 10 - checksum
	}

	return value, nil
}

//returns first digit to be used to form a Transaction ID based on the company
func firstDigit(locale string) string {

	switch locale {
	case _l.CBUSEnglish:
		return "1"
	case _l.CB2USEnglish:
		return "2"
	case _l.LONUSEnglish:
		return "3"
	case _l.CBCANEnglish,
		_l.CBCANFrench:
		return "4"
	case _l.CB2CANEnglish,
		_l.CB2CANFrench:
		return "5"
	default:
		return "1"
	}

}
