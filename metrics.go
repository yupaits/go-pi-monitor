package go_pi_monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

type Metrics struct {
	host         *host.InfoStat
	temperatures []host.TemperatureStat
	users        []host.UserStat

	totalPercents   []float64
	percents        []float64
	infoStats       []cpu.InfoStat
	timesStats      []cpu.TimesStat
	totalTimesStats []cpu.TimesStat

	swapMemStat    *mem.SwapMemoryStat
	virtualMemStat *mem.VirtualMemoryStat

	connectionStats []net.ConnectionStat
	interfaces      []net.InterfaceStat
	netIO           []net.IOCountersStat
	netIOPerNic     []net.IOCountersStat

	parts     []disk.PartitionStat
	diskUsage *disk.UsageStat
	diskIO    map[string]disk.IOCountersStat
}

func getMetrics() gin.H {
	metrics := &Metrics{}
	//设备
	metrics.host, _ = host.Info()
	metrics.temperatures, _ = host.SensorsTemperatures()
	metrics.users, _ = host.Users()
	//CPU
	metrics.infoStats, _ = cpu.Info()
	metrics.totalTimesStats, _ = cpu.Times(false)
	metrics.timesStats, _ = cpu.Times(true)
	metrics.totalPercents, _ = cpu.Percent(500*time.Millisecond, false)
	metrics.percents, _ = cpu.Percent(500*time.Millisecond, true)
	//内存
	metrics.swapMemStat, _ = mem.SwapMemory()
	metrics.virtualMemStat, _ = mem.VirtualMemory()
	//网络
	metrics.connectionStats, _ = net.Connections("all")
	metrics.interfaces, _ = net.Interfaces()
	metrics.netIO, _ = net.IOCounters(false)
	metrics.netIOPerNic, _ = net.IOCounters(true)
	//磁盘
	metrics.parts, _ = disk.Partitions(true)
	metrics.diskUsage, _ = disk.Usage("/")
	metrics.diskIO, _ = disk.IOCounters()

	return gin.H{
		"host":              metrics.host,
		"temperatures":      metrics.temperatures,
		"users":             metrics.users,
		"cpu.info":          metrics.infoStats,
		"cpu.percent":       metrics.totalPercents,
		"cpu.percentPerCpu": metrics.percents,
		"cpu.timesPerCpu":   metrics.timesStats,
		"cpu.times":         metrics.totalTimesStats,
		"mem.swap":          metrics.swapMemStat,
		"mem.virtual":       metrics.virtualMemStat,
		"net.connections":   metrics.connectionStats,
		"net.interfaces":    metrics.interfaces,
		"net.io":            metrics.netIO,
		"net.ioPerNic":      metrics.netIOPerNic,
		"disk.parts":        metrics.parts,
		"disk.usage":        metrics.diskUsage,
		"disk.io":           metrics.diskIO,
	}
}
