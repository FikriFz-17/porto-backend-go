package dto

import "mime/multipart"

type SkillResponse struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

type CreateSkillRequest struct {
	Name     string                `form:"name" binding:"required"`
	IconFile *multipart.FileHeader `form:"icon" binding:"required"`
}

type UuidSkillRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type UpdateSkillRequest struct {
	UUID     string                `json:"-"`
	Name     string                `form:"name" json:"name" binding:"required"`
	IconFile *multipart.FileHeader `form:"icon" `
}
