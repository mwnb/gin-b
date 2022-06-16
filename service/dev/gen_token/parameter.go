package gentoken

type GenTokenParameter struct {
	Name    string `form:"name" binding:"required"`
	Account string `form:"account" binding:"required"`
	Extra   string `form:"extra"`
}
