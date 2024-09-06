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
	DB       int8   `json:",omitempty"`
}
type CacheRedisConf struct {
	Addr     string `json:",omitempty"`
	Password string `json:",omitempty"`
	DB       int8   `json:",omitempty"`
}
