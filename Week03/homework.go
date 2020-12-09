package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func myServer(ctx context.Context, addr string) error {
	h := http.NewServeMux()
	h.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello, bro. Welcome to my server[%s]!", addr)
	})

	s := http.Server{
		Addr: addr,
		Handler: h,
	}

	go func() {
		select {
		case <-ctx.Done():
			if err := s.Shutdown(context.Background()); err != nil {
				ctx.Done()
			}
			fmt.Println("shutdown http server complete!")
		}
	}()

	return s.ListenAndServe()
}

func mySignal(ctx context.Context, sig chan os.Signal) error{
	select{
	case <-sig:
		ctx.Done()
		time.Sleep(time.Nanosecond)		// 等待shutdown完成的打印
		return errors.New("signal exiting! ")
	case <- ctx.Done():
		sig <- syscall.SIGTERM
		return nil
	}
}


func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("waiting for signal")

	g, cancel := errgroup.WithContext(context.Background())
	g.Go(func() error {
		err := myServer(cancel, ":8800")
		if err != nil {
			cancel.Done()
		}
		return err
	})

	g.Go(func() error {
		return mySignal(cancel, sig)
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}