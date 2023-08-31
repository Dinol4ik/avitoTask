package storage

import (
	"avitoTask/internal/models"
	"github.com/gofiber/fiber/v2"
)

type Repo interface {
	AddUser(ctx *fiber.Ctx, user string) error
	AddSegment(ctx *fiber.Ctx, segment string) error
	DeleteSegment(ctx *fiber.Ctx, segment string) (int, error)
	AddSegmentsForUser(ctx *fiber.Ctx, segmentName string, userId int) error
	GetSegmentsIdByNames(ctx *fiber.Ctx, segmentNames []string) (models.SegmentsId, error)
	DeleteSegmentsForUser(ctx *fiber.Ctx, segmentsId []int, userId int) error
	GetUserSegments(ctx *fiber.Ctx, userId int) ([]models.Segment, error)
	DeleteSegmentForAllUsers(ctx *fiber.Ctx, segmentsId int) error
	CsvCreate(ctx *fiber.Ctx, dateStart string, dateEnd string, userId int) ([]models.Csv, error)
	FindUsers(ctx *fiber.Ctx, percent float64) ([]models.UserId, error)
}
