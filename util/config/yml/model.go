package yml

type Config struct {
	Http     Http     `yml:"http"`
	DataBase DataBase `yml:"database"`
}

type DataBase struct {
	Type     string `yml:"type"`
	Host     string `yml:"host"`
	Port     string `yml:"port"`
	Instance string `yml:"instance"`
	Username string `yml:"username"`
	Password string `yml:"password"`
}

type Http struct {
	Addr    string  `yml:"addr"`
	Context Context `yml:"context"`
}

type Context struct {
	Timeout int `yml:"timeout"`
}
