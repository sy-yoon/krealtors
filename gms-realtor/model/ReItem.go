package model

type ReItem struct {
	Id            int32       `json:"id" bson:"_id"`
	Title         string      `json:"title" bson:"title"`
	UserId        string      `json:"userId" bson:"userId"`
	ReType        int32       `json:"type" bson:"type"`
	TxType        int32       `json:"txType" bson:"txType"`
	Currency      string      `json:"currency" bson:"currency"`
	CityId        string      `json:"cityId" bson:"cityId"`
	Thumbnail     string      `json:"thumbnail" bson:"thumbnail"`
	Images        []string    `json:"images" bson:"images"`
	Price         uint64      `json:"price" bson:"price"`
	Heating       string      `json:"heating" bson:"heating"`
	Facing        string      `json:"facing" bson:"facing"`
	Content       string      `json:"content" bson:"content"`
	Address       string      `json:"address" bson:"address"`
	GeoLocation   string      `json:"location" bson:"location"`
	Bedroom       int32       `json:"bedroom" bson:"bedroom"`
	Bathroom      int32       `json:"bathroom" bson:"bathroom"`
	Parking       int32       `json:"parking" bson:"parking"`
	Area          string      `json:"area" bson:"area"`
	Furnished     bool        `json:"furnished" bson:"furnished"`
	Garage        bool        `json:"garage" bson:"garage"`
	AvailableDate string      `json:"availableDate" bson:"availableDate"`
	Options       OptionsType `json:"options" bson:"options"`
	CreatedDate   uint64      `json:"createdDt" bson:"createdDt"`
	UpdatedDate   uint64      `json:"updatedDt" bson:"updatedDt"`
}

type LocationType struct {
	Latitude  float64 `json:"lat" bson:"lat"`
	Longitude float64 `json:"lng" bson:"lng"`
}

type OptionsType struct {
	Airconditioner bool `json:"aircon" bson:"aircon"`
	Refrigerator   bool `json:"refrigerator" bson:"refrigerator"`
	WashingMachine bool `json:"washingmachine" bson:"washingmachine"`
	Dishwasher     bool `json:"dishwasher" bson:"dishwasher"`
	Microwave      bool `json:"microwave" bson:"microwave"`
	TV             bool `json:"tv" bson:"tv"`
	Doorlock       bool `json:"doorlock" bson:"doorlock"`
}
