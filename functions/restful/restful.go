package restful

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/dimfeld/httptreemux/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log    *zap.SugaredLogger
	apiMux *httptreemux.TreeMux
	db     *sqlx.DB
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func init() {
	log, _ = GetProductionLogger("BasicFunction", "0.0.1")
	//apiMux = httptreemux.NewContextMux()
	apiMux = httptreemux.New()
	db, _ = sqlx.Connect("mysql", "user:password@tcp(localhost:7801)/learning")
	// if err != nil {
	// 	panic(err)
	// }
	//funcframework.RegisterHTTPFunction("/", HookFunction)

	functions.HTTP("HookFunction", HookFunction)

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

	// apiMux.Handle(http.MethodGet, "/root", RootFunction)
	// apiMux.Handle(http.MethodGet, "/hello", HelloFunction)
	//g := apiMux.NewGroup("/")
	// apiMux.GET("/", HomeFunction)
	// apiMux.GET("/hello", HelloFunction)

	//apiMux.ServeHttp(w, r)

	http.HandleFunc("/", Handler)
	//RootFunction(w, r)
	//user := []User{}

	// er := db.Select(&user, "select * from people")

	// if er != nil {
	// 	fmt.Println(er)
	// 	return
	// }
	// log.Infow(fmt.Sprintf("%+v", user))
}
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Hello, %s!", params["name"])
}

func Handler(w http.ResponseWriter, r *http.Request) {
	Router().ServeHTTP(w, r)
}
func Router() http.Handler {
	router := httptreemux.New()
	router.GET("/", httptreemux.HandlerFunc(HelloWorld))
	router.GET("/:name", MyHandler)
	//router.ServeHTTP(w, r)
	return router
}

func RootFunction(w http.ResponseWriter, r *http.Request) {
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
		log.Infow(fmt.Sprintf("%v method called", r.Method))
	}

}

func HelloFunction(w http.ResponseWriter, r *http.Request, m map[string]string) {
	//params := httptreemux.ContextParams(r.Context())
	fmt.Fprint(w, "Hello World\n")
	log.Infow("Hello Function called")
}
func HomeFunction(w http.ResponseWriter, r *http.Request) {
	//params := httptreemux.ContextParams(r.Context())
	fmt.Fprint(w, "Hello Home\n")
	log.Infow("Home Function called")
}
