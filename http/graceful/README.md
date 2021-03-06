# Graceful HTTP Shutdown

Graceful HTTP Shutdown was introduced into Golang in 1.8 via [Server.Shutdown](https://golang.org/pkg/net/http/#Server.Shutdown).

This example that I have here is currently the best I've come across without triggering an error for exiting too soon. 

Ideally I'd prefer to have the [Server.ListenAndServe](https://golang.org/pkg/net/http/#Server.ListenAndServe) as the mainblocking call in main but I've struggled to get that working without the odd error when killing the process. I don't know if that odd error could ever cause problems but I thought for now I'll just put the method that I know seems to work flawlessy. 

Updated this example based on some code [campoy](https://github.com/campoy) did in a [justforfunc episode](https://www.youtube.com/watch?v=SWKuYLqouIY&list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ&index=3) reviewing the [ursho](https://github.com/douglasmakey/ursho) project. 

## References

1. https://golang.org/pkg/net/http