package rest

import (
	"github.com/asaskevich/govalidator"
	"github.com/c-4u/loyalty-card/domain/service"
	"github.com/gofiber/fiber/v2"
)

type RestService struct {
	Service *service.Service
}

func NewRestService(service *service.Service) *RestService {
	return &RestService{
		Service: service,
	}
}

// CreateGuest godoc
// @Summary create a new guest
// @ID createGuest
// @Tags Guest
// @Description Router for create a new guest
// @Accept json
// @Produce json
// @Param body body CreateGuestRequest true "JSON body for create a new guest"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guests [post]
func (t *RestService) CreateGuest(c *fiber.Ctx) error {
	var req CreateGuestRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	guestID, err := t.Service.CreateGuest(c.Context(), &req.Name, &req.Doc)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *guestID})
}

// FindGuest godoc
// @Summary find a gust
// @ID findGuest
// @Tags Guest
// @Description Router for find a guest
// @Accept json
// @Produce json
// @Param guest_id path string true "Guest ID"
// @Success 200 {object} Guest
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /guests/{guest_id} [get]
func (t *RestService) FindGuest(c *fiber.Ctx) error {
	guestID := c.Params("guest_id")
	if !govalidator.IsUUIDv4(guestID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "guest_id is not a valid uuid",
		})
	}

	guest, err := t.Service.FindGuest(c.Context(), &guestID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(guest)
}

// CreateScore godoc
// @Summary create a new score
// @ID createScore
// @Tags Score
// @Description Router for create a new score
// @Accept json
// @Produce json
// @Param body body CreateScoreRequest true "JSON body for create a new score"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /scores [post]
func (t *RestService) CreateScore(c *fiber.Ctx) error {
	var req CreateScoreRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	scoreID, err := t.Service.CreateScore(c.Context(), &req.Date, &req.GuestID, &req.Description)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *scoreID})
}

// FindScore godoc
// @Summary find a score
// @ID findScore
// @Tags Score
// @Description Router for find a score
// @Accept json
// @Produce json
// @Param score_id path string true "Score ID"
// @Success 200 {object} Score
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /scores/{score_id} [get]
func (t *RestService) FindScore(c *fiber.Ctx) error {
	scoreID := c.Params("score_id")
	if !govalidator.IsUUIDv4(scoreID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "score_id is not a valid uuid",
		})
	}

	score, err := t.Service.FindScore(c.Context(), &scoreID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(score)
}

// UseScore godoc
// @Summary use a score
// @ID useScore
// @Tags Score
// @Description Router for use a score
// @Accept json
// @Produce json
// @Param score_id path string true "Score ID"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /scores/{score_id}/use [post]
func (t *RestService) UseScore(c *fiber.Ctx) error {
	scoreID := c.Params("score_id")
	if !govalidator.IsUUIDv4(scoreID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "score_id is not a valid uuid",
		})
	}

	err := t.Service.UseScore(c.Context(), &scoreID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// AddTag godoc
// @Summary add a tag to a score
// @ID addTag
// @Tags Score
// @Description Router for add a tag to a score
// @Accept json
// @Produce json
// @Param score_id path string true "Score ID"
// @Param body body AddTagRequest true "JSON body for add a tag to a score"
// @Success 200 {object} HTTPResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /scores/{score_id}/tags [post]
func (t *RestService) AddTag(c *fiber.Ctx) error {
	var req AddTagRequest

	scoreID := c.Params("score_id")
	if !govalidator.IsUUIDv4(scoreID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "score_id is not a valid uuid",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	err := t.Service.AddTag(c.Context(), &scoreID, &req.TagID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(HTTPResponse{Msg: "successful request"})
}

// CreateTag godoc
// @Summary create a new tag
// @ID createTag
// @Tags Tag
// @Description Router for create a new tag
// @Accept json
// @Produce json
// @Param body body CreateTagRequest true "JSON body for create a new tag"
// @Success 200 {object} IDResponse
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /tags [post]
func (t *RestService) CreateTag(c *fiber.Ctx) error {
	var req CreateTagRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	tagID, err := t.Service.CreateTag(c.Context(), &req.Name)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(IDResponse{ID: *tagID})
}

// FindTag godoc
// @Summary find a tag
// @ID findTag
// @Tags Tag
// @Description Router for find a tag
// @Accept json
// @Produce json
// @Param tag_id path string true "Tag ID"
// @Success 200 {object} Tag
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /tags/{tag_id} [get]
func (t *RestService) FindTag(c *fiber.Ctx) error {
	tagID := c.Params("tag_id")
	if !govalidator.IsUUIDv4(tagID) {
		return c.Status(fiber.StatusBadRequest).JSON(HTTPResponse{
			Msg: "tag_id is not a valid uuid",
		})
	}

	tag, err := t.Service.FindTag(c.Context(), &tagID)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}

// SearchTags godoc
// @Summary search tags
// @ID searchTags
// @Tags Tag
// @Description Router for search tags
// @Accept json
// @Produce json
// @Param q query SearchTagsRequest false "query for search tags"
// @Success 200 {array} Tag
// @Failure 400 {object} HTTPResponse
// @Failure 403 {object} HTTPResponse
// @Router /tags [get]
func (t *RestService) SearchTags(c *fiber.Ctx) error {
	var req SearchTagsRequest

	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HTTPResponse{Msg: err.Error()})
	}

	tag, err := t.Service.SearchTags(c.Context(), &req.Name)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(HTTPResponse{Msg: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}
