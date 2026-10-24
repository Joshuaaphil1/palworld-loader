package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.6" )

func b6GyaX4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3N4XCzWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rALTCV2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bn92J4oVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bthzXuTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BA6hqjvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57Q8MCLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3gCdqCX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCO4aNptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQfzU4kGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3naJGJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIRRoqMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBWAZdgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpI9nIaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DU7WnNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbcgbNvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76auIbk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4VnLuCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gmy4VfCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVHNXCVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0lhYeiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xXNpqapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shC5D9OLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgTyMTWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlcggZnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QFKIl9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKAIJMDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1X89wTe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qu8LTytSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9XYRw6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxhtWTxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0i7WUqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhQANIosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8cF1vxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZw413T2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwGKEapbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bcuKHvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sk5rrVDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5nshy1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2cw6HceyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgchghB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDwblt80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FFCrE1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aluxKRb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yv6vLglNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhGRfRYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HvJzfYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xr2aoYAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HapYErdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ailwj8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZeK40jNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfkxfpluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsg6ZCToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rO5GltYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1v1Kl8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sh8kHHr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjkMzVypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTx3kjbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRosmKlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTQqHb7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVTfC6WKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ww2ZiPhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgKS9JmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uHjPDNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9OkYbWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzJOpvJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIVMCGYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNsm8UfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORKhcwc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qQOIlKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFpQ0QpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8F031lEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqiOasLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEg8zLn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7vLm6D6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HrJ6fs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gj089QlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpUEjsGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KfGnuE2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TF5V10QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmybNvGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHV21lzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fqxRtLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Flj6zxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCFlpvXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYm6ObYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBQ62B53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0B6mj7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cguR7MBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JyxvJsbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXoYu9lKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1K7nfwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wp9RfjdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NO58QoziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbwFQJIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oiWk0jBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMWVTBiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1tKcKhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tym0M88CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwOQmZ52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKFpytwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2fcPNDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZX5jfRvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2A3VhNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTuqY9ZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNvMU9bpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIkVh14zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jmu2IdMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PC3igawXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jub2EYrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76IzM9iSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akZqUO7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBDzVmyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnGi3PihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KC2Uv54EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVVbdeiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cYDYiF2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOCI4hR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtBTmPrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53KHmjMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAfnOdISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPNV6JLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2kbJwxEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJM4NnYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJy9YZr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJUIFJHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0JiL6OTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmOYcvlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYpZQELDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWwf7wURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThuLFb0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRDbOTP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjhtdN60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJ2NqGl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqp6vq74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGgH915vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJwIEBysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsEPjZT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTw5IRLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3q1bPyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCUCjhC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98s3rlugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lL88c6IwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HW2YFOM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkfaJGlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WU92mN3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhLYwS4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVAGhrYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XARVlt9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHroZtGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrVAVkYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3WJxKpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbAu5IHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGhCVbUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5K9IV5pqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKoLv9spWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grFgQMTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzGH00MKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0svf0fsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYUhQkbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U44ByZJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IInyTXUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJC8V5x0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPdEbQRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YuSKsmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xywZ9DeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4d7TwZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrJxqRVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZO6uTXo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeTP7uTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gE3Y0Q2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6AJGfP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHg3PtGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fo8S0lfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBXcneZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hc1sK8HQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95ZeSzhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Qv6dtqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65olFpRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqxfHY80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrqmfShFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQzALYIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJ9uIweoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDcWf101Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuROTq74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0m5IzFd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8xY5pGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2b6OYiiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dB7bVhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYVBUG9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDmFQTKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnU4h90HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjMG6KlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SRpuTZLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTmilDRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OxlZIp96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8IBs6H9TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BktW8QCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Phnw6VTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXMXUjLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reMXtgbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czdAhRsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qa0UaRguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtZ7Z7aHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FI5Dap89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQI014EzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TB4VpT4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISwCRGIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hS5Yg2uhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDAqhduRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMHkfSptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVyindAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmpPLMgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfOJwAgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F26Uj3C0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BleXV1S3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6doxE0uxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2BvTMidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bW6WUbjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3D0D8UJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBt3Ne4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWbe8Yz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHNqCLkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwGJxumRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCIxU7TWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oqie5ndnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmYHWGUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YwALbhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ncl8GLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtl1FQvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYOXKOTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uV4bTUFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foQMw3IWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lltg7WB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUpU2FU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9YG1gvx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6cUyA6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZ6DxwT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPko5gcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgTQCpFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVa4ZqiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZF8WMCuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NM8H80JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1izLbGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRFbcU38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXi1wLStWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2f4MUnixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNpTV5LXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5lpxaLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcUZSs0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsQ4buPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8NAckWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqulfDkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PVvzB7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Gn61y0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HV0m9X4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIkKSC5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7pnYYL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbO7XarpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFgN6hgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzJbOeWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNl7R8OBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func po4bz3Y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syhy7MzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSsN3wpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zs7RX5xfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtrgdGzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYnHyMghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEe89qQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5slbJAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIUYbK3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYby7DF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7wLwiFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbMVPlO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSlOfAZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rq0o81CPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgzsDvriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujZt2dd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGSdeLrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciLQ37U9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIkXW0FQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COqQyX7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o6fwMA8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7NHvSmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8gkyRHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJFmDP6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6foKsubAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbewMVCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x06AtwDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZjM5XJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVFHarh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPgF1l78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0f2h9QzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZoMPYQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZuQHzCMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1X7B9cXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDTheQluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 97CY4Y2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jG1oNElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Rez3nAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvkfA5JFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqckD9ZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGQre09zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvrdFNiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujR5e0BwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTz7CVPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpPdlTpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3idDqvsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6vdolIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPcGDguWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3kFmAWUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTtRZXg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTAa4af9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxyc1E2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EdR5PH7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3WqD9HUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRArhWfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sd2STJohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGlyT4EMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjjMp4jQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfYEZTLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUj5o530Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9LJcO6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dszUPyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuTDR1laWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXfVZRMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdjQiCVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCpue8kRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxo1i2ryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EANDjIvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9Y7jKYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2FUtbJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpKf1ZGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDH76qnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsSJ0TCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7t8AmP4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EETIFbP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uh0s1hRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUQyxj0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8QmxiuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wU8t9G0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8td13tjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pPLMQxRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAbsAOnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99aU30rHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ocy7QeFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKs7b8PMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtuFkcReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqmjyYwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXeHf74jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0b4fFykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdUEJc1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxO6BR1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzFafej7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5vL7HG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZWJUQhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lz6TOxVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVZ74aHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DgvHL7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oo534XnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKzmkHVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybHyhfuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSsBeggjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJL7kKZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlGNG3IrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0YuSLlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPsGSwoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2xa0nNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bynlRCOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0mkQ05AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQgEh3qgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHb4n7WxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5iR07gXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiFl5VU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 856zFKVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKF9dHnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3igyNA0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjNwsoa2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qA9hVzI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrXa8yiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2KviLFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeeohCOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQ8B1EAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqqz7gVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmyWmlnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gi6tJ8diWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tke05E2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CA9WJyt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5y4woakxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8RN9h0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIOJRcclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkuNo5jLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pX7qFByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoN12ujdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58Fmb9RlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0OytLD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6XZJUe2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ve07EOlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRgROT1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqvDgwIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2AHeiSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dr3OC19LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svMqeVEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDeZM5gbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRlKu6wiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73WvDsyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsHquh6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iInqSmUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzqtkNgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SltyOZxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZEwL1aLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZL8bXsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2N9LKF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSwny5dOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBrteai9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1T0bC0PqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WjyANm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LdIbmOeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEuXjnTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vs6A7t0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmKftdDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrpBmbyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlFDKwXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2rOSpF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKHfASdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E89CNR9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xd331nsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELXw71bBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPvYfQhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFJWiDumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UGLhssEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHyiaFeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOMkbdrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAbpgEkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ha4dldk5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nFakDR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZgwzjVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTaisKttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lV8gWuWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsTWbLOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9ojpeQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ro8o3dL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZPGojw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haNrRpttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzyoGlm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxbwanL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DWqP8mrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cWCCoIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0TuXTCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70OUMX58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T7xVB4fcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHS9fsWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozcipYTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Us1bD2ztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFaOC6QiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkFHU2d0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIi6hFasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJOHbeOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zA4SqneDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func anSSNsGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gu5xMkYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdP9Ooo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1h99tSx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dz2goWpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5cBmbqEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 159XWhrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIcccCSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBVtPkRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIt6B5HhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcsI4RqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKQrVX6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sYDCTDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFsgzY56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkxjPzDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5LqHGhPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ananolByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tpWvdmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3wBlu6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMcs3jB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rncRmWCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXrp8lAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IByc212VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9p0U3DwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpklTEQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlcmyVE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjWbFS7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNHeKqyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWNrtOlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWN6y0AIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqXi2UYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snfZUuWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bdkPMjh3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVHS58OCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnZW23XuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7QRARo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z91YxaLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bH4C7SsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEyAKVsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwwcrcpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nc5kfdi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4txs7a3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9YlzsvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mY1OLlsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VizTrA2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNjGRJ1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HedlXv8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmuHzQ6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bm3g48aTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AO29gykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3sqPH8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paZuxSKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRTlmY06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UbQIsAAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWbonEluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leFhRapQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOmbgWsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vid2SWQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvNUPWooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tt9MrnRBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQxFG5eiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMjIBgoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjMBfiObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3RCVknkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rN7vnIlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXzvYc9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwKkmZJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RhxB0oFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8vgS5F7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mf0xwDxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXbowX3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vykjDoPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRAVnUAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YE1PjwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJX8nUkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAyniNKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQIXVSIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txDncpomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iC4WrdVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5lWyJaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vaxli9DTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxlHqEePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHRBPGW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYagzsFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvAXycSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSf60lEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ksrqse5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtvbUpRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAsENjdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GcbzHqg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eft9PqyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCciw08uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDn8cGqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cyHpEnIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xY0hnecdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ow0Q11v8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kz1qwRxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YG8TOEHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNZWCDF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZxRpPEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUUqA2ytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3duQCJZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rUNRSwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTnn9JvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sut9b4NlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kh5Am59jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Clk3RVMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGsGf7EFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2blKehQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5RvOvqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSlJHisiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0rRSkH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXJMLV5kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyQMVRvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJuDbXAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RTiLg7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTnZfNOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ap02alLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DrLzHlrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4gStnyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcICE52WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ma4SMEMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92sgqn60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdKNbx7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKuhSkXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARzIowURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87pbtm0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nKwnsDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgQfksszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B77lOQGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFcoVUEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GuJvU5zYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1emHvwiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HaPBCxOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMLfKOzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5kZKaucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGKnK1vKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20hywOBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLUJuSMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiYafXcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JnxuBMGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzMMcGZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPD8gpBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlR7yr6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWzkjoxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjhQPisgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drXE7uF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgmBPPhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 46FFi4dlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkW4cSvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rBBxlCOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c56TNj47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xhzHv5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGr6Mxx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ikvyrxYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LV1o7auwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CpfT1IJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPsFhW4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAtzVlhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuKMPbH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNLWI1gmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hjkWpn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOmtwLEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAEwrLABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nueFYCh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mH9YJGKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Js8QJ6oOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFXUIu6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7kcsDVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MxTurjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcDTPNOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyLZymvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8PPOQPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lQZ8qJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZYDGR6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0X2ddyPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFHXBT5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYgNJFkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuM3Vg2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuUER7bgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYdlrrSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAM30g3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrUElzkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCaSxx2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGPKoxt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4d5xjWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Md1PsIYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZuE0B6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWWefajcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dm2vjT0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TAXnkQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXy1lesiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKLmrd8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9mYQX9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PObMSW8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dcbWXDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8gTt3wKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhtIzUdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func logopfqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceAyLzcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1jTvDLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESQJJVy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNl3WwbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9FWjB8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDCAv0tNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdL6Hv5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKX4SN12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ud9xzB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wn7PVUkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQnLBuRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfepjH50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYInNd18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38pCOGT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PELfRtXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2AgaCjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSk9OnYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0iyxcq1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJK55l9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJB29RGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUpDxKX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Njb1fCTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oc3dHEsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dg0073DwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCGa6pJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpCp23bNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8bSxW9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVbeeP5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B9kvdrOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctPvDTkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsObvE2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEVoVSYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEEQe9ciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ht9G1CWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taEwVhiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7TVGd2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyKEG7X2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbt8DxjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bx43F4OPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWZhRXTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFih0QGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAG1njg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMcqD7K4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1iBWuL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRXASXe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExpFdN1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6UE6EXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhjfV4BeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmLrJvQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAtolSTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiclLSVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YPgu9HEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SpKis48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQaoxa2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upoHLhH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCRKHVfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQTeEgFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJyMm9GuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvrdVza4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvU7xEjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkvR6yZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrNJ7oBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfGYmdbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TUfhpyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpTLIn1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lTalDQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGLuHYc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyPMz2J1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28bQ6wztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72ttxlGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hOY83BleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLnmwSAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uJfp5cZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vCLk0L5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9ccVKy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BEpuXF2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZB4wB5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8INaiY1yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USok8hdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miuJULGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRIlAaTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s21EbdBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNbSGgbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8u4VOpjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpP3W5jfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmH8N3dOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlqYTGqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LAyEN3UQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8s1uZ5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDBywQAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UrrbpY4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QGgzKjvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQCLkZzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oonK51loWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFo74U80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Llus3Ua8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1cXMJV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3z6duqW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UnmsDxJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNGwFnSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOfK1ZiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ee3DbKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myFfoMtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwB5JKPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQMvwWstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUhRviAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7PNH897Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UO1gPE7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyqjJv1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzI3F2FSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRa13xmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KNukkE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6r1REIYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0zbwop6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZlXYZI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkGNzyDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGxHeso5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51b9oSuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yXIr0YFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kf9ZKE8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRKx2Q5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YkncCwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSLAoai5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XKy8GumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9oP533TrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2yVQRcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8jzXws7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kc4YTuiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olVEGJk5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GF3ecjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tzKd4sOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Us0mAIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pt2nbbp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reh1QWUiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9c7uro9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zV37R3EfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rid6GocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwZmkGdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2owaGVdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyW3Pb2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmFJrHONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Klx04blXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FKi0SLsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxkBWNsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qn0C56jUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAaMeMpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qzqCOOtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VL3TcNZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZWT7TsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjFOE82YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x573pAllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6G66iWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KU5ufi9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJZ0gCPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPm4IlYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxcCOxjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wl8VxaymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBh47ogzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxfi9y11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVTLYGH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOgoOeJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5OKJxG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o07SqkKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKOfSZVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Rw1rh7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvo5Cqk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQBLwbtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M50m4gAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUuKVmivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52MzfFUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCxoys4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WUKNZU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBM6x9DlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mliGbTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSj0JnShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iREaL86GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3kgXw5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lN6XnrfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxzy97rVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDzPvvQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FW4AkImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cALAC3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8e3FNnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcBlT3cYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGvTxn2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func av33HFocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhyJ23RDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func heRA7SdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6ncgmdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6bSrnxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7kxve4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTNTEAAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUgN0IoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

