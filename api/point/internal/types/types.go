// Code generated by goctl. DO NOT EDIT.
package types

type GetUserEarnRequest struct {
	UserId  int64 `form:"user_id" validate:"required"`
	PointId int64 `form:"point_id" validate:"omitempty"`
}

type GetUserEarnResponse struct {
	Id          int64  `json:"id"`
	Point       int64  `json:"point"`
	Description string `json:"description"`
	ExpiredAt   string `json:"expired_at"`
	CreatedAt   string `json:"created_at"`
}

type GetUserUseRequest struct {
	UserId  int64 `form:"user_id" validate:"required"`
	PointId int64 `form:"point_id" validate:"omitempty"`
}

type GetUserUseResponse struct {
	Id          int64  `json:"id"`
	Point       int64  `json:"point"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type GetUserPointRequest struct {
	UserId int64 `path:"user_id"`
}

type GetUserPointResponse struct {
	Point int64 `json:"point"`
}
