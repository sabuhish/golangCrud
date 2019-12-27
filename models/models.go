package models


type Car struct {
	ID  string     `json: "id,omitempty"`
	Model  string     `json: "model,omitempty"`
	Price  string     `json: "price,omitempty"`
	Interior *Interioe `json:"sad"`
}


type Interioe struct {
	Material string `json:"materaila"`
	Display string `json:"display"`

}