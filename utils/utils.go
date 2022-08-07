package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"os/user"
	"strings"
	"time"
)

func GetStaticDir() (staticDir string) {
	usr, _ := user.Current()
	if usr.HomeDir == "/home/ec2-user" {
		return "./dist"
	}
	return "./front/dist"
}

func getDbConf() (dbConf string) {
	usr, _ := user.Current()
	if usr.HomeDir == "/home/ec2-user" {
		return "myapp:admin@tcp(10.0.2.10:3306)/myapp?charset=utf8mb4"
	}
	return "myapp:admin@/myapp"
}

func LogFirstAccess() {
	fmt.Println(time.Now().String() + " : Server Listening..." + "\n")
}
func LogPostForm(r *http.Request) {
	str, _ := url.QueryUnescape(r.PostForm.Encode())
	array := strings.Split(str, "&")
	fmt.Println(time.Now().String() + " : " + r.URL.String() + " is called." + "  Posted Form is :")
	for _, form := range array {
		fmt.Println(form)
	}
}