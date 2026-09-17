package settings

import "github.com/isv933/go-samples/url-shortener/duration"

type Settings struct {
	ListenAddress      string            `json:"listen_address"`
	ServerTimeout      duration.Duration `json:"server_timeout"`
	DbConnectionString string            `json:"db_connection_string"`
}

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}

func NewSettings() Settings {
	return Settings{ListenAddress: ":18000", ServerTimeout: must(duration.ParseISO8601("PT5S")), DbConnectionString: "postgresql://test:test123@server.lan:55432/go_samples"}
}
