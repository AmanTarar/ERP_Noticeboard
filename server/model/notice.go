package model

import (
	"time"
)


type Notice struct{
	
	Notice_id  string `bson:"_id"`
	
	Created_at time.Time `bson:"createdAt"`
	Title string `bson:"title"`
	Text string `bson:"text"`
	Creator_id string `bson:"creatorId"`
	Creator_name string `bson:"creatorName"`
	Role    string `bson:"role"`
	From_Date time.Time `bson:"fromDate"`
	To_Date time.Time `bson:"toDate"`
	SendTo []TeamID `bson:"sendTo"`	
	
}

type TeamID struct{

	Name string `json:"name"`
	Id string `json:"_id"` 
}


type UserDetails struct{
	Data struct {
		Email string `json:"email"`
		Name string `json:"name"`
		Team []struct {
			Id string `json:"_id"`
			Name string `json:"name"`
		} `json:"teams"`
		Role_Data struct{
			Role string `json:"role"`
			Id string `json:"_id"`
		} `json:"roleData"`
	}`json:"data"`


}






// {
//     "statusCode": 200,
//     "message": "action successful",
//     "status": true,
//     "type": "Default",
//     "data": {
//         "_id": "63ce887193e913067d03127b",
//         "email": "amantarar01@gmail.com",
//         "name": "Aman Tarar",
//         "fatherName": "Gajendra Singh",
//         "phones": [
//             {
//                 "contact": "6280912015",
//                 "primary": true,
//                 "verified": false,
//                 "_id": "63ce887193e913067d03127c"
//             },
//             {
//                 "contact": "",
//                 "primary": false,
//                 "verified": false,
//                 "_id": "63ce887193e913067d03127d"
//             },
//             {
//                 "contact": "6280912015",
//                 "primary": true,
//                 "verified": true,
//                 "_id": "63ce889093e913067d0312d1"
//             },
//             {
//                 "contact": "",
//                 "primary": false,
//                 "verified": true,
//                 "_id": "63ce889093e913067d0312d2"
//             }
//         ],
//         "gender": "male",
//         "joinedAs": 1,
//         "experience": {
//             "totalWorkExperience": {
//                 "years": null,
//                 "months": null
//             },
//             "releventWorkExperience": {
//                 "years": null,
//                 "months": null
//             }
//         },
//         "education": [
//             {
//                 "qualificationId": null,
//                 "college": null,
//                 "passOutYear": null
//             }
//         ],
//         "previousCompanies": [
//             {
//                 "name": "",
//                 "designation": "",
//                 "salary": null,
//                 "salarySlips": [],
//                 "_id": "63edd138d54f91f127f637e0"
//             }
//         ],
//         "status": true,
//         "active": true,
//         "designationId": "5e33d9b611f3a90cf90a0f4a",
//         "addresses": [
//             {
//                 "line1": null,
//                 "line2": null,
//                 "city": "",
//                 "state": "",
//                 "country": "India",
//                 "postalCode": null,
//                 "type": 1,
//                 "verified": false,
//                 "active": false,
//                 "_id": "63ce887193e913067d03127f"
//             },
//             {
//                 "line1": null,
//                 "line2": null,
//                 "city": "",
//                 "state": "",
//                 "country": "India",
//                 "postalCode": null,
//                 "type": 2,
//                 "verified": false,
//                 "active": false,
//                 "_id": "63ce887193e913067d031280"
//             },
//             {
//                 "line1": null,
//                 "line2": null,
//                 "city": "",
//                 "state": "",
//                 "country": "India",
//                 "postalCode": null,
//                 "type": 1,
//                 "verified": true,
//                 "active": true,
//                 "_id": "63ce889093e913067d0312db"
//             },
//             {
//                 "line1": null,
//                 "line2": null,
//                 "city": "",
//                 "state": "",
//                 "country": "India",
//                 "postalCode": null,
//                 "type": 2,
//                 "verified": true,
//                 "active": true,
//                 "_id": "63ce889093e913067d0312dc"
//             }
//         ],
//         "aadharPictures": [
//             "CHM/2023/589/otherDocs/1676529874547.jpeg"
//         ],
//         "emergencyContacts": [
//             {
//                 "index": 1,
//                 "name": "",
//                 "relation": "",
//                 "contact": "",
//                 "verified": false,
//                 "_id": "63ce887193e913067d031281"
//             },
//             {
//                 "index": 2,
//                 "name": "",
//                 "relation": "",
//                 "contact": "",
//                 "verified": false,
//                 "_id": "63ce887193e913067d031282"
//             },
//             {
//                 "index": 1,
//                 "name": "",
//                 "relation": "",
//                 "contact": "",
//                 "verified": true,
//                 "_id": "63ce889093e913067d0312d6"
//             },
//             {
//                 "index": 2,
//                 "name": "",
//                 "relation": "",
//                 "contact": "",
//                 "verified": true,
//                 "_id": "63ce889093e913067d0312d7"
//             }
//         ],
//         "teams": [
//             {
//                 "_id": "642fc5e87385415e2a7c7936",
//                 "name": "GOLANG",
//                 "isActive": true,
//                 "users": [],
//                 "isDeleted": false,
//                 "createdAt": "2023-04-07T07:27:36.275Z",
//                 "updatedAt": "2023-04-07T07:27:36.275Z"
//             }
//         ],
//         "roleId": "5e3179cba6fd0d4f78b06f27",
//         "employeeId": "CHM/2023/589",
//         "joiningDate": "2023-01-25T00:00:00.000Z",
//         "officialEmail": {
//             "email": "aman.tarar@yopmail.com",
//             "password": ""
//         },
//         "skypeId": {
//             "email": "aman.tarar@chicmic.co.in",
//             "password": ""
//         },
//         "hubstaffId": {
//             "email": "aman.tarar@chicmic.co.in",
//             "password": ""
//         },
//         "isDeleted": false,
//         "verificationStatus": 3,
//         "bankInfo": [
//             {
//                 "name": "State Bank Of India",
//                 "account": "41611504679",
//                 "ifsc": "SBIN0050022",
//                 "active": false,
//                 "_id": "63edd138d54f91f127f637db"
//             }
//         ],
//         "verifiedAllAssets": false,
//         "minInTime": "10:00",
//         "maxInTime": "10:30",
//         "isPartTime": false,
//         "leaveTakenInResignation": 0,
//         "workingAt": 1,
//         "policyVersions": [
//             {
//                 "policyNumber": 1,
//                 "policyVersionNumber": 1,
//                 "policyAccepted": true,
//                 "_id": "64647aeb9658c3f95d1222c3"
//             },
//             {
//                 "policyNumber": 1,
//                 "policyVersionNumber": 1.1,
//                 "policyAccepted": true,
//                 "_id": "64647aeb9658c3f95d1222c4"
//             },
//             {
//                 "policyNumber": 1,
//                 "policyVersionNumber": 1.2,
//                 "policyAccepted": true,
//                 "_id": "64647aeb9658c3f95d1222c5"
//             }
//         ],
//         "updatedAt": "2023-05-17T06:57:47.648Z",
//         "dateOfBirth": "2001-07-12T00:00:00.000Z",
//         "aadharNumber": "2678 9965 7441",
//         "panNumber": "CBPPT2512Q",
//         "panPicture": "CHM/2023/589/otherDocs/1676529877819.pdf",
//         "proflePicture": "CHM/2023/589/dp/1676529869842.jpeg",
//         "isRelaxationAllowed": false,
//         "mentorId": "61fba872f4f70d6c0b3effbd",
//         "organizationType": 2,
//         "overallPolicyAccepted": true,
//         "designationData": {
//             "_id": "5e33d9b611f3a90cf90a0f4a",
//             "status": true,
//             "isDeleted": false,
//             "name": "Trainee",
//             "__v": 0,
//             "isActive": true
//         },
//         "roleData": {
//             "_id": "5e3179cba6fd0d4f78b06f27",
//             "name": "Individual",
//             "role": "IND",
//             "permissions": [
//                 "canReadOwnProject"
//             ],
//             "description": "Permissions for Individual \n  - Can view his allocated projects.\n  - Can create, view & update  his Timesheet \n  - Can view Timesheet History",
//             "status": true
//         },
//         "baseUrl": "https://cm-erp-files.s3.ap-south-1.amazonaws.com/"
//     }
// }