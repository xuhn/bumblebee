package api

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/xuhn/optimusprime/log"
	"github.com/xuhn/optimusprime/net"
	"github.com/xuhn/optimusprime/task"
)

type APIRequest struct {
	Action      string `json:"Action"`
	RequestUUID string `json:"request_uuid"`
}

func init() {
	// 注册HTTP任务
	net.RouteHTTP = RouteHTTP
	task.RegisterHTTPTaskHandle("GetItem", http.HandlerFunc(GetItem), 20*time.Second)
	task.RegisterHTTPTaskHandle("ListItem", http.HandlerFunc(ListItem), 20*time.Second)
	task.RegisterHTTPTaskHandle("Signature", http.HandlerFunc(Signature), 20*time.Second)
	task.RegisterHTTPTaskHandle("UserLogin", http.HandlerFunc(UserLogin), 20*time.Second)
}

func RouteHTTP(w http.ResponseWriter, r *http.Request) {
	log.DEBUGF("request: %v", r)

	action := r.URL.Query().Get("Action")
	requestUUID := r.URL.Query().Get("request_uuid")
	if r.Method == "POST" {
		contentType := ""
		contentTypes, ok := r.Header["Content-Type"]
		if ok && len(contentTypes) > 0 {
			contentType = contentTypes[0]
		}
		log.DEBUGF("contentType: %s", contentType)
		if contentType == "application/json" {
			safe := &io.LimitedReader{R: r.Body, N: 1 << 26}
			requestbody, _ := ioutil.ReadAll(safe)
			r.Body.Close()
			bf := bytes.NewBuffer(requestbody)
			r.Body = ioutil.NopCloser(bf)

			apiRequest := new(APIRequest)
			json.Unmarshal(requestbody, apiRequest)

			action = apiRequest.Action
			requestUUID = apiRequest.RequestUUID
		} else if contentType == "application/x-www-form-urlencoded" {
			safe := &io.LimitedReader{R: r.Body, N: 1 << 26}
			requestbody, _ := ioutil.ReadAll(safe)
			r.Body.Close()
			bf := bytes.NewBuffer(requestbody)
			r.Body = ioutil.NopCloser(bf)

			apiRequest := new(APIRequest)
			xml.Unmarshal(requestbody, apiRequest)

			action = apiRequest.Action
			requestUUID = apiRequest.RequestUUID

		} else if strings.Contains(contentType, "multipart/form-data") {
			r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("uploadfile")
			log.DEBUGF("form: %v", r.Form)
			action = r.Form.Get("Action")
			log.DEBUGF("action: %s", action)
			if err != nil {
				log.ERRORF("err: %v", err)
				return
			}
			defer file.Close()
			fmt.Fprintf(w, "%v", handler.Header)
			f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.ERRORF("err: %v", err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		} else {
			r.ParseForm()
			action = r.PostForm.Get("Action")
			requestUUID = r.PostForm.Get("request_uuid")
		}
	} else if r.Method == "GET" {
		r.ParseForm()
		action = "Signature"
	}
	w.Header().Set("request_uuid", requestUUID)
	r.Header["request_uuid"] = []string{requestUUID}

	if action == "" {
		log.ERROR("can not find action")
		OutputResponse(w, generateErrorCode("MISSING_ACTION"))
		return
	}

	task, err := task.NewHTTPTask(action)

	if err != nil {
		log.DEBUG("new task fail, ", err)
		OutputResponse(w, generateErrorCode("MISSING_ACTION"))
		return
	}

	_, err = task.Run(w, r)

	if err != nil {
		log.ERRORF("run task(%s) fail:%s", task.Id, err)
		return
	}
}
