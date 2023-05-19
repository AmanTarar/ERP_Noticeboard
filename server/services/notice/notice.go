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
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func AddNoticeService(ctx *gin.Context,Input_Notice request.NoticeRequest){


	var notice model.Notice
	var userDetails model.UserDetails

	tokn:=ctx.Request.Header
	// fmt.Println("header",tokn["Authorization"])

	if tokn["Authorization"][0]==""{

		response.ShowResponse("Bad Request",int64(utils.HTTP_BAD_REQUEST),"Token not present","",ctx)
		return
	}
	claims,_:=token.DecodeToken(tokn["Authorization"][0])

	

	notice.Creator_id=claims.Id
	fmt.Println("claimsid",claims.Id)

	//get request to database to get user info

	req,err:=http.NewRequest("GET",os.Getenv("GetForUserPath")+claims.Id,ctx.Request.Body)

	req.Header.Add("authorization",tokn["Authorization"][0])
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error",err)
		response.ShowResponse("Bad Request",int64(utils.HTTP_BAD_REQUEST),"error in get request","",ctx)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error",err)
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}

	errr:=json.Unmarshal(body,&userDetails)
	fmt.Println("user's team details is :",userDetails)
	if errr != nil {
		fmt.Println("error",err)
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),errr.Error(),"",ctx)
		return
	}

	
	if !utils.IsAuthorized(userDetails.Data.Role_Data.Role){
		fmt.Println("not authorized to add notice")
		response.ShowResponse("Forbidden",int64(utils.HTTP_FORBIDDEN),"Not enough priviledge","",ctx)
		return
		
		
	}

	
	notice.Text=Input_Notice.Text

	notice.Title=Input_Notice.Title

	notice.SendTo=Input_Notice.SendTo
	
	notice.From_Date,_=time.Parse("02 Jan 2006",Input_Notice.From_Date)

	notice.Created_at=time.Now()

	notice.Creator_name=userDetails.Data.Name 

	noticeToDate,_:=time.Parse("02 Jan 2006",Input_Notice.To_Date)

	notice.To_Date = noticeToDate.Add(24*time.Hour)
	
	notice.Notice_id=primitive.NewObjectID().Hex()

	fmt.Println("notice",notice)

	result,err:= db.Collection.InsertOne(context.TODO(), notice)
	fmt.Println("result id",result)



	if err!=nil{

		fmt.Println("error inserting error in db")
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}

	//update call to users table for new tag

	reqst,_:=http.NewRequest("GET",os.Getenv("UpdateUser")+claims.Id,ctx.Request.Body)

	req.Header.Add("authorization",tokn["Authorization"][0])
	resp, Err := http.DefaultClient.Do(reqst) 
	if Err != nil {
		fmt.Println("error",err)
		response.ShowResponse("Bad Request",int64(utils.HTTP_BAD_REQUEST),"error in update request","",ctx)
		return
	}
	fmt.Println("update user response",resp)

	fmt.Println("result",result)
	response.ShowResponse("success",int64(utils.HTTP_OK),"successully added notice",result,ctx)

	
}





