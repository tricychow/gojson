package gojson

import "testing"

func Test_Js2kv_1(t *testing.T){

	var tests = []struct{
		input string
		want string
		sep string
	}{
		{"{\"k1\":{\"age\":25,\"name\":\"tricy\",\"education\":{\"bachelor\":\"CUMTB\",\"master\":\"BUAA\"}}," +
			"\"k2\":null,\"k3\":true,\"k4\":[1,2,3],\"k5\":1.2,\"k6\":\"abc\"}",
			"{\"k1-age\":25,\"k1-education-bachelor\":\"CUMTB\",\"k1-education-master\":\"BUAA\"," +
			"\"k1-name\":\"tricy\",\"k2\":null,\"k3\":true,\"k4\":[1,2,3],\"k5\":1.2,\"k6\":\"abc\"}",
			"-"},
	}
	for _, test := range tests{
		if got, _ := Js2kv(test.input, test.sep);got!=test.want{
			t.Errorf("Js2kv(%v)=%v", test.input, got)
		}
	}
}

func Test_Kv2js_1(t *testing.T){
	var tests = []struct{
		input string
		want string
		sep string
	}{
		{	"{\"k1-age\":25,\"k1-education-bachelor\":\"CUMTB\",\"k1-education-master\":\"BUAA\"," +
			"\"k1-name\":\"tricy\",\"k2\":null,\"k3\":true,\"k4\":[1,2,3],\"k5\":1.2,\"k6\":\"abc\"}",
			"{\"k1\":{\"age\":25,\"education\":{\"bachelor\":\"CUMTB\",\"master\":\"BUAA\"}," +
			"\"name\":\"tricy\"},\"k2\":null,\"k3\":true,\"k4\":[1,2,3],\"k5\":1.2,\"k6\":\"abc\"}",
			"-"},
	}
	for _, test := range tests{
		if got, _ := Kv2js(test.input, test.sep);got!=test.want{
			t.Errorf("Kv2js(%v)=%v", test.input, got)
		}
	}
}

