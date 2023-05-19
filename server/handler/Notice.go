package handler

import (
	"main/server/request"
	"main/server/services/notice"

	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func AddNoticeHandler(context *gin.Context) {


	utils.SetHeader(context)


	var notice_input request.NoticeRequest
	utils.RequestDecoding(context,&notice_input)
	notice.AddNoticeService(context,notice_input)

}

func GetNoticeHandler(context *gin.Context) {


	utils.SetHeader(context)

	notice.GetNotice(context)
}


func UpdateNoticeHandler(context *gin.Context){


	utils.SetHeader(context)
	var notice_input request.NoticeRequest
	utils.RequestDecoding(context,&notice_input)
	notice.UpdateNotice(context,notice_input)
}

func DeleteNoticeHandler(context *gin.Context){


	utils.SetHeader(context)


	notice.DeleteNotice(context)
}

func GetNoticesHandler(context *gin.Context){

	utils.SetHeader(context)

	notice.UserGetNotices(context)
}

func DeleteAllNoticesHandler(context *gin.Context){

	utils.SetHeader(context)

	notice.DeleteAllNotices(context)
}

func GetNoticesHistoryHandler(context *gin.Context){

	utils.SetHeader(context)

	notice.GetnoticesHistory(context)
}

// func DecodeTOKEN(context *gin.Context){


// 	utils.SetHeader(context)
// 	// tokn:=context.Request.Header.Get("token")
// 	tokn:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2M2NlNjZiZTkzZTkxMzA2N2QwMjg4YjgiLCJlbWFpbCI6Iml0dWRhZGh3YWwxNzI3QGdtYWlsLmNvbSIsInRpbWUiOjE2ODQ0MDE4ODI3ODUsImlhdCI6MTY4NDQwMTg4Mn0.yLyUnoMCzWLCUlLZdmc1ZYky3jeKcKE7ik5Xc_Q0Cp8"
// 	claims,_:=token.DecodeToken(tokn)
// 	fmt.Println("claims",claims)
// }



