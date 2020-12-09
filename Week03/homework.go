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
)

func myServer(ctx context.Context) error {
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
		case <-ctx.Done():
			s.Shutdown(context.Background())
			fmt.Println("http server shutdown completed！")
		}
	}()

	return s.ListenAndServe()
}

func mySignal(ctx context.Context, sig chan os.Signal) error{
	select{
	case <-sig:
		ctx.Done()
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
		err := myServer(cancel)
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
