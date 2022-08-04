package database

var connection string

//  function init akan selalu di akses pertama kali meskipun tidak di initialisasi
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
