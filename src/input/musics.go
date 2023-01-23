package input

type InputMusic struct {
	Music_ID    uint   `gorm:"primaryKey" json:"Music_id"`
	Name        string `json:"name" valid:"type(string)"`
	Location    string `json:"location" valid:"type(string)"`
	Description string `json:"description" valid:"type(string)"`
	Price       string `json:"price" valid:"type(string)"`
	Status      string `json:"status" valid:"type(string)"`
	Stock       int    `json:"stock" valid:"type(int)"`
	Category    string `json:"category" valid:"type(string)"`
	Image       string `json:"image"`
	Rating      int    `json:"rating" valid:"type(int)"`
}