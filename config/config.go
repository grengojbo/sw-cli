package config

// Config structure config file
// Hosts
// 		Connect -> api/ssh/sntp/telnet
type Config struct {
	// Port uint `default:"7000" env:"PORT"`
	// DB   struct {
	// 	Name     string `default:"qor_example"`
	// 	Adapter  string `default:"mysql"`
	// 	User     string
	// 	Password string
	// 	Host     string `default:"localhost"`
	// 	Port     uint   `default:"3306"`
	// 	Debug    bool   `default:"false"`
	// }
	// Redis struct {
	// 	Host     string `default:"localhost"`
	// 	Port     uint   `default:"6379"`
	// 	Protocol string `default:"tcp"`
	// 	Password string
	// }
	// Session struct {
	// 	Name    string `default:"sessionid"`
	// 	Adapter string `default:"cookie"`
	// }
	Hosts []struct {
		Name     string `default:"mk000"`
		Ip       string `default:"192.168.88.1"`
		Vendor   string `default:"mikrotik"`
		Connect  string `default:"api"`
		Dogovor  string `default:""`
		Ads      uint64 `default:"3"`
		Provider string `default:"triolan"`
		Adress   string `default:"none"`
		Ping     bool   `default:"true"`
		Username string `default:"admin" env:"SW_USER"`
		Passwd   string `default:"admin" env:"SW_PASSWD"`
	}
	Secret string `default:"secret"`
	Limit  int    `default:"5"`
}
