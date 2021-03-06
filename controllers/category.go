package controllers

import (
	"strconv"

	"github.com/zmisgod/blogApi/models"
)

//CategoryController 分类的控制器
type CategoryController struct {
	BaseController
}

//@router /:categoryId [get]
func (t *CategoryController) Get() {
	var (
		err        error
		categoryID int
	)
	if categoryID, err = strconv.Atoi(t.Ctx.Input.Param(":categoryId")); err != nil {
		t.SendJSON(404, "", "invalid params")
	}
	res, err := models.TagAll(categoryID, t.page, t.pageSize, "category")
	t.CheckError(err)
	t.SendJSON(200, res, "successful")
}
