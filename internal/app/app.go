package app

import (
	//"LaTest/internal/usecase/repo"
	//warehouse "LaTest/internal/usecase/warehouse"
	"avitoTask/internal/app/config"
	"avitoTask/internal/controller/http"
	"avitoTask/internal/storage/repo"
	"avitoTask/internal/usecase/experiment"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
	//"net/rpc"
	//"os"
	//"os/signal"
	//"syscall"
)

type App struct {
	cfg *config.Config
}

func (a *App) Run() error {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zapcore.Level(*a.cfg.Logger.Level))
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		os.Stdout,
		atom,
	)
	logger := zap.New(zapCore)
	logger = logger.With(zap.String("service", "experiment"))
	log := logger.Sugar()
	atom.SetLevel(zapcore.Level(*a.cfg.Logger.Level))
	log.Infof("logger initialized successfully")

	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", a.cfg.Postgres.Host, a.cfg.Postgres.User, a.cfg.Postgres.Password, a.cfg.Postgres.DBName)
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	repository := repo.New(db)

	experimentUC := experiment.NewExperimentUsecase(&repository, log)

	httpServer := http.NewServer(*a.cfg, *log, experimentUC)
	log.Info("application has started")
	go httpServer.Run()

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	log.Debug("waiting for httpServer to shut down")

	log.Info("application has been shut down")

	return nil
}

func New(cfg *config.Config) *App {
	return &App{cfg}
}
