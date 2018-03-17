package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"optimusprime/log"
)

func OutputResponse(w http.ResponseWriter, r interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	b, err := json.Marshal(r)
	if err != nil {
		return
	}

	w.Write(b)
	log.INFOF(" HTTP_RESPONSE (%s)\n    %+v", w.Header().Get("request_uuid"), r)
	return
}

func InputRequest(r *http.Request, ptr interface{}) error {
	body := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		log.ERRORF("InputRequest: %s", err)
		return fmt.Errorf("")
	}

	data := make(map[string]string)
	for k, _v := range body {
		switch v := _v.(type) {
		case string:
			data[k] = v
		case float64:
			data[k] = strconv.FormatFloat(v, 'f', -1, 64)
		}
	}
	dataArr := make(map[string][]string)
	re, err := regexp.Compile(`[A-Za-z0-9]\.[0-9]`)
	if err != nil {
		log.ERRORF("InputRequest: %s", err)
		return fmt.Errorf("")
	}
	for k, v := range data {
		if re.MatchString(k) {
			_k := strings.Split(k, ".")[0]
			if _, ok := dataArr[_k]; !ok {
				dataArr[_k] = []string{v}
			} else {
				dataArr[_k] = append(dataArr[_k], v)
			}
		}
	}

	v := reflect.ValueOf(ptr).Elem()
	for i, l := 0, v.NumField(); i < l; i++ {
		fieldTag := v.Type().Field(i).Tag
		fieldValue := v.Field(i)
		if !fieldValue.CanSet() {
			continue
		}
		key := fieldTag.Get("key")
		defaultValue := fieldTag.Get("default")
		required := fieldTag.Get("required")
		value, ok := data[key]
		if !ok {
			value = defaultValue
			if required != "" {
				if _, _ok := dataArr[key]; !_ok {
					return fmt.Errorf("%s", key)
				}
			}
		}
		if fieldValue.Kind() == reflect.Slice {
			if value, ok := dataArr[key]; ok {
				for _, _v := range value {
					elem := reflect.New(fieldValue.Type().Elem()).Elem()
					if err := populate(elem, _v); err != nil {
						log.ERRORF("InputRequest: %s", err)
						return fmt.Errorf("%s", key)
					}
					fieldValue.Set(reflect.Append(fieldValue, elem))
				}
			}
		} else {
			if err := populate(fieldValue, value); err != nil {
				log.ERRORF("InputRequest: %s", err)
				return fmt.Errorf("%s", key)
			}
		}
	}

	log.INFOF(" HTTP_REQUEST (%s) | (%s:%s)\n    %+v", data["request_uuid"], data["client_ip"], data["Action"], ptr)
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Uint32:
		i, _ := strconv.ParseUint(value, 10, 32)
		v.SetUint(i)

	case reflect.Int:
		i, _ := strconv.ParseInt(value, 10, 64)
		v.SetInt(i)

	case reflect.Bool:
		b, _ := strconv.ParseBool(value)
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
