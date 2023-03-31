package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/dimfeld/httptreemux/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	router *httptreemux.TreeMux
	log    *zap.SugaredLogger
)

type Data struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}

func GetProductionLogger(appName string, appVersion string) (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": appName,
		"version": appVersion,
	}
	lb, err := config.Build()
	if err != nil {
		return nil, err
	}
	return lb.Sugar(), nil
}

func myHandler(w http.ResponseWriter, r *http.Request, m map[string]string) {
	fmt.Fprint(w, "Hello, World!")
	log.Infow("/hello called")
}
func rootHandler(w http.ResponseWriter, r *http.Request, m map[string]string) {
	now := time.Now()
	dat := []Data{}
	dat = append(dat, Data{1, "Amarnath", "/", now})
	dat = append(dat, Data{2, "Amar", "/", now})

	jData, err := json.Marshal(dat)
	if err != nil {
		// handle error
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
	//fmt.Fprint(w, "Hello, Root!")
	log.Infow("/ called")
}

func RegisterRoutes(router *httptreemux.TreeMux) {
	router.Handle(http.MethodGet, "/hello", myHandler)
	router.GET("/", rootHandler)
}

func init() {
	log, _ = GetProductionLogger("BasicFunction", "0.0.1")
	router = httptreemux.New()
	RegisterRoutes(router)

	funcframework.RegisterHTTPFunction("/", func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r)
	})

}

func main() {

	funcframework.Start("3000")
}
