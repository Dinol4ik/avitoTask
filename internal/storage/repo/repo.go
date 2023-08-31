package repo

import (
	"avitoTask/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repo {
	return Repo{db}
}

const addUserSql = `INSERT INTO public.user (name) VALUES ($1) `

func (r *Repo) AddUser(ctx *fiber.Ctx, user string) error {
	_, err := r.db.ExecContext(ctx.Context(), addUserSql, user)
	if err != nil {
		return err
	}
	return nil
}

const addSegmentSql = `INSERT INTO public.segment(name) VALUES ($1)`

func (r *Repo) AddSegment(ctx *fiber.Ctx, segment string) error {
	_, err := r.db.ExecContext(ctx.Context(), addSegmentSql, segment)
	if err != nil {
		return err
	}
	return nil
}

const deleteSegmentSql = `
	UPDATE public.segment
	SET is_removed = true
	WHERE name = $1 returning id`

func (r *Repo) DeleteSegment(ctx *fiber.Ctx, segment string) (int, error) {
	var segmentId int
	err := r.db.GetContext(ctx.Context(), &segmentId, deleteSegmentSql, segment)
	if err != nil {
		return 0, err
	}
	return segmentId, nil
}

const addSegmentForUserSql = `
		INSERT INTO public.user_segment(user_id, segment_id, is_removed)
		VALUES ($1,
        	(SELECT id from segment WHERE name = $2), (SELECT is_removed FROM segment WHERE name = $2))`

func (r *Repo) AddSegmentsForUser(ctx *fiber.Ctx, segmentName string, userId int) error {
	_, err := r.db.ExecContext(ctx.Context(), addSegmentForUserSql, userId, segmentName)
	if err != nil {
		return err
	}

	return nil
}

const deleteSegmentForUserSql = `UPDATE user_segment
	SET is_removed = true, updated_at = now()
	WHERE user_id = $1 and segment_id = any ($2)`

func (r *Repo) DeleteSegmentsForUser(ctx *fiber.Ctx, segmentsId []int, userId int) error {
	_, err := r.db.ExecContext(ctx.Context(), deleteSegmentForUserSql, userId, pq.Array(segmentsId))
	if err != nil {
		return err
	}

	return nil
}

const deleteSegmentForAllUserSql = `UPDATE user_segment
	SET is_removed = true, updated_at = now()
	WHERE segment_id =  $1`

func (r *Repo) DeleteSegmentForAllUsers(ctx *fiber.Ctx, segmentsId int) error {
	_, err := r.db.ExecContext(ctx.Context(), deleteSegmentForAllUserSql, segmentsId)
	if err != nil {
		return err
	}

	return nil
}

const getUserSegments = `SELECT s.name from user_segment
               INNER JOIN public.segment s on s.id = user_segment.segment_id
               WHERE user_id = $1 and user_segment.is_removed = false`

func (r *Repo) GetUserSegments(ctx *fiber.Ctx, userId int) ([]models.Segment, error) {
	var segments []models.Segment
	err := r.db.SelectContext(ctx.Context(), &segments, getUserSegments, userId)
	if err != nil {
		return nil, err
	}

	return segments, err
}

const segmentsIdSql = `SELECT id FROM segment WHERE name = ANY($1)`

func (r *Repo) GetSegmentsIdByNames(ctx *fiber.Ctx, segmentNames []string) (models.SegmentsId, error) {
	var segmentsId models.SegmentsId
	err := r.db.SelectContext(ctx.Context(), &segmentsId.Id, segmentsIdSql, pq.Array(segmentNames))
	if err != nil {
		return models.SegmentsId{}, err
	}

	return segmentsId, err
}

const csvCreate = `select u.name as user_name,s.name as segment_name,created_at,updated_at,user_segment.is_removed
                from user_segment
                inner join public.segment s on s.id = user_segment.segment_id
                inner join public."user" u on u.id = user_segment.user_id
                WHERE user_id = $1 and created_at >= $2 and created_at <= $3`

func (r *Repo) CsvCreate(ctx *fiber.Ctx, dateStart string, dateEnd string, userId int) ([]models.Csv, error) {
	var selectRows []models.Csv
	err := r.db.SelectContext(ctx.Context(), &selectRows, csvCreate, userId, dateStart, dateEnd)
	if err != nil {
		return []models.Csv{}, err
	}

	return selectRows, err
}

const usersWithoutSegmentSql = `select id 
				from "user"
				order by random()
				limit (SELECT count(*) from "user")*$1::float`

func (r *Repo) FindUsers(ctx *fiber.Ctx, percent float64) ([]models.UserId, error) {
	var userId []models.UserId
	err := r.db.SelectContext(ctx.Context(), &userId, usersWithoutSegmentSql, percent)
	if err != nil {
		return nil, err
	}
	return userId, nil
}
