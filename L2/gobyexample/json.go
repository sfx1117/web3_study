package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Friuts []string
}

type response2 struct {
	Page   int      `json:"page"`
	Friuts []string `json:"fruits"`
}

func main() {
	boolB, _ := json.Marshal(true)
	fmt.Println(string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	flB, _ := json.Marshal(2.34)
	fmt.Println(string(flB))
	strB, _ := json.Marshal("fsdrwe34")
	fmt.Println(string(strB))

	slc := []string{"ass", "fdd", "dfsre"}
	slcB, _ := json.Marshal(slc)
	fmt.Println(string(slcB))

	mapp := map[string]string{"aaa": "111", "bbb": "222"}
	mapB, _ := json.Marshal(mapp)
	fmt.Println(string(mapB))

	resD1 := response1{
		Page:   1,
		Friuts: []string{"aaaa", "bbbb", "cccc"},
	}
	resB1, _ := json.Marshal(resD1)
	fmt.Println(string(resB1))

	resD2 := response2{
		Page:   1,
		Friuts: []string{"aaaa", "bbbb", "cccc"},
	}
	resB2, _ := json.Marshal(resD2)
	fmt.Println(string(resB2))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	err := json.Unmarshal(byt, &dat)
	if err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Friuts[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
}

//true
//1
//2.34
//"fsdrwe34"
//["ass","fdd","dfsre"]
//{"aaa":"111","bbb":"222"}
//{"Page":1,"Friuts":["aaaa","bbbb","cccc"]}
//{"page":1,"fruits":["aaaa","bbbb","cccc"]}
//map[num:6.13 strs:[a b]]
//6.13
//a
//{1 [apple peach]}
//apple
//{"apple":5,"lettuce":7}
//{1 [apple peach]}
