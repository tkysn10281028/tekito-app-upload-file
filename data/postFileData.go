package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/user"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
type UploadedFileInfoModel struct{
	FileInfoId sql.NullString
	FileName sql.NullString
	FileContent sql.NullString
	MimeType sql.NullString
	UserId sql.NullString
	PostedDate sql.NullString
}

type UploadedFileInfoJson struct{
	FileInfoId string `json:"fileInfoId"`
	FileName string `json:"fileName"`
	FileContent string `json:"fileContent"`
	MimeType string `json:"mimeType"`
	UserId string `json:"userId"`
	PostedDate string `json:"postedDate"`
}


func init() {
	usr ,_ := user.Current()
	var dbConf string
	if(usr.HomeDir =="/home/ec2-user"){
		dbConf = "myapp:@Pleasure1@tcp(10.0.2.10:3306)/myapp?charset=utf8mb4"
	}else{
		dbConf =  "myapp:@Pleasure1@/myapp"
	}
	var err error
	Db, err = sql.Open("mysql", dbConf)
	if err != nil {
		panic(err)
	}
}



func PostFIleData(fileName string,fileContent string,mimeType string,userId string,postedDate string) (bool) {
	statement := postFileSQL()
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()
	result, err := stmt.Exec(fileName,fileContent,mimeType,userId,postedDate)
	if err != nil {
		fmt.Println(err,result)
		return false
	}
	return true
}

func GetFileDataByUserId(userId string,date string) (bool,[]byte){
	
	infoJsonList := []UploadedFileInfoJson{}
	statement := getFileDataByUserIdSQL()
	rows, err := Db.Query(statement,userId,date)

	if err != nil{
		fmt.Println(err)
		return false,nil
	}
	defer rows.Close()
	for rows.Next(){
		infoModel := UploadedFileInfoModel{}
		err = rows.Scan(
			&infoModel.FileInfoId,
			&infoModel.FileName,
			&infoModel.FileContent,
			&infoModel.MimeType,
			&infoModel.UserId,
			&infoModel.PostedDate)
		if err != nil{
		fmt.Println(err)
		return false,nil
		}
		infoJson := UploadedFileInfoJson{
			FileInfoId:infoModel.FileInfoId.String,
			FileName:infoModel.FileName.String,
			FileContent:infoModel.FileContent.String,
			MimeType: infoModel.MimeType.String,
			UserId: infoModel.UserId.String,
			PostedDate: infoModel.PostedDate.String,
		}
		infoJsonList = append(infoJsonList, infoJson)
	}
	output , err := json.MarshalIndent(&infoJsonList,"","\t\t")
	if err != nil{
		fmt.Println(err)
		return false,nil
	}
	return true,output
}

func postFileSQL()(string){
	return`
	INSERT INTO 
		UPLOADED_FILE_INFO
   		(
   			FILE_NAME,
   			FILE_CONTENT,
   			MIME_TYPE,
			USER_ID,
			POSTED_DATE
		)
   		VALUES(
			?,
   			?,
			?,
			?,
			?
		)
	`
}

func getFileDataByUserIdSQL() (string) {
	return `
	SELECT
		FILE_INFO_ID,
		FILE_NAME,	
		FILE_CONTENT,
		MIME_TYPE,
		USER_ID,
		POSTED_DATE
	FROM
		UPLOADED_FILE_INFO
	WHERE
		USER_ID = ?
	AND
		POSTED_DATE = ?
	ORDER BY FILE_INFO_ID DESC
	LIMIT 10
	`
}