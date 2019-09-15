package Facade

import "fmt"

// CPU
type CPU struct {
}

func (CPU) start() {
	fmt.Println("启动了CPU...")
}

// 内存
type Memory struct {
}

func (Memory) start() {
	fmt.Println("启动了内存...")
}

// 硬盘
type Disk struct {
}

func (Disk) start() {
	fmt.Println("启动了硬盘...")
}

// 开机按钮
type StartBtn struct {
}

func (StartBtn) start() {
	cpu := &CPU{}
	cpu.start()
	memory := &Memory{}
	memory.start()
	disk := &Disk{}
	disk.start()
}
