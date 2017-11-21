package Utils

type Config struct {
	ID      string
	Service Service
	Backup  Backup
}

type Service struct {
	Route       string
	EndPoint    string
	ContentType string
}

type Backup struct {
	Dir  string
	File string
}
