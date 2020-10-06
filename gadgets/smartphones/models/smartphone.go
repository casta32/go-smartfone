package models

// Smartphone model structure for smartphone
type Smartphone struct {
	Id            int64 // similar to long java -ID autoI
	Name          string
	Price         int
	CounterOrigin string
	Os            string
}

type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CounterOrigin string `json:"contry_origin"`
	Os            string `json:"os"`
}
