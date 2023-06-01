package main

import (
	"log"
	"reflect"
	"strings"
	"sync"
)

func main() {
	log.Println("start")
	wg := new(sync.WaitGroup)
	//var wg sync.WaitGroup
	//wg的方法都是指针方法，所以直接用变量访问不到，用指针变量可以
	typ := reflect.TypeOf(wg)
	log.Printf("%+v", typ)
	log.Print(typ.NumMethod())
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		args := make([]string, 0, method.Type.NumIn())
		replys := make([]string, 0, method.Type.NumOut())
		for j := 1; j < cap(args); j++ { //方法method的In(0)是self
			args = append(args, method.Type.In(j).String())
		}
		for j := 0; j < cap(replys); j++ {
			replys = append(replys, method.Type.Out(j).String())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			//typ.Name(),//打印不出来，因为typ是指针，要取elem
			method.Name,
			strings.Join(args, ","),
			strings.Join(replys, ","))
	}
}
