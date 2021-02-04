package datatypes

import (
	"time"
)

type PictureJSON struct {
	PictureID    int        `json:"pictureId"`
	CatPixUserID int        `json:"catPixUserId"`
	FileName     *string    `json:"fileName"`
	CreateDate   *time.Time `json:"createDate,omitempty"`
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
}
