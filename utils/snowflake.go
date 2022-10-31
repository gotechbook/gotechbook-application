package utils

import (
	"strconv"

	"github.com/sony/sonyflake"
)

var (
	SFlake    *SnowFlake
	MachineID uint16
)

// SnowFlake SnowFlake算法结构体
type SnowFlake struct {
	sFlake *sonyflake.Sonyflake
}

func init() {
	SFlake = NewSnowFlake()
}

func getMachineID() (mID uint16, err error) {
	mID = MachineID
	return
}

func NewSnowFlake() *SnowFlake {
	st := sonyflake.Settings{}
	// machineID是个回调函数
	st.MachineID = getMachineID
	return &SnowFlake{
		sFlake: sonyflake.NewSonyflake(st),
	}
}

func (s *SnowFlake) GetID() (uint64, error) {
	return s.sFlake.NextID()
}

func (s *SnowFlake) GenerateID() string {
	id, _ := s.sFlake.NextID()
	return strconv.FormatInt(int64(id), 10)
}
