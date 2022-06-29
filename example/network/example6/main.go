// 抓包本机网卡  解决线上偶现问题来不及抓包的情况
//https://www.codeprj.com/blog/b76a9e1.html

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	device            = "enp0s3" // 指定监控网卡名称信息 根据实际情况填写
	snapshotLen int32 = 10000000
	err         error
	timeout     = 30 * time.Second
	handle      *pcap.Handle
)

func main() {
	handle, err = pcap.OpenLive(device, snapshotLen, false, timeout)
	filter := "host www.baidu.com" // 指定抓取域名
	err = handle.SetBPFFilter(filter)
	if err != nil {
		os.Exit(-1)
	}
	defer handle.Close()
	log := initLog()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		applicationLayer := packet.ApplicationLayer()
		if applicationLayer != nil {
			log.Info(string(applicationLayer.LayerContents()))
		}
		if err := packet.ErrorLayer(); err != nil {
			fmt.Println("Error decoding some part of the packet:", err)
		}
	}
	for {
		time.Sleep(time.Second)
	}
}

func initLog() *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   "./logs/package.log", // 日志文件路径
		MaxSize:    10,                   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 5,                    // 日志文件最多保存多少个备份
		MaxAge:     7,                    // 文件最多保存多少天
		Compress:   true,                 // 是否压缩, 压缩后1M约占20Kb
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),                                        // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
	logger := zap.New(core)

	logger.Info("log 初始化成功")
	return logger
}
