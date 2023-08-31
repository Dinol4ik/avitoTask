package experiment

import (
	"avitoTask/internal/models"
	"avitoTask/internal/models/params"
	"avitoTask/internal/storage"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
	"time"
)

type UseCase struct {
	repo   storage.Repo
	logger *zap.SugaredLogger
}

func NewExperimentUsecase(r storage.Repo, logger *zap.SugaredLogger) *UseCase {
	return &UseCase{
		repo:   r,
		logger: logger,
	}
}
func (uc *UseCase) AddUser(ctx *fiber.Ctx, user models.User) error {
	err := uc.repo.AddUser(ctx, user.Name)
	if err != nil {
		return err
	}
	return nil
}
func (uc *UseCase) AddSegment(ctx *fiber.Ctx, segment models.Segment) error {
	if segment.PercentUsers == 0 {
		err := uc.repo.AddSegment(ctx, segment.Name)
		if err != nil {
			return err
		}
	}
	result, err := uc.repo.FindUsers(ctx, segment.PercentUsers)
	if err != nil {
		return err
	}
	for _, segmentName := range result {
		err = uc.repo.AddSegmentsForUser(ctx, segment.Name, segmentName.Id)
		if err != nil {
			return err
		}
	}
	return nil
}
func (uc *UseCase) DeleteSegment(ctx *fiber.Ctx, segment models.Segment) error {
	deletedSegmentId, err := uc.repo.DeleteSegment(ctx, segment.Name)
	if deletedSegmentId != 0 {
		err = uc.repo.DeleteSegmentForAllUsers(ctx, deletedSegmentId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uc *UseCase) AddSegmentsForUser(ctx *fiber.Ctx, params params.SegmentParams) error {
	for _, segmentName := range params.SegmentName {
		err := uc.repo.AddSegmentsForUser(ctx, segmentName, params.UserId)
		if err != nil {
			return err
		}
	}
	return nil
}
func (uc *UseCase) GetUserSegments(ctx *fiber.Ctx, userId int) (models.UserSegments, error) {
	userSegments := models.UserSegments{}
	segments, err := uc.repo.GetUserSegments(ctx, userId)
	if err != nil {
		return userSegments, err
	}

	userSegments.UserId = userId
	userSegments.Segments = segments

	return userSegments, nil
}

func (uc *UseCase) DeleteUserSegments(ctx *fiber.Ctx, params params.SegmentParams) error {
	segmentsId, err := uc.repo.GetSegmentsIdByNames(ctx, params.SegmentName)
	if err != nil {
		return err
	}
	err = uc.repo.DeleteSegmentsForUser(ctx, segmentsId.Id, params.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UseCase) CsvCreate(ctx *fiber.Ctx, params models.DateFilter) (string, error) {
	querryResult, err := uc.repo.CsvCreate(ctx, params.DateStart, params.DateEnd, params.UserId)
	filename := fmt.Sprintf("%d.csv", time.Now().Unix())
	clientsFile, err := os.OpenFile("history/"+filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer clientsFile.Close()

	if _, err = clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}
	err = gocsv.MarshalFile(&querryResult, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		return "", err
	}

	return filename, nil
}
