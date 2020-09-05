package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"

	"github.com/panjf2000/ants/v2"
)

var(
	requestURL = "http://localhost:8081/v1/seckill?name=iphone"
	users = 20
)


func doGet(i interface{}) {
	rsp, err := http.Get(requestURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	res, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("user", i, ":", string(res))
}

func main() {
	var wg sync.WaitGroup
	log.Println("cpu:", runtime.NumCPU())
	p, _ := ants.NewPoolWithFunc(runtime.NumCPU(), func(i interface{}) {
		doGet(i)
		wg.Done()
	})
	for i := 0; i < users; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	defer p.Release()
}
