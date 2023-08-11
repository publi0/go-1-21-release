package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"
)

func main() {
	//novas funcçoes no cotext
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	context.AfterFunc(ctx, func() {
		fmt.Println("Executa após cancelamento")
	})

	time.Sleep(2 * time.Second)

	// Novo erro padrao
	err := DestroyEverything()

	if err != nil {
		if errors.Is(err, errors.ErrUnsupported) {
			fmt.Println("operation is not supported")
		} else {
			fmt.Println("An error occurred:", err)
		}
	}

	//Math novas conversoes

	bigInt := big.NewInt(0)
	bigInt.SetString("1234567890123456789012345678901234567890", 10) // 40 digits

	// Convert bigInt to a float64
	floatVal, accuracy := bigInt.Float64()

	fmt.Printf("Float64 value: %f\n", floatVal) //Float64 value: 1234567890123456846996462118072609669120.000000
	fmt.Printf("Accuracy: %v\n", accuracy)      // Accuracy: Above
	//  Below Accuracy = -1
	//  Exact Accuracy = 0
	//  Above Accuracy = +1

	//	Sync Package
	f := func() {
		fmt.Println("Executing function...")
	}

	onceFunc := sync.OnceFunc(f)

	go onceFunc()
	go onceFunc()

	// simulate waiting for goroutines to finish
	time.Sleep(1 * time.Second)
}

func DestroyEverything() error {
	return errors.ErrUnsupported
}
