package http

import (
	"avitoTask/internal/app/config"
	"avitoTask/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
)

type Server struct {
	config       *config.Config
	logger       *zap.SugaredLogger
	experimentUC usecase.Experiment
}

func NewServer(config config.Config, logger zap.SugaredLogger, expUC usecase.Experiment) *Server {
	return &Server{
		config:       &config,
		logger:       &logger,
		experimentUC: expUC,
	}
}

func (s *Server) Run() error {
	app := fiber.New()
	app.Use(logger.New())
	s.RegisterRoutes(app)
	return app.Listen(s.config.HTTPServer.Address)
}

func (s *Server) RegisterRoutes(app *fiber.App) {
	app.Get("/api/health-check", s.HealthCheck)
	app.Post("/add", s.AddUser)                             // +
	app.Post("/add-segment", s.AddSegment)                  //+
	app.Post("/add-segment-for-user", s.AddSegmentsForUser) // +
	app.Get("/get-user-segments", s.GetUserSegments)        // +
	app.Put("/delete-user-segments", s.DeleteUserSegments)  // +
	app.Delete("/delete-segment", s.DeleteSegment)          // +
	app.Post("/get-user-history", s.CreateCsv)
	app.Get("/save-history/:fileName", s.SaveCsv)
}
