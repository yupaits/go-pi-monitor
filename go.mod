module go-pi-monitor

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/gin-gonic/gin v1.4.0
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190213061140-3a22650c66bd
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
)
