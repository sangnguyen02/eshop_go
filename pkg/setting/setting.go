package setting

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string

	FileMaxSize   int
	FileAllowExts []string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Port        string
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v. Using default values.", err)
	}

	AppSetting.JwtSecret = os.Getenv("JWT_SECRET")
	AppSetting.PageSize, _ = strconv.Atoi(os.Getenv("PAGE_SIZE"))
	AppSetting.PrefixUrl = os.Getenv("PREFIX_URL")
	// AppSetting.RuntimeRootPath = os.Getenv("RUNTIME_ROOT_PATH")

	AppSetting.ImageSavePath = os.Getenv("IMAGE_SAVE_PATH")
	AppSetting.ImageMaxSize, _ = strconv.Atoi(os.Getenv("IMAGE_MAX_SIZE"))
	AppSetting.ImageAllowExts = strings.Split(os.Getenv("IMAGE_ALLOW_EXTS"), ",")
	AppSetting.FileMaxSize, _ = strconv.Atoi(os.Getenv("FILE_MAX_SIZE"))

	AppSetting.ExportSavePath = os.Getenv("EXPORT_SAVE_PATH")
	AppSetting.QrCodeSavePath = os.Getenv("QRCODE_SAVE_PATH")
	AppSetting.FontSavePath = os.Getenv("FONT_SAVE_PATH")

	AppSetting.LogSavePath = os.Getenv("LOG_SAVE_PATH")
	AppSetting.LogSaveName = os.Getenv("LOG_SAVE_NAME")
	AppSetting.LogFileExt = os.Getenv("LOG_FILE_EXT")
	AppSetting.TimeFormat = os.Getenv("TIME_FORMAT")

	// Server
	ServerSetting.RunMode = os.Getenv("RUN_MODE")
	ServerSetting.HttpPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))
	readTimeout, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	ServerSetting.ReadTimeout = time.Duration(readTimeout) * time.Second
	writeTimeout, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	ServerSetting.WriteTimeout = time.Duration(writeTimeout) * time.Second

	// Database
	DatabaseSetting.Port = os.Getenv("DB_PORT")
	DatabaseSetting.Type = os.Getenv("DB_TYPE")
	DatabaseSetting.User = os.Getenv("DB_USER")
	DatabaseSetting.Password = os.Getenv("DB_PASSWORD")
	DatabaseSetting.Host = os.Getenv("DB_HOST")
	DatabaseSetting.Name = os.Getenv("DB_NAME")
	DatabaseSetting.TablePrefix = os.Getenv("DB_TABLE_PREFIX")
}
