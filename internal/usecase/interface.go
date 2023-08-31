package usecase

import (
	"avitoTask/internal/models"
	"avitoTask/internal/models/params"
	"github.com/gofiber/fiber/v2"
)

type Experiment interface {
	AddUser(ctx *fiber.Ctx, user models.User) error
	AddSegment(ctx *fiber.Ctx, segment models.Segment) error
	DeleteSegment(ctx *fiber.Ctx, segment models.Segment) error
	AddSegmentsForUser(ctx *fiber.Ctx, params params.SegmentParams) error
	GetUserSegments(ctx *fiber.Ctx, userId int) (models.UserSegments, error)
	DeleteUserSegments(ctx *fiber.Ctx, params params.SegmentParams) error
	CsvCreate(ctx *fiber.Ctx, params models.DateFilter) (string, error)
}
