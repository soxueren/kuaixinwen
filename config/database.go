package config

var DbConfig = map[string]map[string]string{
	"mysql_dev": {
		"host":     "localhost",
		"username": "root",
		"password": "Admin888",
		"port":     "3306",
		"database": "kuaixinwen",
		"charset":  "utf8",
		"protocol": "tcp",
		"driver":   "mysql", // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
	},
	"mssql_dev": {
		"host":     "192.168.200.248",
		"username": "gcore",
		"password": "gcore",
		"port":     "3306",
		"database": "test",
		//"charset":  "utf8",
		//"protocol": "tcp",
		"driver":   "mssql",
	},
	"postgres_dev": {
		"host":     "localhost",
		"username": "root",
		"password": "",
		"port":     "3306",
		"database": "test",
		//"charset":  "utf8",
		//"protocol": "tcp",
		"driver": "postgres",
	},
	"oracle_dev": {
		//"host":     "localhost",
		"username": "root",
		"password": "",
		//"port":     "3306",
		"database": "test",
		//"charset":  "utf8",
		//"protocol": "tcp",
		"driver": "oracle",
	},
	"sqlite_dev": {
		//"host":     "localhost",
		//"username": "root",
		//"password": "",
		//"port":     "3306",
		"database": "./foo.db",
		//"charset":  "utf8",
		//"protocol": "tcp",
		"driver": "sqlite",
	},
}
