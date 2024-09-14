package pkg

import (
	"log/slog"
	"os"
	"github.com/lmittmann/tint"
	"strconv"
	"runtime"
	"syscall"
)

func Log(json bool,debug bool,color bool) {
	if json {
		if debug {
			handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
			slog.SetDefault(slog.New(handler))
		} else {
			handler := slog.NewJSONHandler(os.Stderr, nil)
			slog.SetDefault(slog.New(handler))
		}
	} else {
		if debug {
			if color {
				handler := tint.NewHandler(os.Stderr, &tint.Options{Level:slog.LevelDebug})
				slog.SetDefault(slog.New(handler))
			} else {
				handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
				slog.SetDefault(slog.New(handler))
			}
		} else {
			if color {
				handler := tint.NewHandler(os.Stderr, nil)
				slog.SetDefault(slog.New(handler))
			} else {
				handler := slog.NewTextHandler(os.Stderr, nil)
				slog.SetDefault(slog.New(handler))
			}
		}
	}
}

func CPU() string{
	return strconv.Itoa(runtime.NumCPU())
}

func Memory() string {
	var info syscall.Sysinfo_t
	err := syscall.Sysinfo(&info)
	if err != nil {
		slog.Error("Error getting sysinfo: " + err.Error())
		os.Exit(1)
	}

	totalMemory := uint64(info.Totalram) * uint64(info.Unit)
	return strconv.FormatUint(totalMemory/1024/1024,10)
}
