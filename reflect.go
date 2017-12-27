package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
	LifeExpectance int
}

func main()  {
	sparrow := &Bird{"Sparrow", 3}

	//fmt.Println(reflect.ValueOf(sparrow))
	//fmt.Println(reflect.ValueOf(sparrow).Elem())

	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()

	//fmt.Println(typeOfT)

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
