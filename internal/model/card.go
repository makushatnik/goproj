package model

type CardType int

type Card struct {
	ID     int
	Number string
	Type   CardType
	UserID int
}

const (
	Undefined CardType = iota
	Visa
	Mastercard
	UnionPay
	Mir
	JCB
)

func (ct CardType) String() string {
	switch ct {
	case Visa:
		return "visa"
	case Mastercard:
		return "mastercard"
	case UnionPay:
		return "unionpay"
	case Mir:
		return "mir"
	case JCB:
		return "jcb"
	}
	return "undefined"
}
