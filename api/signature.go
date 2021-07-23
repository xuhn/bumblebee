/*
*author:cszixin.liu@ucloud.cn
*function:在vip 2.0 下申请vip
time:2016.8.28
*/
package api

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"

	_ "github.com/xuhn/bumblebee/dao"

	"github.com/xuhn/optimusprime/log"
)

type (
	SignatureRequest struct {
		Signature string `key:"signature" required:"true"`
		Timestamp string `key:"timestamp " required:"true"`
		Nonce     string `key:"nonce" required:"true"`
		Echostr   string `key:"echostr" required:"true"`
	}
)

func Signature(w http.ResponseWriter, r *http.Request) {
	req := new(SignatureRequest)
	r.ParseForm()
	req.Signature = strings.Join(r.Form["signature"], "")
	req.Timestamp = strings.Join(r.Form["timestamp"], "")
	req.Nonce = strings.Join(r.Form["nonce"], "")
	req.Echostr = strings.Join(r.Form["echostr"], "")

	apiRes := map[string]interface{}{}

	token := "hello_world"
	list := []string{req.Nonce, req.Timestamp, token}
	sort.Strings(list)

	log.DEBUGF("handle/GET func | list: %v", list) //获取请求的方法
	hashcode := convsha1(strings.Join(list, ""))

	log.DEBUGF("handle/GET func | hashcode: %s , signature: %s", hashcode, req.Signature) //获取请求的方法

	if hashcode == req.Signature {
		w.Write([]byte(req.Echostr))
		return
	}

	OutputResponse(w, apiRes)
	return
}

//对字符串进行SHA1哈希
func convsha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
