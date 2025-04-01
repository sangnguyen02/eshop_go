package setting

import (
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
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

// Setup initializes the configuration instance
func Setup() {
	// Tải file .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	// App settings
	AppSetting.JwtSecret = getEnv("JWT_SECRET", "default-secret")
	AppSetting.PageSize = getEnvAsInt("PAGE_SIZE", 10)
	AppSetting.PrefixUrl = getEnv("PREFIX_URL", "http://127.0.0.1:8000")
	AppSetting.RuntimeRootPath = getEnv("RUNTIME_ROOT_PATH", "runtime/")

	AppSetting.ImageSavePath = getEnv("IMAGE_SAVE_PATH", "upload/images")
	AppSetting.ImageMaxSize = getEnvAsInt("IMAGE_MAX_SIZE", 10*1024*1024)
	AppSetting.ImageAllowExts = getEnvAsSlice("IMAGE_ALLOW_EXTS", []string{".jpg", ".jpeg", ".png"}, ",")

	AppSetting.ExportSavePath = getEnv("EXPORT_SAVE_PATH", "export/")
	AppSetting.QrCodeSavePath = getEnv("QRCODE_SAVE_PATH", "qrcode/")
	AppSetting.FontSavePath = getEnv("FONT_SAVE_PATH", "fonts/")

	AppSetting.LogSavePath = getEnv("LOG_SAVE_PATH", "logs/")
	AppSetting.LogSaveName = getEnv("LOG_SAVE_NAME", "log")
	AppSetting.LogFileExt = getEnv("LOG_FILE_EXT", "log")
	AppSetting.TimeFormat = getEnv("TIME_FORMAT", "20060102")

	// Server settings
	ServerSetting.RunMode = getEnv("RUN_MODE", "debug")
	ServerSetting.HttpPort = getEnvAsInt("HTTP_PORT", 8000)
	ServerSetting.ReadTimeout = time.Duration(getEnvAsInt("READ_TIMEOUT", 60)) * time.Second
	ServerSetting.WriteTimeout = time.Duration(getEnvAsInt("WRITE_TIMEOUT", 60)) * time.Second

	// Database settings
	DatabaseSetting.Type = getEnv("DB_TYPE", "sqlserver")
	DatabaseSetting.User = getEnv("DB_USER", "sa")
	DatabaseSetting.Password = getEnv("DB_PASSWORD", "123456789")
	DatabaseSetting.Host = getEnv("DB_HOST", "localhost:1433")
	DatabaseSetting.Name = getEnv("DB_NAME", "AdventureWorks2017")
	DatabaseSetting.TablePrefix = getEnv("DB_TABLE_PREFIX", "")
}

// getEnv lấy giá trị từ biến môi trường, nếu không có thì trả về giá trị mặc định
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt lấy giá trị từ biến môi trường và chuyển thành int
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// getEnvAsSlice lấy giá trị từ biến môi trường và chuyển thành slice
func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	return strings.Split(valueStr, sep)
}
