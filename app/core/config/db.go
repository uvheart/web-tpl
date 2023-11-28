package config

type DBItemConf struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Database     string `yaml:"database"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Charset      string `yaml:"charset"`
	Timeout      int    `yaml:"timeout"`
	WriteTimeOut int    `yaml:"write_time_out"`
	ReadTimeout  int    `yaml:"read_time_out"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max-open-conns"`
}

type DBLog struct {
	Enable bool   `yaml:"enable"`
	Level  string `yaml:"level"`
	Path   string `yaml:"path"`
	Type   string `yaml:"type"`
	Format string `yaml:"format"`
}

type dbItem struct {
	Write DBItemConf `yaml:"write"`
	Read  DBItemConf `yaml:"read"`
}
