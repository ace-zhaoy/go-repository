package contract

type ENTITY[ID comparable] interface {
	GetID() ID
	SetID(ID)
}

type Order struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}
