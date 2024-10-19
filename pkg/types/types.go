package types

type ServerInfoConf struct {
	ServiceId   string `json:",omitempty"`
	ServiceName string `json:",omitempty"`
}

type SqlConf struct {
	Separation  int8            `json:",omitempty"`
	MasterSlave MasterSlaveConf `json:",omitempty"`
	SqlSource   SqlSourceConf   `json:",omitempty"`
}
type MasterSlaveConf struct {
	MasterAddr string        `json:",omitempty"`
	SlaveAddr  SlaveAddrConf `json:",omitempty"`
}
type SlaveAddrConf struct {
	Tag     []string `json:",omitempty"`
	Connect []string `json:",omitempty"`
}
type SqlSourceConf struct {
	Addr string `json:",omitempty"`
}

type BizRedisConf struct {
	Addr     string `json:",omitempty"`
	Password string `json:",omitempty"`
	DB       int    `json:",omitempty"`
}
type CacheRedisConf struct {
	Addr     string `json:",omitempty"`
	Password string `json:",omitempty"`
	DB       int    `json:",omitempty"`
}

type SocketOptionsConf struct {
	// BucketNum cpu个数
	BucketNum uint
	// MaxMessageSize 消息最大字节
	MaxMessageSize int64
	// PingPeriod 每次ping的间隔时长
	PingPeriod int
	// PongPeriod 每次pong的间隔时长，可以是PingPeriod的一倍|两倍
	PongPeriod int
	// WriteWait client的写入等待超时时长
	WriteWait int
	// ReadWait client的读取等待超时时长
	ReadWait int
}
