package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
)

const (
	NOCOVER = iota
	SETCAMERA
	COVERED
)

func main() {

	//a := 1<<64 - 1

	//b := math.MaxUint64

	//fmt.Printf("type:%T value:%v", a, a)
	//fmt.Printf("type:%T value:%v", b, b)
	//	fmt.Println('1' - '0')
	//	fmt.Printf("type: %T\n", '0')
	//	a := 8
	//	a = (a - 1) & a

	//	fmt.Printf(strconv.Itoa(a))
	//fmt.Println(NOCOVER, SETCAMERA, COVERED)

	//a := map[string]int{"test": 123, "abc": 222}
	//b, c := a["test1"]
	//fmt.Println(a["test1"])
	//fmt.Println(b, c)

	/*for a, b := range "test" {
		fmt.Println(a, b)
	}*/

	/*c := []int{1, 2, 3}
	  fmt.Println(c[len(c):])
	  fmt.Println(string(79))
	  fmt.Println(string(88))
	  fmt.Println('H')*/

	/*type S int
	  a := S(0)
	  b := make([]*S, 2)
	  b[0] = &a
	  c := new(S)
	  b[1] = c*/

	//type TestValue struct {
	//    IntValue    int      `qiudianzan:"int_value"`
	//    StringValue string   `qiudianzan:"string_value"`
	//    IntArray    []int    `qiudianzan:"int_array"`
	//    StringArray []string `qiudianzan:"string_array"`
	//}

	//fmt.Println([]int{1, 2, 3}[0:2])

	//fmt.Println('a' - 'c')

	//for idx, str := range "testabc" {
	//    fmt.Println(idx, str)
	//}

	//a := map[int]int{123: 678}
	//b, c := a[123]
	//d, e := a[111]
	//g := a[123]
	//fmt.Println(b, c)
	//fmt.Println(d, e)
	//fmt.Println(g)

	//select {}

	//a := float64(10)
	//fmt.Println(a.(int))
	//
	//var b sync.WaitGroup
	//b.Add()

	//fmt.Println('a' - 'b')

	//fmt.Println("abc"[1:2])
	//fmt.Println(string(100))
	//fmt.Println(strings.Count(""))

	//a := []string{"test", "123", "abc"}
	//fmt.Println(cap(a))
	////a = a[1:]
	////fmt.Println(cap(a))
	//
	//b := make([]string, 0)
	//fmt.Println(cap(b))
	//
	//c := []string{}
	//fmt.Println(cap(c))
	//
	//d := a[:0]
	//fmt.Println(cap(d), len(d))
	//
	//d = append(d, "uuu")
	//fmt.Println(d, cap(d), len(d))
	//fmt.Println(a, cap(a), len(a))
	//
	//d = append(d, "qqq")
	//fmt.Println(d, cap(d), len(d))
	//fmt.Println(a, cap(a), len(a))

	//a := int(12345) % (1e9 + 7)
	//fmt.Println(a)

	//sort.Search()
	fmt.Println('9' - '0')
}

func FindPhoneNumber(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return regexp.MustCompile("[0-9]+").Find(b)
}

func bind(configMap map[string]string, result interface{}) error {

	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("qiudianzan")
		fmt.Println(tag)
	}
	return nil
}
