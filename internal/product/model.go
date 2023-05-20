package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `gorm:"type:varchar(60);" json:"name"`
	Sku         string   `gorm:"not null" json:"sku"`
	Price       int      `json:"price"`
	Color       string   `gorm:"type:varchar(15);" json:"color"`
	Measures    Measures `gorm:"embedded" json:"measures"`
	Size        string   `gorm:"type:varchar(5);" json:"size"`
	Material    string   `gorm:"type:varchar(20)" json:"material"`
	Category    string   `gorm:"type:varchar(25);" json:"category"`
	Quantity    int      `json:"quantity"`
	Description string   `json:"description"`
	Images      Images   `gorm:"embedded" json:"images"`
}

type Measures struct {
	High int `json:"high"`
	Wide int `json:"wide"`
	Long int `json:"long"`
}

type Images struct {
	ImageMain string `json:"image_main"`
	Image2    string `json:"image_2"`
	Image3    string `json:"image_3"`
	Image4    string `json:"image_4"`
	Image5    string `json:"image_5"`
	Image6    string `json:"image_6"`
	Image7    string `json:"image_7"`
}
