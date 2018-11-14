# Status
### Package status provides helpers for GRPC codes

`go get github.com/Snow-Sight/pkg/grpc-status`

Running `status.Error(e error)` will return a GRPC style error. If `e` is a `*era.Error` and with a known era code, the code will be converted into a known GRPC code.