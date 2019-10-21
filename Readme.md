Small Caching middleware utility for AutoPilot Contact API

This contact middleWare can either be used on the client side or the server side.


All implementation share the same testing functions for consistency.

To use on your machine, you will need to populate the two constants in the `constants.go` file.


**How to:*** 

```bash
// Setup the Api Key in constants.go 

// cd server

go get
go run server.go

```



**TODO:** 

- Logging
- Better error handling in the MiddleWare delete
- A bit more coverage - 80% would be nice (Currently 74%)
- Benchmarks


