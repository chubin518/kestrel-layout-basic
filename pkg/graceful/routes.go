package graceful

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/chubin518/kestrel-layout-basic/buildinfo"
	"github.com/chubin518/kestrel-layout-basic/pkg/result"
	timeutils "github.com/chubin518/kestrel-layout-basic/pkg/utils/time-utils"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type RoutesManager struct {
	routes []*Routes
}

func NewRoutesManager() *RoutesManager {
	routes := make([]*Routes, 0)
	routes = append(routes, DefaultRoutes())
	return &RoutesManager{
		routes: routes,
	}
}

type Routes struct {
	ConfigRoutes func(*gin.Engine) error
}

func DefaultRoutes() *Routes {
	return &Routes{
		ConfigRoutes: func(server *gin.Engine) error {

			server.GET("/health", healthHandler)
			server.GET("/metrics", metricsHandler)

			server.NoMethod(func(ctx *gin.Context) {
				result.METHOD_NOT_ALLOWED.JSON(ctx)
			})

			server.NoRoute(func(ctx *gin.Context) {
				result.NOT_FOUND.JSON(ctx)
			})
			return nil
		},
	}
}

func (r *RoutesManager) Append(routes ...*Routes) {
	r.routes = append(r.routes, routes...)
}

func (r *RoutesManager) ConfigRoutes(engine *gin.Engine, ctx context.Context) error {
	for _, route := range r.routes {
		if err := ctx.Err(); err != nil {
			return err
		}
		if route.ConfigRoutes == nil {
			continue
		}
		if err := route.ConfigRoutes(engine); err != nil {
			return err
		}
	}
	return nil
}

func healthHandler(ctx *gin.Context) {
	result.OK.WithData(map[string]any{
		"app": map[string]any{
			"name":    buildinfo.Name,
			"env":     buildinfo.Environment,
			"version": buildinfo.Version,
		},
		"build": map[string]any{
			"time":    buildinfo.BuildTime,
			"version": buildinfo.BuildVersion,
		},
		"runtime": map[string]any{
			"version": fmt.Sprintf("go version %s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
		},
	}).JSON(ctx)
}

func metricsHandler(ctx *gin.Context) {
	// 获取CPU信息
	cpuInfo, _ := cpu.PercentWithContext(ctx, time.Second, false)
	// 获取内存信息
	memInfo, _ := mem.VirtualMemory()
	// 获取磁盘信息（根目录）
	diskInfo, _ := disk.Usage("/")
	// 进程信息
	procInfo, _ := process.NewProcessWithContext(ctx, int32(os.Getpid()))
	pcpu, _ := procInfo.CPUPercent()
	pmem, _ := procInfo.MemoryPercent()
	createTime, _ := procInfo.CreateTimeWithContext(ctx)
	cmd, _ := procInfo.Cmdline()
	result.OK.WithData(map[string]any{
		"server": map[string]any{
			"cpu":  cpuInfo[0],
			"mem":  memInfo.UsedPercent,
			"disk": diskInfo.UsedPercent,
		},
		"process": map[string]any{
			"pid":  procInfo.Pid,
			"cpu":  pcpu,
			"mem":  pmem,
			"cmd":  cmd,
			"time": timeutils.NORM_DATETIME.String(time.UnixMilli(createTime)),
		},
	}).JSON(ctx)
}
