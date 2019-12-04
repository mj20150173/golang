package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"


)


func defaultHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	//get 파라미터 출력
	fmt.Println("default : ", r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("parm : ", r.Form["test_parm"])
	//parameter 전체 출력
	for k,v :=range r.Form {
		fmt.Println("key: ",k)
		fmt.Println("val: ", strings.Join(v,""))
	}
	fmt.Fprintf(w, "mj's server working")
}



func main() {
//기본 url 핸들러 메소드 지정
http.HandleFunc("/",defaultHandler)
//server start
err:=http.ListenAndServe(":9090",nil)
//handle error
if err!=nil{
	log.Fatal("ListenAndServe: ", err)
}else {
	fmt.Println("ListenAndServe started! port 9090")
}
























/*
	buffer := func() chan int {
		ch := make(chan int, 100)

		go func() {
			for i := 0; i < 100; i++ {
				ch <- i

			}

		}()

		return ch
	}

	consumer := buffer()
	producer := buffer()

	for i := 0; i < 200; i++ {
		select {
		case n := <-consumer:
			fmt.Println("Consumer: ", n)
			time.Sleep(50 * time.Millisecond)

		case n := <-producer:
			fmt.Println("Producer: ", n)
			time.Sleep(50 * time.Millisecond)
		}
	}
	close(consumer)
	close(producer)
*/
}
