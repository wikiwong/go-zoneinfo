# go-zoneinfo
A lil lib which reads the local OS [zoneinfo](https://www.iana.org/time-zones) db.

# why?
[Go's `Location` type](https://github.com/golang/go/blob/master/src/time/zoneinfo.go#L18-L35) in the `time` package does not contain any public fields exposing a timezone's offset from UTC. This is a slightly modified subset of Go's `time` package which makes the interesting properties of `Location` public. 

I am simply printing out values, you may do what you please with them :)
