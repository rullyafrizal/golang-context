package src

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	ctxA := context.Background()

	ctxB := context.WithValue(ctxA, "A1", "rully")
	ctxC := context.WithValue(ctxA, "A2", "zidan")

	ctxD := context.WithValue(ctxB, "B1", "raufa")
	ctxE := context.WithValue(ctxB, "B2", "khadafie")

	ctxF := context.WithValue(ctxC, "C1", "mohammad")

	fmt.Println(ctxA)
	fmt.Println(ctxB)
	fmt.Println(ctxC)
	fmt.Println(ctxD)
	fmt.Println(ctxE)
	fmt.Println(ctxF)

	// ambil value dari context dengan menggunakan key
	fmt.Println(ctxB.Value("A1"))
	fmt.Println(ctxC.Value("A2"))

	// mengambil value dari parent context
	fmt.Println(ctxD.Value("A1"))
	fmt.Println(ctxF.Value("A2"))

	// hasil nil karena context / parent context tidak memiliki value dengan key "B1"
	fmt.Println(ctxF.Value("B1"))

	// parent tidak bisa mengambil key dari child context, hasil nil
	fmt.Println(ctxA.Value("A1"))
}

func TestContextWithCancel(t *testing.T) {
	fmt.Printf("Total Goroutine : %d\n", runtime.NumGoroutine())

	// dengan leak
	// destination := createCounter()

	// tanpa leak
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := createCounter2(ctx)

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}
	cancel()

	// berguna untuk menghentikan proses sementara
	time.Sleep(time.Second * 1)

	// jika jumlah Goroutine 3, maka terjadi leak
	fmt.Printf("Total Goroutine : %d\n", runtime.NumGoroutine())
}

// Goroutine leak
func createCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		var counter int = 1

		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

// Tanpa Goroutine leak
func createCounter2(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		var counter int = 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Printf("Total Goroutine : %d\n", runtime.NumGoroutine())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	destination := createCounterTimeout(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	fmt.Printf("Total Goroutine : %d\n", runtime.NumGoroutine())
}

func createCounterTimeout(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Printf("Total Goroutine : %d\n", runtime.NumGoroutine())

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	destination := createCounterTimeout(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	fmt.Printf("Total Goroutine : %d\n", runtime.NumGoroutine())
}

func operation1(ctx context.Context) error {
	// Let's assume that this operation failed for some reason
	// We use time.Sleep to simulate a resource intensive operation
	time.Sleep(499 * time.Millisecond)
	return errors.New("failed")
}

func operation2(ctx context.Context) {
	// We use a similar pattern to the HTTP server
	// that we saw in the earlier example
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("halted operation2")
	}
}

func TestContext2(t *testing.T) {
	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(context.Background())

	// Run two operations: one in a different go routine
	go func() {
		err := operation1(ctx)
		// If this operation returns an error
		// cancel all operations using this context
		if err != nil {
			cancel()
		}
	}()

	// Run operation2 with the same context we use for operation1
	// this below will be halted due to error in operation1
	// because of the same context
	operation2(ctx)
}
