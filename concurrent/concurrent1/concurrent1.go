package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

// error, 没有等待协程结束
//func makeThumbnails2(filenames []string) {
//	for _, f := range filenames {
//		go thumbnail.ImageFile(f) // NOTE: ignoring errors
//	}
//}

func testConcurrent1(fileNames []string) {
	ch := make(chan struct{})
	for _, f := range fileNames {
		go func(f string) {
			// do something
			fmt.Printf("file: %s\n", f)
			ch <- struct{}{}
		}(f)
	}
	for range fileNames {
		fmt.Println("range")
		<-ch
	}
}

func testConcurrent2(fileNames []string) error {
	ch := make(chan error)

	for _, f := range fileNames {
		go func(f string) {
			fmt.Printf("file: %s\n", f)
			err := errors.New(f)
			ch <- err
		}(f)
	}
	for range fileNames {
		if err := <-ch; err != nil {
			fmt.Printf("err: %+v\n", err)
			return err // NOTE: incorrect: goroutine leak!
		}
	}
	return nil
}

func testConcurrent3(fileNames []string) (thumbFiles []string, err error) {
	type item struct {
		fileName string
		err      error
	}
	ch := make(chan item, len(fileNames))
	for _, f := range fileNames {
		go func(f string) {
			var it item
			it.fileName, it.err = f, errors.New(f)
			ch <- it
		}(f)
	}
	// index := 0
	for range fileNames {
		// index += 1
		it := <-ch
		fmt.Printf("item: %+v\n", it)
		// if it.err != nil && index == 4 {
		if it.err != nil {
			return nil, it.err
		}
		thumbFiles = append(thumbFiles, it.fileName)
	}
	return thumbFiles, nil
}

func testConcurrent4(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of working goroutines
	for f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()
			err := errors.New(f)
			if err != nil {
				log.Println(err)
				return
			}
			sizes <- int64(10)
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total

}

func main() {
	fileNames := []string{"dir1", "dir2", "dir3"}
	// testConcurrent2(fileNames)

	resultFiles, err := testConcurrent3(fileNames)
	fmt.Printf("resultFiles: %+v, error: %+v\n", resultFiles, err)

	//ch := make(chan string)
	//for _, f := range fileNames {
	//	go func(f string) {
	//		ch <- f
	//	}(f)
	//}
	//count := testConcurrent4(ch)
	//fmt.Printf("count: %+v", count)
	fmt.Println("end")
}
