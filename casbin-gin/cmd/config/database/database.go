package database

var (
	Config = new(Database)
	//DatabasesConfig = make(map[string]*Database)
)

type Database struct {
	Driver          string `yaml:"driver"`
	Source          string `yaml:"source"`
	ConnMaxIdleTime int
	ConnMaxLifeTime int
	MaxIdleConns    int
	MaxOpenConns    int
	Registers       []DBResolverConfig
}

type DBResolverConfig struct {
	Sources  []string
	Replicas []string
	Policy   string
	Tables   []string
}

func (d *Database) Setup() {
	//db, err :=
	//TODO
	//runtime.ApplicationContext.SetDb("*")
}
