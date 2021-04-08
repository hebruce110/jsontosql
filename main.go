/**
Author: Bruce
*/
package main

import (
	"flag"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
	"io/ioutil"
	"reflect"
	"strings"
)

func main() {
	flag.String("Tool","","Convert json string to sql statement to create table")
	jsonStr := flag.String("json","","json string, for example -json '{\"key\":\"value\"}'")
	jsonPath := flag.String("json_path","","json file path,for example /tmp/test.json。If json parameter is specified, the content of json parameter will be read first")
	flag.Parse()
	if *jsonStr == "" {
		jsonBytes,err  := ioutil.ReadFile(*jsonPath)
		if err != nil {
			fmt.Println("Failed to open the json file, please specify the correct json file path!")
			panic(err)
		}
		*jsonStr = string(jsonBytes)
	}
	var jsonDict map[string]interface{}
	err := jsoniter.Unmarshal([]byte(*jsonStr), &jsonDict)
	if err != nil{
		fmt.Println(err)
		panic("Parsing failed, please check the json format!")
	}
	tablesColumn := []string{}
	for k,value := range jsonDict {
		var dataType string
		switch reflect.TypeOf(value).Kind() {
		case reflect.Array,reflect.Slice,reflect.Map:
			dataType = "Array"
		case reflect.String:
			dataType = "String"
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int8, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			dataType = "UInt64"
		case reflect.Float32,reflect.Float64:
			dataType = "Float64"
			//When converting json to map, int will be judged as float, if there is no . symbol, correct to int type
			if !strings.Contains(cast.ToString(value),".") {
				dataType = "UInt64"
			}
		case reflect.Bool:
			dataType = "UInt8"
		default:
			dataType = "String"
		}
		tablesColumn = append(tablesColumn,fmt.Sprintf("`%s` %s",k,dataType))
		//fmt.Println(k,"-----------",reflect.TypeOf(value).Kind(),"-----------",reflect.TypeOf(value).Name())
	}
	sql := fmt.Sprintf("\n**********JSON CONVERT TO CREATE TABLE SQL **********\n\nCREATE TABLE IF NOT EXISTS jsonTableName (\n%s\n)\n\n******************************",strings.Join(tablesColumn,",\n"))
	fmt.Println(sql)
}

