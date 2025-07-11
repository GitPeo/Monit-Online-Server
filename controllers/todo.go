package controllers

import (
	"fmt"
	"monit/global"
	"monit/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fetch(ctx *gin.Context) {
	var todos []models.Todo
	if result := global.Db.Where("deleted = ?", false).Order("`order` DESC").Find(&todos); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	todoResponses := models.TodosToResponses(todos)
	ctx.JSON(http.StatusOK, gin.H{
		"todos": todoResponses,
	})

}

func Update(ctx *gin.Context) {
	var todoReqs []models.TodoReq
	// 尝试从请求体中解析 JSON 数据到 todos 切片
	if err := ctx.ShouldBindJSON(&todoReqs); err != nil {
		// 若解析失败，返回 400 状态码和错误信息
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("todoReqs: ")
	fmt.Println(todoReqs)
	// 调用 update 函数更新数据库中的记录
	if err := update(todoReqs); err != nil {
		// 若更新过程中出现错误，返回 400 状态码和错误信息
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 更新成功，返回 200 状态码和成功信息
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// update 批量更新或插入 Todo 记录
func update(todos []models.TodoReq) error {
	var allTodos []models.Todo
	// 从数据库获取所有 Todo 记录
	if err := global.Db.Find(&allTodos).Error; err != nil {
		return err
	}

	// 创建一个映射，键为 Tid，值为 *Todo 指针
	collect := make(map[int]*models.Todo)
	for i := range allTodos {
		collect[allTodos[i].Tid] = &allTodos[i]
	}
	fmt.Println("collect: ")
	fmt.Println(collect)

	var newTodos []models.Todo
	var updateTodos []models.Todo

	// 遍历传入的 TodoReq 列表
	for _, item := range todos {
		oldTodo, exists := collect[item.Tid]
		if !exists {
			// 若记录不存在，创建新的 Todo 记录
			todo1 := models.RequestToTodo(item)
			newTodos = append(newTodos, todo1)
		} else if oldTodo.Version < item.Version {
			todo2 := models.RequestToTodo(item)
			todo2.Id = oldTodo.Id
			updateTodos = append(updateTodos, todo2)
		}
	}

	// 批量插入新的 Todo 记录
	if len(newTodos) > 0 {
		if err := global.Db.Create(&newTodos).Error; err != nil {
			return err
		}
	}

	// 批量更新现有的 Todo 记录
	fmt.Println("updateTodos: ")
	for _, item := range updateTodos {
		global.Db.Model(&item).Where("id = ?", item.Id).Updates(item)
	}

	return nil
}
