package request

import "main/server/model"



type NoticeRequest struct {

	
	Title string `bson:"title"`
	Text string `bson:"text"`
	SendTo []model.TeamID `json:"sendTo"`
	From_Date string`bson:"fromDate"`
	To_Date string  `bson:"toDate"`

}