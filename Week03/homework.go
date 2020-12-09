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

func myServer(done chan int) error{
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
		case <-done:
			s.Shutdown(context.Background())
			fmt.Println("http server shutdown completed！")
		}
	}()

	return s.ListenAndServe()
}


func main() {
	sig := make(chan os.Signal, 1)
	done := make(chan int, 1)

	var g errgroup.Group
	g.Go(func() error {
		return myServer(done)
	})

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("waiting for signal")

	g.Go(func() error{
		s := <- sig
		fmt.Printf("accept signal: %s\n", s)
		done <- 0
		return nil
	})

	if err := g.Wait(); err != nil {
		done <- 0
		sig <- syscall.SIGTERM
	}

	<-done
	fmt.Println("system exiting!")
	//time.Sleep(time.Millisecond)
}
