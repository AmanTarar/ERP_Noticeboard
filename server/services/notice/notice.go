package notice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/services/token"
	"main/server/utils"
	"net/http"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func AddNoticeService(ctx *gin.Context,Input_Notice request.NoticeRequest){




	// tokn:=token.GenerateToken(clms,ctx)
	tokn:=ctx.Request.Header.Get("authorization")
	claims,_:=token.DecodeToken(tokn)

	var notice model.Notice

	notice.Creator_id=claims.Id
	fmt.Println("claimsid",claims.Id)

	//get request to database to get user info

	req,err:=http.NewRequest("GET","https://timedragon.staging.chicmic.co.in/v1/user?_id=63ce887193e913067d03127b",ctx.Request.Body)

	req.Header.Add("authorization",tokn)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
	fmt.Println("error",err)
	return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	// fmt.Println("resp body",string(body))

	var userDetails model.UserDetails
	err=json.Unmarshal(body,&userDetails)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	fmt.Println("user's team details is :",userDetails)

	notice.Text=Input_Notice.Text

	notice.Title=Input_Notice.Title

	notice.SendTo=Input_Notice.SendTo
	
	notice.From_Date,_=time.Parse("02 Jan 2006",Input_Notice.From_Date)

	notice.Created_at=time.Now()

	notice.Creator_name=userDetails.Data.Name


	notice.To_Date,_=time.Parse("02 Jan 2006",Input_Notice.To_Date)

	notice.SendTo=Input_Notice.SendTo
	
	notice.Notice_id=primitive.NewObjectID().Hex()

	fmt.Println("notice",notice)

	result,err:= db.Collection.InsertOne(context.TODO(), notice)
	fmt.Println("result id",result)



	if err!=nil{

		fmt.Println("error inserting error in db")
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}

	fmt.Println("result",result)
	response.ShowResponse("success",int64(utils.HTTP_OK),"successully added notice",result,ctx)


	

	
}

func GetNoticeByTeamId(ctx *gin.Context) {
	// set header.
	

	tokn:=ctx.Request.Header.Get("authorization")
	claims,_:=token.DecodeToken(tokn)



	var notice model.Notice
	// we get params with mux.
	paramID:= ctx.Query("id")
	
	fmt.Println("param id",paramID)
	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(paramID)
	fmt.Println("id",id)

	// We create filter.
	// filter := bson.M{"_id": id}

	filter := bson.M{"SendTo": bson.M{"$_id":claims.Id}}
	
	err:= db.Collection.FindOne(context.TODO(), filter).Decode(&notice)

	//IF THE teamID of the user who is requesting ,is present inside sendTo object

	



	if err != nil {
		fmt.Println("error in db toooo find one")
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}

	response.ShowResponse("success",int64(utils.HTTP_OK),"get notice success",notice,ctx)
}







func GetNotice(ctx *gin.Context) {
	
	

	// tokn:=ctx.Request.Header.Get("authToken")
	// claims,_:=token.DecodeToken(tokn)



	var notice model.Notice
	// we get params with mux.
	paramID:= ctx.Query("id")



	


	// We create filter.
	filter := bson.M{"_id": paramID}
	
	err:= db.Collection.FindOne(context.TODO(), filter).Decode(&notice)


	if err != nil {
		fmt.Println("error in db to find one")
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}

	response.ShowResponse("success",int64(utils.HTTP_OK),"get notice success",notice,ctx)
}


func GetNotices(ctx *gin.Context) {
	


	tokn:=ctx.Request.Header.Get("authorization")
	claims,_:=token.DecodeToken(tokn)
	fmt.Println("cliams",claims.Id)

	req,err:=http.NewRequest("GET","https://timedragon.staging.chicmic.co.in/v1/user?_id=63ce887193e913067d03127b",ctx.Request.Body)

	req.Header.Add("authorization",tokn)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
	fmt.Println("error",err)
	return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	// fmt.Println("resp body",string(body))

	var userDetails model.UserDetails
	err=json.Unmarshal(body,&userDetails)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	fmt.Println("user's team details is :",userDetails.Data.Team)
	// we created Notice array
	var notices []model.Notice


	filter := bson.M{"SendTo": bson.M{"_id":"2"}}
	// filters:=bson.D{"sendTo.id":{"$in":"2"}}

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := db.Collection.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println("error in filter")
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}

	var notice model.Notice

	
	
	er := cur.Decode(&notice) 
	if er != nil {
		log.Fatal(er)
	}

	fmt.Println("cursor",cur)


	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var notice model.Notice

		fmt.Println("inside cursor loop")
		
		err := cur.Decode(&notice) 
		if err != nil {
			log.Fatal(err)
		}

	

		// add item our to array
		notices = append(notices, notice)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	response.ShowResponse("success",int64(utils.HTTP_OK),"get success",notices,ctx)
}


func DeleteNotice(ctx *gin.Context) {

	

	// prepare filter.
	
	url:=ctx.Request.URL
	fmt.Println("url",url)
	T:=strings.Split(url.String(),"/")
	fmt.Println("t",T[len(T)-1])

	filter := bson.M{"_id": T[len(T)-1]}

	deleteResult, err := db.Collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}
	fmt.Println("ERROR 2")
	response.ShowResponse("Success",int64(utils.HTTP_OK),"This record is successfully deleted",deleteResult,ctx)
	
}

func DeleteAllNotices(ctx *gin.Context) {


	filter:=bson.M{}
	result,err:=db.Collection.DeleteMany(ctx,filter)

	fmt.Println("result",result)

	if err!=nil{
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}

	response.ShowResponse("success",int64(utils.HTTP_OK),"delete successfully",result,ctx)


}


func UpdateNotice(ctx *gin.Context,notice_input request.NoticeRequest) {


		// //Get id from parameters
		// id, _ := primitive.ObjectIDFromHex(ctx.Query("id"))
		// fmt.Println("notice input:",notice_input)

		var noticeupdate model.Notice
		
		// Create filter
		filter := bson.M{"_id": ctx.Query("id")}

		noticeupdate.Title=notice_input.Title
		noticeupdate.Text=notice_input.Text
		noticeupdate.From_Date,_=time.Parse("02 Jan 2006",notice_input.From_Date)
		noticeupdate.To_Date,_=time.Parse("02 Jan 2006",notice_input.To_Date)
		noticeupdate.SendTo=notice_input.SendTo


	
	update := bson.M{"$set": bson.M{"title": notice_input.Title,"text":notice_input.Text,"fromDate":noticeupdate.From_Date,"toDate":noticeupdate.To_Date,"sendTo":notice_input.SendTo}}

	fmt.Println("update ",update)
		

	err := db.Collection.FindOneAndUpdate(context.TODO(), filter, update)
	if err.Err()!=nil{
		fmt.Println("server errrrrrrrr",err.Err())

		// response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Err().Error(),"",ctx)
		
		return
	}

	response.ShowResponse("success",int64(utils.HTTP_OK),"Update successful",noticeupdate,ctx)


}