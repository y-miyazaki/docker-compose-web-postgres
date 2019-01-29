package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

type dbInfo struct {
	dbms            string
	host            string
	user            string
	pass            string
	name            string
	port            string
	logmode         bool
	connMaxLifetime int
	maxIdleConn     int
	maxOpenConn     int
}

// TestResponse responses common json data.
type TestResponse struct {
	Name    string `json:"name"`
	Age     uint64 `json:"age"`
	Message string `json:"message,omitempty"`
}

// TestUser struct.
type TestUser struct {
	ID   uint64 `gorm:"column:id;not null;primary_key"`
	Name string `gorm:"column:name"`
	Age  uint64 `gorm:"column:age"`
}

// TableName for TestUser struct
func (TestUser) TableName() string {
	return "test_user"
}

func main() {
	var err error
	r := chi.NewRouter()
	NewConfig()

	db, err = NewSQL()
	if err != nil {
		panic(err)
	}
	InitializeRouter(r)
	Router(r)
	http.ListenAndServe(":8080", r)
}

// NewConfig get config datas.
func NewConfig() {
	dir, _ := os.Getwd()
	viper.SetConfigName("app")
	viper.AddConfigPath(dir + "/config") // path to look for the config file in
	viper.SetConfigType("json")          // viper.SetConfigType("YAML")としてもよい
	viper.AutomaticEnv()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("change config.", e.Name)
	})
	_ = viper.ReadInConfig()

}

// NewSQL returns new SQL.
func NewSQL() (*gorm.DB, error) {
	info := dbInfo{
		host:            viper.GetString("database.host"),
		user:            viper.GetString("database.user"),
		pass:            viper.GetString("database.pass"),
		name:            viper.GetString("database.name"),
		port:            viper.GetString("database.port"),
		logmode:         viper.GetBool("database.logmode"),
		connMaxLifetime: viper.GetInt("database.conn_max_lifetime"),
		maxIdleConn:     viper.GetInt("database.max_idle_conn"),
		maxOpenConn:     viper.GetInt("database.max_open_conn"),
	}
	fmt.Printf("data = %v", info)

	connect := "host=" + info.host + " port=" + info.port + " user=" + info.user + " dbname=" + info.name + " sslmode=disable password=" + info.pass
	db, err := gorm.Open("postgres", connect)
	if err != nil {
		return nil, errors.Wrap(err, "can't open database.")
	}
	db.DB().SetConnMaxLifetime(time.Duration(info.connMaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(info.maxIdleConn)
	db.DB().SetMaxOpenConns(info.maxIdleConn)
	db.LogMode(info.logmode)
	db.SingularTable(true)
	return db, nil
}

// InitializeRouter initializes Mux and middleware
func InitializeRouter(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
}

// Router sets router.
func Router(r *chi.Mux) {
	r.HandleFunc("/", test)
}

// test function.
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	testUser := &TestUser{}
	err := db.First(testUser).Error
	if err != nil {
		fmt.Printf("error = %v", err)
	}
	response := TestResponse{
		Name:    testUser.Name,
		Age:     testUser.Age,
		Message: "test",
	}
	ResponseJSON(w, 200, response)
}

// ResponseJSON function response as json with ResponseWriter
func ResponseJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json, _ := json.Marshal(data)
		_, _ = w.Write(json)
	}
}
