package gojson

import "testing"

func Test_Js2kv_1(t *testing.T){
	data := "{\"k1\":{\"age\":25,\"name\":\"tricy\",\"education\":{\"bachelor\":\"CUMTB\",\"master\":\"BUAA\"}}," +
		"\"k2\":null,\"k3\":true,\"k4\":[1,2,3],\"k5\":1.2,\"k6\":\"abc\"}"
	kv, err := Js2kv(data, "-")
	if err != nil{
		t.Error("Js2kvc测试出错")
	}else{
		if kv != "{\"k1-age\":25,\"k1-education-bachelor\":\"CUMTB\",\"k1-education-master\":\"BUAA\"," +
			"\"k1-name\":\"tricy\",\"k2\":null,\"k3\":true,\"k4\":[1,2,3],\"k5\":1.2,\"k6\":\"abc\"}"{
			t.Error("Js2kvc测试出错")
		}else{
			t.Log("Js2kvc测试通过")
		}
	}
}

func Test_Kv2js_1(t *testing.T){
	data := "{\"k1-age\":25,\"k1-name\":\"tricy\",\"k1-education-bachelor\":\"CUMTB\",\"k1-education-master\":\"BUAA\"," +
		"\"k2\":null,\"k3\":true,\"k4\":[1,2,3],\"k5\":1.2,\"k6\":\"abc\"}"
	js, err := Kv2js(data, "-")
	if err != nil{
		t.Error("Kv2js测试出错")
	}else{
		if js != "{\"k1\":{\"age\":25,\"education\":{\"bachelor\":\"CUMTB\",\"master\":\"BUAA\"},\"name\":\"tricy\"}," +
		"\"k2\":null,\"k3\":true,\"k4\":[1,2,3],\"k5\":1.2,\"k6\":\"abc\"}"{
			t.Error("Kv2js测试出错")
		}else{
			t.Log("Kv2js测试通过")
		}
	}
}

