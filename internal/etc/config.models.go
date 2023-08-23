package etc

type Configuration struct {
	Db   db
	Web  web
	File file
}

type web struct {
	Listen string
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}

type file struct {
	TemplateFilePath string
	StorageFilePath  string
	ConvertURL       string
}
