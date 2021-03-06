package api

import (
	"strconv"

	"github.com/bo-er/todo-simpler/admin/response"
	"github.com/bo-er/todo-simpler/todo"
	"github.com/gin-gonic/gin"
)

// Todo 定义Todo APi
type Todo struct {
	TodoService todo.TodoService
}

type paginationInput struct {
	Page     int `form:"page"`      //当前页
	PageSize int `form:"page_size"` // 每页大小
}

type addUserTodoInput struct {
	UserID              int    `json:"user_id"`                                             //用户id(新增)
	UserTodoTitle       string `json:"user_todo_title" binding:"required"`                  //用户Todo的标题
	UserTodoDescription string `json:"user_todo_description"`                               //用户Todo的描述
	UserTodoDueTime     string `json:"user_todo_due_time" example:"2021-01-22 16:09:00"`    //用户Todo截止时间
	UserTodoRemindTime  string `json:"user_todo_remind_time" example:"2021-01-21 16:09:00"` //用户Todo提前多久通知
	Status              bool   `json:"status"`                                              //用户Todo状态
}

type findUserTodo struct {
	UserID     string `json:"user_id"`      //用户id
	UserTodoID string `json:"user_todo_id"` //用户Todo id
}

// AddUserTodo 是添加新的UserTodo的路由处理函数
func (t *Todo) AddUserTodo(c *gin.Context) {
	var json addUserTodoInput
	if err := c.ShouldBindJSON(&json); err != nil {
		response.ResError(c, err)
		return
	}
	var status int
	if json.Status {
		status = 1
	} else {
		status = 0
	}
	userID := strconv.Itoa(json.UserID)
	userTodoID, err := t.TodoService.AddUserTodo(userID, json.UserTodoTitle, json.UserTodoDescription,
		json.UserTodoDueTime, json.UserTodoRemindTime, status)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}

// GetUserTodo 获得一条todo的路由处理函数
func (t *Todo) GetUserTodo(c *gin.Context) {
	var json findUserTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		//fmt.Print(err)
		response.ResError(c, err)
		return
	}

	//userid,err:=strconv.Atoi(json.UserID)
	userTodoID, err := t.TodoService.GetUserTodo(json.UserID, json.UserTodoID)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}

// GetUserAllTodos 获得所有todo的路由处理函数
func (t *Todo) GetUserAllTodos(c *gin.Context) {
	// var page paginationInput
	// if err := c.ShouldBind(&page); err != nil {
	// 	response.ResError(c, err)
	// 	return
	// }
	page, _ := strconv.Atoi(c.PostForm("page"))
	pagesize, _ := strconv.Atoi(c.PostForm("page_size"))
	resultTodo, err := t.TodoService.GetUserAllTodos(page, pagesize)
	// fmt.Print(page)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, resultTodo)
}

//UpdateUserTodo 更新单个todo的路由处理函数
func (t *Todo) UpdateUserTodo(c *gin.Context) {
	var json addUserTodoInput
	if err := c.ShouldBindJSON(&json); err != nil {
		response.ResError(c, err)
		return
	}
	var status int
	if json.Status {
		status = 1
	} else {
		status = 0
	}
	userID := strconv.Itoa(json.UserID)
	userTodoID, err := t.TodoService.UpdateUserTodo(userID, json.UserTodoTitle, json.UserTodoDescription,
		json.UserTodoDueTime, json.UserTodoRemindTime, status)
	if err != nil {
		response.ResError(c, err)
	}
	response.ResSuccess(c, userTodoID)
}
