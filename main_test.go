package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"uploadFilePJ/utils"

	"github.com/stretchr/testify/assert"
)

type UploadedFileInfoJson struct{
	FileInfoId string `json:"fileInfoId"`
	FileName string `json:"fileName"`
	FileContent string `json:"fileContent"`
	MimeType string `json:"mimeType"`
	UserId string `json:"userId"`
	PostedDate string `json:"postedDate"`
}
func TestIndexHandler(t *testing.T) {
    // テスト用ハンドラ作成
    mux := http.NewServeMux()
    mux.Handle("/",http.FileServer(http.Dir(utils.GetStaticDir())))
    // /のリクエスト用テストコード
    req := httptest.NewRequest("GET", "/", nil)
    res := httptest.NewRecorder()
    mux.ServeHTTP(res,req)
    assert.Equal(t, http.StatusOK, res.Code)
}
func TestPostFileHandler1(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/postFile", postFile)
    values := url.Values{}
    values.Set("fileName", "test.txt")
    values.Add("base64Data", "test")
    values.Add("mimeTypeString", "text/plain")
    values.Add("userId", "001")
    values.Add("postedDate", "2022/8/7")
    
    req := httptest.NewRequest("POST","/api/v1/postFile",strings.NewReader(values.Encode()),)
    res := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    mux.ServeHTTP(res,req)
    assert.Equal(t, http.StatusNoContent, res.Code)
}
func TestPostFileHandler2(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/postFile", postFile)
    values := url.Values{}
    values.Set("fileName", "test.txt")
    values.Add("base64Data", "test")
    values.Add("mimeTypeString", "text/plain")
    values.Add("userId", "001")

    req := httptest.NewRequest("POST","/api/v1/postFile",strings.NewReader(values.Encode()),)
    res := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    mux.ServeHTTP(res,req)
    assert.Equal(t, http.StatusNoContent, res.Code)
}
func TestPostFileHandler3(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/postFile", postFile)
    values := url.Values{}
    values.Set("fileName", "test.txt")
    values.Add("base64Data", "test")
    values.Add("mimeTypeString", "text/plain")

    req := httptest.NewRequest("POST","/api/v1/postFile",strings.NewReader(values.Encode()),)
    res := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    mux.ServeHTTP(res,req)
    assert.Equal(t, http.StatusNoContent, res.Code)
}
func TestPostFileHandler4(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/postFile", postFile)
    values := url.Values{}
    values.Set("fileName", "test.txt")
    values.Add("base64Data", "test")

    req := httptest.NewRequest("POST","/api/v1/postFile",strings.NewReader(values.Encode()),)
    res := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    mux.ServeHTTP(res,req)
    assert.Equal(t, http.StatusNoContent, res.Code)
}
func TestPostFileHandler5(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/postFile", postFile)
    values := url.Values{}
    values.Set("fileName", "test.txt")
    req := httptest.NewRequest("POST","/api/v1/postFile",strings.NewReader(values.Encode()),)
    res := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    mux.ServeHTTP(res,req)
    assert.Equal(t, http.StatusNoContent, res.Code)
}
func TestPostFileHandler6(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/postFile", postFile)
    values := url.Values{}
    req := httptest.NewRequest("POST","/api/v1/postFile",strings.NewReader(values.Encode()),)
    res := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    mux.ServeHTTP(res,req)
    assert.Equal(t, http.StatusNoContent, res.Code)
}
func TestPostFileHandler7(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/postFile", postFile)
    values := url.Values{}
    values.Set("fileName", "")
    values.Add("base64Data", "")
    values.Add("mimeTypeString", "")
    values.Add("userId", "")
    values.Add("postedDate", "")
    
    req := httptest.NewRequest("POST","/api/v1/postFile",strings.NewReader(values.Encode()),)
    res := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    mux.ServeHTTP(res,req)
    assert.Equal(t, http.StatusNoContent, res.Code)
}
func TestGetFileInfoByUserId1(t *testing.T) {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/getFileInfoByUserId", getFileInfoByUserId)
    values := url.Values{}
    values.Set("userId", "001")
    values.Add("date", "2022/8/7")
    req := httptest.NewRequest("POST","/api/v1/getFileInfoByUserId",strings.NewReader(values.Encode()),)
    res := httptest.NewRecorder()
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    mux.ServeHTTP(res,req)
    // 実行結果をJSONで取得
    var infoJsonList []UploadedFileInfoJson
    json.Unmarshal(res.Body.Bytes(),&infoJsonList)
    // 期待結果を作成したJSONファイルから取得
    jsonFile,err := os.Open("./jsonFileForTest/getFileInfo.json")
    if err != nil{
		panic(err)
	}
	jsonData ,err := ioutil.ReadAll(jsonFile)
	if err != nil{
		panic(err)
	}
	var infoJsonListExpected []UploadedFileInfoJson
    json.Unmarshal(jsonData,&infoJsonListExpected)

    assert.Equal(t,infoJsonListExpected,infoJsonList)
}
