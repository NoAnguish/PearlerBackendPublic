package config

type DatabaseConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type S3Conf struct {
	Endpoint    string `yaml:"endpoint"`
	Bucket      string `yaml:"bucket"`
	Region      string `yaml:"region"`
	Credentials struct {
		AccessKeyId     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
	} `yaml:"credentials"`
}

type Config struct {
	Server   ServerConf   `yaml:"server"`
	Database DatabaseConf `yaml:"database"`
	S3       S3Conf       `yaml:"s3"`
}
