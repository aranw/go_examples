# Graceful HTTP Shutdown

Graceful HTTP Shutdown was introduced into Golang in 1.8 via [Server.Shutdown](https://golang.org/pkg/net/http/#Server.Shutdown).

This example that I have here is currently the best I've come across without triggering an error for exiting too soon. 

Ideally I'd prefer to have the [Server.ListenAndServe](https://golang.org/pkg/net/http/#Server.ListenAndServe) as the mainblocking call in main but I've struggled to get that working without the odd error when killing the process. I don't know if that odd error could ever cause problems but I thought for now I'll just put the method that I know seems to work flawlessy. 

Would love to see alternatives and suggestion for this problem as I know pretty much most if not all apps I implement in Go would use a web server in some way and I'd use a graceful shutdown technique like this. 

## References

1. https://golang.org/pkg/net/http