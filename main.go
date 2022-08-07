package main

import (
	"net/http"
	"uploadFilePJ/utils"
)


func main()  {
	server := http.Server{
		Addr: ":8082",
	}
	http.Handle("/", http.FileServer(http.Dir(utils.GetStaticDir())))
	http.HandleFunc("/api/v1/postFile", postFile)
	http.HandleFunc("/api/v1/getFileInfoByUserId", getFileInfoByUserId)
	utils.LogFirstAccess()
	server.ListenAndServe()
}