func GetNotice(ctx *gin.Context) {
	

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





func UserGetNotices(ctx *gin.Context) {
	

	fmt.Println("inside get notices")
	// we created Notice array
	var notices []model.Notice
	var userDetails model.UserDetails
	
	tokn:=ctx.Request.Header
	if tokn["Authorization"][0]==""{

		response.ShowResponse("Bad Request",int64(utils.HTTP_BAD_REQUEST),"Token not present","",ctx)
		return
	}
	claims,er:=token.DecodeToken(tokn["Authorization"][0])
	if er!=nil{
		fmt.Println("token decoding failed",er)
		response.ShowResponse("Unauthorized",int64(utils.HTTP_UNAUTHORIZED),er.Error(),"",ctx)

	}

	req,err:=http.NewRequest("GET",os.Getenv("GetForUserPath")+claims.Id,ctx.Request.Body)
	
	req.Header.Add("Authorization",tokn["Authorization"][0])
	res, err := http.DefaultClient.Do(req) //hit on get route
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

	err=json.Unmarshal(body,&userDetails)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	fmt.Println("user's team details is :",userDetails.Data.Team)
	fmt.Println("role id",userDetails.Data.Role_Data.Id)

	var teamIds []string
	for i:=0;i<len(userDetails.Data.Team);i++{

		teamIds=append(teamIds,userDetails.Data.Team[i].Id)

	}
	fmt.Println("team id",teamIds)
	
	filter:=bson.M{"sendTo.id": bson.M{"$in": teamIds}}
	
	
	cur,err:=db.Collection.Find(context.TODO(),filter)


	
	// cur, err := db.Collection.Find(context.TODO(),filter)

	if err != nil {
		fmt.Println("error in filter")
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}
	



	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var notice model.Notice
		
		err := cur.Decode(&notice) 
		if err != nil {
			log.Fatal(err)
		}

		if notice.From_Date.Before(time.Now())&& notice.To_Date.After(time.Now()){
			
			notices = append(notices, notice)
		}

		
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	

	response.ShowResponse("success",int64(utils.HTTP_OK),"get success",notices,ctx)
}


func DeleteNotice(ctx *gin.Context) {

	

	
	var userDetails model.UserDetails
	url:=ctx.Request.URL
	fmt.Println("url",url)
	T:=strings.Split(url.String(),"/")
	fmt.Println("t",T[len(T)-1])


	//token decoding
	tokn:=ctx.Request.Header
	claims,er:=token.DecodeToken(tokn["Authorization"][0])
	if er!=nil{
		fmt.Println("token decoding failed",er)
	}
	fmt.Println("claims",claims.Id)

	//get request to erp db to get user info
	req,err:=http.NewRequest("GET",os.Getenv("GetForUserPath")+claims.Id,ctx.Request.Body)
	
	req.Header.Add("Authorization",tokn["Authorization"][0])
	res, err := http.DefaultClient.Do(req) //hit on get route
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

	err=json.Unmarshal(body,&userDetails)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	//authorization check
	if !utils.IsAuthorized(userDetails.Data.Role_Data.Role){

		response.ShowResponse("Forbidden",int64(utils.HTTP_FORBIDDEN),"Not enough priviledge","",ctx)

		return
	}


	// prepare filter.
	filter := bson.M{"_id": T[len(T)-1]}

	deleteResult, err := db.Collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}
	
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
	var userDetails model.UserDetails



		//token decoding
	tokn:=ctx.Request.Header
	claims,er:=token.DecodeToken(tokn["Authorization"][0])
	if er!=nil{
		fmt.Println("token decoding failed",er)
	}
	fmt.Println("claims",claims.Id)

	//get request to erp db to get user info
	req,err:=http.NewRequest("GET",os.Getenv("GetForUserPath")+claims.Id,ctx.Request.Body)
	
	req.Header.Add("Authorization",tokn["Authorization"][0])
	res, err := http.DefaultClient.Do(req) //hit on get route
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

	err=json.Unmarshal(body,&userDetails)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	//authorization check
	if !utils.IsAuthorized(userDetails.Data.Role_Data.Role){

		response.ShowResponse("Forbidden",int64(utils.HTTP_FORBIDDEN),"Not enough priviledge","",ctx)

		return
	}

		
		// Create filter
		filter := bson.M{"_id": ctx.Query("id")}

		noticeupdate.Title=notice_input.Title
		noticeupdate.Text=notice_input.Text
		noticeupdate.From_Date,_=time.Parse("02 Jan 2006",notice_input.From_Date)
		noticeupdate.To_Date,_=time.Parse("02 Jan 2006",notice_input.To_Date)
		noticeupdate.SendTo=notice_input.SendTo


	
	update := bson.M{"$set": bson.M{"title": notice_input.Title,"text":notice_input.Text,"fromDate":noticeupdate.From_Date,"toDate":noticeupdate.To_Date,"sendTo":notice_input.SendTo}}

	fmt.Println("update ",update)
		

	errr := db.Collection.FindOneAndUpdate(context.TODO(), filter, update)
	if errr.Err()!=nil{
		fmt.Println("server errrrrrrrr",errr.Err())

		// response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Err().Error(),"",ctx)
		
		return
	}

	response.ShowResponse("success",int64(utils.HTTP_OK),"Update successful",noticeupdate,ctx)


}

func GetnoticesHistory(ctx *gin.Context){


	fmt.Println("inside creator get notices")
	// we created Notice array
	var notices []model.Notice
	var userDetails model.UserDetails
	
	tokn:=ctx.Request.Header
	claims,er:=token.DecodeToken(tokn["Authorization"][0])
	if er!=nil{
		fmt.Println("token decoding failed",er)
	}
	fmt.Println("claims",claims.Id)

	req,err:=http.NewRequest("GET",os.Getenv("GetForUserPath")+claims.Id,ctx.Request.Body)
	
	req.Header.Add("Authorization",tokn["Authorization"][0])
	res, err := http.DefaultClient.Do(req) //hit on get route
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

	err=json.Unmarshal(body,&userDetails)
	if err != nil {
		fmt.Println("error",err)
		return
	}
	fmt.Println("user's team details is :",userDetails.Data.Team)
	fmt.Println("role id",userDetails.Data.Role_Data.Id)
	fmt.Println("role name",userDetails.Data.Role_Data.Role)
	


	if !utils.IsAuthorized(userDetails.Data.Role_Data.Role){

		response.ShowResponse("Forbidden",int64(utils.HTTP_FORBIDDEN),"Not enough priviledge","",ctx)

		return 
	}
	filter:=bson.M{}

	var cur *mongo.Cursor
	if ctx.Query("sortKey")==""{

		cur, err = db.Collection.Find(context.TODO(),filter)

	}else{
		sortkey:=ctx.Query("sortKey")
		sortDir:=ctx.Query("sortDirection")
		sortDirInt,_:=strconv.Atoi(sortDir)
	
		opts := options.Find().SetSort(bson.D{{Key: sortkey, Value:sortDirInt}})
		
		
		cur, err = db.Collection.Find(context.TODO(),filter,opts)

	}
	

	if err != nil {
		
		response.ShowResponse("server error",int64(utils.HTTP_INTERNAL_SERVER_ERROR),err.Error(),"",ctx)
		return
	}
	



	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var notice model.Notice
		
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
	
	//sort the notices data before showresponse

	response.ShowResponse("success",int64(utils.HTTP_OK),"get success",notices,ctx)


}