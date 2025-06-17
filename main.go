package main

import (
	"encoding/json"
	"funWithRealtime/hub"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	h := hub.New()
	go h.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.Serve(h, w, r)
	})

	var (
		qu []int
		mu sync.Mutex
	)

	go func() {
		for {
			mu.Lock()
			qu = append(qu, rand.Intn(9)+1)
			mu.Unlock()
			time.Sleep(200 * time.Millisecond)
		}
	}()
	go func() {
		for {
			time.Sleep(10 * time.Second)

			mu.Lock()
			sum := 0
			for _, v := range qu {
				sum += v
			}
			count := len(qu)
			qu = qu[:0]
			mu.Unlock()

			if count == 0 {
				continue
			}
			avg := float64(sum) / float64(count)

			msg, _ := json.Marshal(map[string]any{
				"avg": avg,
			})

			h.Broadcast <- msg
			log.Println("Sent avg:", avg)
		}
	}()
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
