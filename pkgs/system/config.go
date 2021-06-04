package system

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v3"
)

var (
	SystemConfigVar SystemConfig
)

type SystemConfig struct {
	NetCoCo struct {
		Port int `yaml:"port"`
		DB   struct {
			SQL struct {
				Hostname string `yaml:"hostname"`
				Username string `yaml:"username"`
				Password string `yaml:"password"`
				DBName   string `yaml:"db_name"`
			} `yaml:"sql"`
		} `yaml:"db"`
		Security struct {
			Salt          string `yaml:"salt"`
			HttpAuthUsers []struct {
				Username string `yaml:"username"`
				Password string `yaml:"password"`
			} `yaml:"http_auth_users"`
		} `yaml:"security"`
	} `yaml:"netcoco"`
}

func ParseSystemConfig(path string) {
	filename := filepath.Clean(path)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &SystemConfigVar)
	if err != nil {
		panic(err)
	}

}

func SetConfigFromEnvironment() {
	if os.Getenv("COCO_PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("COCO_PORT"))
		if err != nil {
			log.Fatal(err)
			os.Exit(0)
		}
		SystemConfigVar.NetCoCo.Port = port
	} else {
		SystemConfigVar.NetCoCo.Port = 8080
	}

	db_hostname := os.Getenv("COCO_DB_HOSTNAME")
	db_username := os.Getenv("COCO_DB_USERNAME")
	db_password := os.Getenv("COCO_DB_PASSWORD")
	db_name := os.Getenv("COCO_DB_NAME")
	salt := os.Getenv("COCO_SEC_SALT")

	if db_hostname == "" {
		log.Panic("Please enter COCO_DB_HOSTNAME environment")
	} else if db_username == "" {
		log.Panic("Please enter COCO_DB_USERNAME environment")
	} else if db_password == "" {
		log.Panic("Please enter COCO_DB_PASSWORD environment")
	} else if db_name == "" {
		log.Panic("Please enter COCO_DB_NAME environment")
	} else if salt == "" {
		log.Panic("Please enter COCO_SEC_SALT environment")
	}

	SystemConfigVar.NetCoCo.DB.SQL.Hostname = db_hostname
	SystemConfigVar.NetCoCo.DB.SQL.Username = db_username
	SystemConfigVar.NetCoCo.DB.SQL.Password = db_password
	SystemConfigVar.NetCoCo.DB.SQL.DBName = db_name

	SystemConfigVar.NetCoCo.Security.Salt = salt

}
