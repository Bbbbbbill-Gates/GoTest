const (
    burst   = 100
    rate    = 100
    seconds = 5
)

store := redis.NewRedis("localhost:6379", "node", "")
fmt.Println(store.Ping())
// New tokenLimiter
limiter := limit.NewTokenLimiter(rate, burst, store, "rate-test")
timer := time.NewTimer(time.Second * seconds)
quit := make(chan struct{})
defer timer.Stop()
go func() {
  <-timer.C
  close(quit)
}()

var allowed, denied int32
var wait sync.WaitGroup
for i := 0; i < runtime.NumCPU(); i++ {
  wait.Add(1)
  go func() {
    for {
      select {
        case <-quit:
          wait.Done()
          return
        default:
          if limiter.Allow() {
            atomic.AddInt32(&allowed, 1)
          } else {
            atomic.AddInt32(&denied, 1)
          }
      }
    }
  }()
}

wait.Wait()
fmt.Printf("allowed: %d, denied: %d, qps: %d\n", allowed, denied, (allowed+denied)/seconds)
