package main

import (
	"net/http"
	"uploadFilePJ/utils"

	"uploadFilePJ/data"
)

func postFile(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fileName := r.PostFormValue("fileName")
	base64Data := r.PostFormValue("base64Data")
	mimeTypeString := r.PostFormValue("mimeTypeString")
	userId := r.PostFormValue("userId")
	postedDate := r.PostFormValue("postedDate")
	isOK := data.PostFIleData(fileName,base64Data,mimeTypeString,userId,postedDate)
	if isOK {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getFileInfoByUserId(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	utils.LogPostForm(r)
	userId := r.PostFormValue("userId")
	date := r.PostFormValue("date")
	isOK,output:= data.GetFileDataByUserId(userId,date)
	if isOK {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}