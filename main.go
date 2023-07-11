package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

//go:embed web-app/dist/*
var templateFS embed.FS

type ProcessInfo struct {
	Pid     int32  `json:"pid"`
	Name    string `json:"name"`
	Cmdline string `json:"cmdline"`
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	front, err := fs.Sub(templateFS, "web-app/dist")
	if err != nil {
		fmt.Println(err)
	}

	r.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/" {
			c.Redirect(http.StatusMovedPermanently, "/ui")
			c.Abort()
			return
		}
	})

	r.StaticFS("/ui", http.FS(front))

	r.GET("/info", func(c *gin.Context) {
		data := make(map[string]interface{})

		// 获取内存信息
		memInfo, err := mem.VirtualMemory()
		if err != nil {
			c.JSON(500, gin.H{"error": "无法获取内存信息"})
			return
		}
		data["mem"] = memInfo

		// 获取CPU信息
		cpuInfo, err := cpu.Info()
		if err != nil {
			c.JSON(500, gin.H{"error": "无法获取CPU信息"})
			return
		}
		data["cpu"] = cpuInfo

		// 获取磁盘信息
		diskInfo, err := disk.Usage("/")
		if err != nil {
			c.JSON(500, gin.H{"error": "无法获取磁盘信息"})
			return
		}
		data["disk"] = diskInfo

		// 获取主机信息
		hostInfo, err := host.Info()
		if err != nil {
			c.JSON(500, gin.H{"error": "无法获取主机信息"})
			return
		}
		data["host"] = hostInfo

		// 获取负载信息
		loadInfo, err := load.Avg()
		if err != nil {
			c.JSON(500, gin.H{"error": "无法获取负载信息"})
			return
		}
		data["load"] = loadInfo

		// 获取网络信息
		netInfo, err := net.IOCounters(true)
		if err != nil {
			c.JSON(500, gin.H{"error": "无法获取网络信息"})
			return
		}
		data["net"] = netInfo

		c.JSON(200, data)
	})

	r.GET("/processes", func(c *gin.Context) {
		sort := c.DefaultQuery("sort", "mem") // 默认按内存排序
		top := c.DefaultQuery("top", "20")    // 默认返回前20个进程

		data := make(map[string]interface{})

		// 获取进程信息
		processList, err := process.Processes()
		if err != nil {
			c.JSON(500, gin.H{"error": "无法获取进程信息"})
			return
		}

		var processInfoList []ProcessInfo

		switch sort {
		case "cpu":
			processInfoList = sortProcessesByCPU(processList)
		default:
			processInfoList = sortProcessesByMemory(processList)
		}

		if len(processInfoList) > 0 {
			topCount, err := strconv.Atoi(top)
			if err != nil || topCount <= 0 {
				topCount = len(processInfoList)
			} else if topCount > len(processInfoList) {
				topCount = len(processInfoList)
			}

			processInfoList = processInfoList[:topCount]
		}

		data["process"] = processInfoList

		c.JSON(200, data)
	})

	err = r.Run(":8656")
	if err != nil {
		fmt.Println("启动服务器失败:", err)
	}
}

// 根据内存占用排序进程
func sortProcessesByMemory(processList []*process.Process) []ProcessInfo {
	sort.Slice(processList, func(i, j int) bool {
		mem1, _ := processList[i].MemoryInfo()
		mem2, _ := processList[j].MemoryInfo()
		return mem1.RSS > mem2.RSS
	})

	return getProcessInfoList(processList)
}

// 根据 CPU 占用排序进程
func sortProcessesByCPU(processList []*process.Process) []ProcessInfo {
	sort.Slice(processList, func(i, j int) bool {
		cpuPercent1, _ := processList[i].CPUPercent()
		cpuPercent2, _ := processList[j].CPUPercent()
		return cpuPercent1 > cpuPercent2
	})

	return getProcessInfoList(processList)
}

// 获取进程信息列表
func getProcessInfoList(processList []*process.Process) []ProcessInfo {
	processInfoList := make([]ProcessInfo, 0, len(processList))

	for _, p := range processList {
		name, _ := p.Name()
		cmdline, _ := p.Cmdline()

		processInfoList = append(processInfoList, ProcessInfo{
			Pid:     p.Pid,
			Name:    name,
			Cmdline: cmdline,
		})
	}

	return processInfoList
}
