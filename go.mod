module go-pi-monitor

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4 // indirect
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190213061140-3a22650c66bd
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
)
