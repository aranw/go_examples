# Worker Pools and Job Queues

Over the years I've built a few different systems that has had the need for an internal worker pool and queue. For the years I've attempted different patterns and made various alterations. I've had various problems with the various implementations I've made. Those ranged from panics, lost work and early termination issues. I wanted to come up with a more robust implementation, hoping I would overcome my past failures. 

## References

1. https://gobyexample.com/worker-pools
2. https://golangbot.com/buffered-channels-worker-pools/
3. http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
4. https://geeks.uniplaces.com/building-a-worker-pool-in-golang-1e6c0fdfd78c