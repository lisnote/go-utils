package config

import "time"

const (
    Duration = 10 * time.Second // 测试持续时间
    Timeout  = 5 * time.Second  // 读写超时
    BufSize  = 32 * 1024        // 缓冲区大小
)
