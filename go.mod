module github.com/ljwtorch/kook-sdk-go

go 1.21

retract (
    v0.0.1 // Published accidentally
    v0.0.2 // Published accidentally
    v0.1.0-dev.1 // Pre-release version
)

require github.com/gorilla/websocket v1.5.1

require golang.org/x/net v0.17.0 // indirect
