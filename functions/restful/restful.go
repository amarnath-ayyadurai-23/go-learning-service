package restful

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

func init() {
	log, _ = GetProductionLogger("BasicFunction", "0.0.1")

	funcframework.RegisterHTTPFunction("/", HookFunction)
	functions.HTTP("Helloworld", HookFunction)

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

// HelloWorld writes "Hello, World!" to the HTTP response.
func HookFunction(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, fmt.Sprintf("[GET] API is Up and Running @ %v\n", now))
		log.Infow("GET method called")
	case http.MethodPost:
		fmt.Fprint(w, fmt.Sprintf("[POST] API is Up and Running @ %v\n", now))
		log.Infow("POST method called")
	case http.MethodPatch:
		fmt.Fprint(w, fmt.Sprintf("[PATCH] API is Up and Running @ %v\n", now))
		log.Infow("PATCH method called")
	case http.MethodDelete:
		fmt.Fprint(w, fmt.Sprintf("[DELETE] API is Up and Running @ %v\n", now))
		log.Infow("DELETE method called")
	default:
		fmt.Fprint(w, fmt.Sprintf("API is Up and Running @ %v\n", now))
	}
}
