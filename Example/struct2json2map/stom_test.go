package struct2json2map

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/stretchr/testify/require"
	"testing"
)

// 结构体 => json => map

type Person struct {
	Name  string `json:"name" structs:"name"`
	Age   int    `json:"age" structs:"age"`
	Hobby string `json:"hobby" structs:"hobby"`
}

func TestT1(t *testing.T) {
	s1 := Person{
		"mzh",
		24,
		"football",
	}

	// struct to json
	jsonRet, jsonErr := json.Marshal(s1)
	require.Equal(t, jsonErr, nil)
	fmt.Println("jsonRet: ", string(jsonRet))

	fmt.Println("=========================")

	// json to map
	var mapRet map[string]interface{}
	errUnmarshal := json.Unmarshal(jsonRet, &mapRet)
	require.Equal(t, errUnmarshal, nil)
	fmt.Println("mapRet: ", mapRet)

	fmt.Println("=========================")

	// todo 有一个坑：value的类型
	for key, val := range mapRet {
		fmt.Printf("key: %v, val: %v, typeOfValue: %T \n", key, val, val)
	}
}

func TestT2(t *testing.T) {
	s1 := Person{
		"mzh",
		24,
		"football",
	}

	mapRet := structs.Map(s1)
	fmt.Println("mapRet: ", mapRet)

	fmt.Println("=========================")

	for key, val := range mapRet {
		fmt.Printf("key: %v, val: %v, typeOfValue: %T \n", key, val, val)
	}
}
