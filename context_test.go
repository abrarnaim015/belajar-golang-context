package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T)  {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWitfhValue(t *testing.T)  {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println("contextA :", contextA)

	fmt.Println("contextB :", contextB)
	fmt.Println("contextC :", contextC)

	fmt.Println("contextD :", contextD)
	fmt.Println("contextE :", contextE)

	fmt.Println("contextF :", contextF)
	fmt.Println("contextF :", contextF.Value("f"))

	fmt.Println("contextA :", contextD.Value("b")) // data dari contextA saya ambil di contextD
}

func CreateCounter(ctx context.Context, wg *sync.WaitGroup) chan int  {
	destination := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(destination)
		counter := 1
		for {
			select {
			case <- ctx.Done():
				return
			default:
				destination <- counter
				counter++

				time.Sleep(1 * time.Second) // memperlama proses untuk pengujaon WithTimeout()
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T)  {
	fmt.Println(runtime.NumGoroutine())

	wg := sync.WaitGroup {}
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx, &wg)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cancel()
	wg.Wait()

	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T)  {
	fmt.Println(runtime.NumGoroutine())

	wg := sync.WaitGroup {}
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
	defer cancel()

	destination := CreateCounter(ctx, &wg)
	for n := range destination {
		fmt.Println("Counter", n)
	}

	wg.Wait()

	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T)  {
	fmt.Println(runtime.NumGoroutine())

	wg := sync.WaitGroup {}
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5 * time.Second))
	defer cancel()

	destination := CreateCounter(ctx, &wg)
	for n := range destination {
		fmt.Println("Counter", n)
	}

	wg.Wait()

	fmt.Println(runtime.NumGoroutine())
}