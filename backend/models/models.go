package models



import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

type Morse struct {

	MorseCode string `json:"morsecode"`
}

type FizzBuzz struct {
	FizzBuzz string `json:"fizzBuzz,omitempty"`

}

type GuestLedger struct {
	Email string `json:"email"`
	Message string `json:"message"`

}

type Counter struct {
	Counter int `json:"counter,omitempty"`

}

type Quotes struct {
	Text string `json:"text"`
	Author string `json:"author"`

}
type Card struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title"`
	SubTitle string `json:"subTitle"`
	Text string `json:"text"`
	Author string `json:"author"`

}