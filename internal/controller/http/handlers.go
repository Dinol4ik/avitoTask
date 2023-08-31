package http

import (
	"avitoTask/internal/models"
	"avitoTask/internal/models/params"
	"avitoTask/internal/models/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func (s *Server) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func (s *Server) DeleteSegment(ctx *fiber.Ctx) error {
	paramsSegment := models.Segment{}
	err := ctx.BodyParser(&paramsSegment)
	if err != nil {
		return err
	}
	err = s.experimentUC.DeleteSegment(ctx, paramsSegment)
	return ctx.Status(fiber.StatusOK).JSON(
		response.DeleteSegment{
			Name: paramsSegment.Name,
		})
}

func (s *Server) AddUser(ctx *fiber.Ctx) error {
	paramsUser := models.User{}
	err := ctx.BodyParser(&paramsUser)
	if err != nil {
		return err
	}
	err = s.experimentUC.AddUser(ctx, paramsUser)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(
		response.AddUserResponse{
			Name: paramsUser.Name,
		})
}
func (s *Server) AddSegment(ctx *fiber.Ctx) error {
	paramsSegment := models.Segment{}
	err := ctx.BodyParser(&paramsSegment)
	if err != nil {
		return err
	}
	err = s.experimentUC.AddSegment(ctx, paramsSegment)
	return ctx.Status(fiber.StatusOK).JSON(
		response.AddSegment{
			Name: paramsSegment.Name,
		})
}
func (s *Server) AddSegmentsForUser(ctx *fiber.Ctx) error {
	segmentParams := params.SegmentParams{}
	err := ctx.BodyParser(&segmentParams)
	if err != nil {
		return err
	}
	err = s.experimentUC.AddSegmentsForUser(ctx, segmentParams)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.Status(fiber.StatusOK).JSON(
		response.AddSegmentsForUser{
			Name: segmentParams.SegmentName,
		})
}

func (s *Server) GetUserSegments(ctx *fiber.Ctx) error {
	userId := params.GetSegmentsParams{}
	err := ctx.BodyParser(&userId)
	if err != nil {
		return err
	}
	result, err := s.experimentUC.GetUserSegments(ctx, userId.UserId)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(
		response.GetUserSegments{
			Name: result,
		})
}

func (s *Server) DeleteUserSegments(ctx *fiber.Ctx) error {
	segmentParams := params.SegmentParams{}
	err := ctx.BodyParser(&segmentParams)
	if err != nil {
		return err
	}
	err = s.experimentUC.DeleteUserSegments(ctx, segmentParams)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.Status(fiber.StatusOK).JSON(
		response.DeleteUserSegments{
			SegmentsName: segmentParams.SegmentName,
		})
}
func (s *Server) CreateCsv(ctx *fiber.Ctx) error {
	csvParams := models.DateFilter{}
	err := ctx.BodyParser(&csvParams)
	if err != nil {
		return err
	}
	filename, err := s.experimentUC.CsvCreate(ctx, csvParams)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	url := fmt.Sprintf("http://localhost:8080/save-history/%s", filename)
	return ctx.Status(fiber.StatusOK).JSON(
		response.CreateCsv{
			Url: url,
		})
}

func (s *Server) SaveCsv(ctx *fiber.Ctx) error {
	fileName := ctx.Params("fileName")
	_, err := os.Stat("history/" + fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

	} else {
		return ctx.Download("history/"+fileName, "user-history.csv")
	}
	return ctx.SendStatus(fiber.StatusBadRequest)
}
