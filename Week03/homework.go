package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func myServer(sig chan os.Signal) error {
	h := http.NewServeMux()
	h.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello, bro. Welcome to my server!")
	})

	s := http.Server{
		Addr: ":8800",
		Handler: h,
	}

	go func() {
		select {
		case <-sig:
			s.Shutdown(context.Background())
			fmt.Println("http server shutdown completed！")
		}
	}()

	return s.ListenAndServe()
}


func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("waiting for signal")

	var g errgroup.Group
	g.Go(func() error {
		err := myServer(sig)
		if err != nil {
			fmt.Println(err)
			sig <- syscall.SIGTERM
		}
		return err
	})

	if err := g.Wait(); err != nil {
		fmt.Println("system exiting!")
	}

}
