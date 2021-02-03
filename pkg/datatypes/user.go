package datatypes

type CatPixUserJSON struct {
	CatPixUserID int     `json:"catPixUserId"`
	Username     *string `json:"username"`
	Password     *string `json:"password"`
}
