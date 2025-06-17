package main

import (
	"encoding/json"
	"funWithRealtime/hub"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	h := hub.New()
	go h.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.Serve(h, w, r)
	})

	qu := make(chan int, 100)

	go func() {
		for {
			qu <- rand.Intn(9) + 1
		}
	}()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			sum := 0
			count := 0
			draining := true
			for draining {
				select {
				case v := <-qu:
					sum += v
					count++
				default:
					draining = false
				}
			}

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
