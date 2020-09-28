// package gojson provides a method to convert json and key-value
package gojson

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"strings"
)

var kv =  make(map[string]interface{})

// Js2kv convert json to key value
// data: json string
// sep: separation
// return: key-value string
func Js2kv(data, sep string)(string, error){
	if !gjson.Valid(data){
		return "", errors.New("invalid json")
	}
	decoder := json.NewDecoder(strings.NewReader(data))
	decoder.UseNumber()
	jsBox := make(map[string]interface{})
	err := decoder.Decode(&jsBox)
	if err!=nil{
		return "", err
	}
	parse(jsBox, "", sep)
	kvb, err := json.Marshal(kv)
	if err!=nil{
		return "", err
	}
	return string(kvb), nil
}

// Kv2js convert key value to json
// data: key-value string
// sep: separation
// return: json string
func Kv2js(data, sep string)(string, error){
	if !gjson.Valid(data){
		return "", errors.New("invalid json")
	}
	err := json.Unmarshal([]byte(data), &kv)
	if err!=nil{
		return "", err
	}
	js := make(map[string]interface{})
	for k, v := range kv{
		keys := strings.Split(k, sep)
		sub := js
		length := len(keys)
		for i:=0;i<length-1;i++{
			vv, ok := sub[keys[i]]
			if !ok{
				t:=make(map[string]interface{})
				sub[keys[i]] = t
				sub = t
			}else{
				sub = vv.(map[string]interface{})
			}
		}
		sub[keys[length-1]] = v
	}
	jsb, err := json.Marshal(js)
	if err!=nil{
		return "", err
	}
	return string(jsb), nil
}

// recursive parse json
func parse(m map[string]interface{}, prefix, sep string){
	for k, v:=range m{
		key := prefix + k
		switch vv:=v.(type) {
		case map[string]interface{}:
			parse(vv, key + sep, sep)
		default:
			kv[key] = vv
		}
	}
}



