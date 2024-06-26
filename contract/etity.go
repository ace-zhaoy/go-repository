package contract

type ENTITY[ID comparable] interface {
	GetID() ID
	SetID(ID)
}

type Order struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func (o *Order) ToString() string {
	s := " ASC"
	if o.Value < 0 {
		s = " DESC"
	}

	return o.Key + s
}
