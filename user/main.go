package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"

	"github.com/panjf2000/ants/v2"
)

func doGet(i interface{}) {
	rsp, err := http.Get("http://localhost:8081/v1/seckill?name=iphone")
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
	runTimes := 20
	log.Println("cpu:", runtime.NumCPU())
	p, _ := ants.NewPoolWithFunc(runtime.NumCPU(), func(i interface{}) {
		doGet(i)
		wg.Done()
	})
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	defer p.Release()
}
