package main

import (
	"fmt"
	"github.com/ssdev-go/download/pkg/base"
	"github.com/ssdev-go/download/pkg/download"
)

func main() {
	finallyCh := make(chan error)
	err := download.Boot().
		URL("https://www.baidu.com/index.html").
		Listener(func(event *download.Event) {
			if event.Key == download.EventKeyFinally {
				finallyCh <- event.Err
			}
		}).
		Create(&base.Options{
			Connections: 8,
		})
	if err != nil {
		panic(err)
	}
	err = <-finallyCh
	if err != nil {
		fmt.Printf("download fail:%v\n", err)
	} else {
		fmt.Println("download success")
	}
}
