package yml

type Config struct {
	DataBase DataBase `yml:"database"`
}

type DataBase struct {
	Type     string `yml:"type"`
	Ip       string `yml:"ip"`
	Port     string `yml:"port"`
	Instance string `yml:"instance"`
	Username string `yml:"username"`
	Password string `yml:"password"`
}
