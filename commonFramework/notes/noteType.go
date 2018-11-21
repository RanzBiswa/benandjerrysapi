package notes

import (
	"strings"
)

const (
	userNote      string = "UserNotes"
	orderNote            = "OrderNotes"
	recipientNote        = "RecipientNotes"
	all                  = "All"
)

//NoteType model a note type
type NoteType struct {
	Value string
}

const (
	//UserNoteParam query param value for user notes
	UserNoteParam string = "UN"
	//OrderNoteParam query param value for order notes
	OrderNoteParam = "ON"
	//RecipientNoteParam query param value for recipient notes
	RecipientNoteParam = "RN"
	//AllNotesParam query param value for all notes
	AllNotesParam = "X"
)

var globalNoteTypeText = map[string]NoteType{
	UserNoteParam:      NoteType{Value: userNote},
	OrderNoteParam:     NoteType{Value: orderNote},
	RecipientNoteParam: NoteType{Value: recipientNote},
	AllNotesParam:      NoteType{Value: all},
}

//GlobalNoteTypeText returns the note type text passsed to service
func GlobalNoteType(ntype string) *NoteType {
	if t, exists := globalNoteTypeText[strings.ToUpper(ntype)]; !exists {
		return nil
	} else {
		return &t
	}
}

var orderNoteTypeText = map[string]NoteType{
	UserNoteParam:  NoteType{Value: userNote},
	OrderNoteParam: NoteType{Value: orderNote},
}

//OrderNoteTypeText returns the order note type text passsed to service
func OrderNoteType(ntype string) *NoteType {
	if t, exists := orderNoteTypeText[strings.ToUpper(ntype)]; !exists {
		return nil
	} else {
		return &t
	}
}

var orderNoteRecordsTypeText = map[string]string{
	"U": "User",
	"S": "System",
	"F": "Financial",
	"N": "None",
}

//OrderNoteRecordsType returns the order note type required to post order notes
func OrderNoteRecordsType(nType string) string {
	return orderNoteRecordsTypeText[strings.ToUpper(nType)]
}

var noteRecordsSubTypeText = map[string]string{
	"U":  "User",
	"M":  "Manifest",
	"PP": "PrePrinted",
	"R":  "Receipt",
	"BT": "BillTo",
	"ST": "ShipTo",
	"D":  "Detail",
	"P":  "Payment",
	"O":  "Order",
	"S":  "System",
	"F":  "Financial",
	"B":  "Billing",
	"Q":  "Queue",
	"MB": "MessageBroker",
	"IA": "InventoryAssignment",
	"CS": "CustomSku",
	"N":  "None",
}

//NoteRecordsSubType returns the order note type required to post order notes
func NoteRecordsSubType(nSubtype string) string {
	return noteRecordsSubTypeText[strings.ToUpper(nSubtype)]
}
