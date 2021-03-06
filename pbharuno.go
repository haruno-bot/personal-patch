package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/haruno-bot/haruno/coolq"
	"github.com/haruno-bot/haruno/logger"
	"github.com/haruno-bot/haruno/plugins"
)

type config struct {
	Version     string `toml:"version"`
	LogsPath    string `toml:"logsPath"`
	ServerPort  int    `toml:"serverPort"`
	CQWSURL     string `toml:"cqWSURL"`
	CQHTTPURL   string `toml:"cqHTTPURL"`
	CQToken     string `toml:"cqToken"`
	ServWSURL   string `toml:"ServWSURL"`
	ServWSToken string `toml:"ServWSToken"`
	ServHTTPURL string `toml:"ServHTTPURL"`
	WebRoot     string `toml:"webroot"`
}

// haruno 晴乃机器人
// 机器人运行的全局属性
type haruno struct {
	startTime   int64
	port        int
	logpath     string
	version     string
	cqWSURL     string
	cqHTTPURL   string
	cqToken     string
	ServWSURL   string
	ServWSToken string
	ServHTTPURL string
	webRoot     string
}

const waitTime = time.Second * 15

var bot = new(haruno)

func (bot *haruno) loadConfig() {
	cfg := new(config)
	_, err := toml.DecodeFile("config.toml", cfg)
	if err != nil {
		logger.Logger.Fatalln("Haruno Initialize fialed:", err)
	}
	bot.startTime = time.Now().UnixNano() / 1e6
	bot.port = cfg.ServerPort
	bot.logpath = cfg.LogsPath
	bot.version = cfg.Version
	bot.cqWSURL = cfg.CQWSURL
	bot.webRoot = cfg.WebRoot
	bot.cqHTTPURL = cfg.CQHTTPURL
	bot.cqToken = cfg.CQToken
	bot.ServWSURL = cfg.ServWSURL
	bot.ServWSToken = cfg.ServWSToken
	bot.ServHTTPURL = cfg.ServHTTPURL
}

// Initialize 从配置文件读取配置初始化
func (bot *haruno) Initialize() {
	bot.loadConfig()
	// 设置环境变量
	os.Setenv("CQHTTPURL", bot.cqHTTPURL)
	os.Setenv("CQWSURL", bot.cqWSURL)
	os.Setenv("CQTOKEN", bot.cqToken)
	logger.Service.SetLogsPath(bot.logpath)
	logger.Service.Initialize()
	plugins.SetupPbPlugins()
	coolq.PbClient.Initialize(bot.cqToken, bot.ServWSToken)
	go coolq.PbClient.Connect(bot.cqWSURL, bot.cqHTTPURL, bot.ServWSURL)
	go coolq.PbClient.RegisterAllPlugins()
}

// Status 运行状态json格式
type Status struct {
	Go      int    `json:"go"`
	Version string `json:"version"`
	Success int    `json:"success"`
	Fails   int    `json:"fails"`
	Start   int64  `json:"start"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := new(Status)
	status.Fails = logger.Service.FailCnt()
	status.Success = logger.Service.SuccessCnt()
	status.Start = bot.startTime
	status.Version = bot.version
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	status.Go = runtime.NumGoroutine()
	json.NewEncoder(w).Encode(status)
}

// Run 启动机器人
func (bot *haruno) Run() {
	r := mux.NewRouter()

	if bot.webRoot != "" {
		_, err := os.Stat(bot.webRoot)
		if err == nil {
			logger.Logger.Println("the web page root found in", fmt.Sprintf("\"%s\"", bot.webRoot))
			page := http.FileServer(http.Dir(bot.webRoot))
			r.Methods(http.MethodGet).Path("/").Handler(page)
			r.Methods(http.MethodGet).PathPrefix("/static").Handler(page)
		}
	}

	r.Methods(http.MethodGet).Path("/status").HandlerFunc(statusHandler)
	r.Methods(http.MethodGet).Path("/logs/-/type=websocket").HandlerFunc(logger.WSLogHandler)
	r.Methods(http.MethodGet).Path("/logs/-/type=plain").HandlerFunc(logger.RawLogHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%d", bot.port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		logger.Logger.Printf("haruno is listening on http://localhost:%d\n", bot.port)

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Logger.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, os.Kill)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()

	srv.Shutdown(ctx)

	logger.Logger.Println("haruno is shutting down")

	os.Exit(0)
}

func main() {
	bot.Initialize()
	bot.Run()
}
