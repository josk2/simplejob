package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	// nhận vào context parent (Background) và trả về context child (ctx) và hàm cancel
	// deadline 10 secs
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup

	// đơn giản hoá worker pool cho ngắn gọn
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	time.Sleep(time.Second)

	// mặc dù ctx sẽ expire theo timeout đã set trước đó
	// ta vẫn gọi cancel để đóng context child và các children của nó
	// để tránh giữ chúng tồn tại không cần thiết
	cancel()

	// sử dụng waitGroup thay cho done channel
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("Start: ", start)
	fmt.Println("Elapsed: ", elapsed)
}

func worker(ctx context.Context, wg *sync.WaitGroup, num int) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello ", num)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
