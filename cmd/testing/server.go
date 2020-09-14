package testing

import (
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/repo"
	"github.com/dhuki/rest-template-v2/pkg/testing/infrastructure/command"
	"github.com/dhuki/rest-template-v2/pkg/testing/infrastructure/query"
	"github.com/dhuki/rest-template-v2/pkg/testing/presenter"
	"github.com/dhuki/rest-template-v2/pkg/testing/usecase"
	"github.com/dhuki/rest-template-v2/utils"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-kit/kit/log"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type server interface {
	Start()
}

type testingServer struct {
	mux          *mux.Router
	db           *gorm.DB
	elasticlient *elasticsearch.Client
	redisClient  *redis.Client
	middlewares  []mux.MiddlewareFunc
	utils        utils.Utils
	logger       log.Logger
}

func NewServer(mux *mux.Router) testingServer {
	return testingServer{
		mux: mux,
	}
}

func (t testingServer) AddPersistent(db *gorm.DB, elasticlient *elasticsearch.Client, redisClient *redis.Client) testingServer {
	t.db = db
	t.elasticlient = elasticlient
	t.redisClient = redisClient
	return t
}

func (t testingServer) AddUtils(utils utils.Utils) testingServer {
	t.utils = utils
	return t
}

func (t testingServer) AddMiddlewares(middlewares []mux.MiddlewareFunc) testingServer {
	t.middlewares = middlewares
	return t
}

func (t testingServer) AddLogger(logger log.Logger) testingServer {
	t.logger = logger
	return t
}

func (t testingServer) Build() server {
	return t
}

func (t testingServer) Start() {
	presenter.NewHttpHandler(
		t.mux,
		usecase.UsecaseImpl{
			TestTableRepo: repo.NewTestTableRepo(
				command.NewTestTableCommand(t.db),
				query.NewTestTableQuery(t.elasticlient),
			),
			Utils: t.utils,
		},
		t.middlewares,
		t.logger,
	)
}
