package main

import (
	"anshoryihsan/simple_rest/app"
	"anshoryihsan/simple_rest/controller"
	"anshoryihsan/simple_rest/helper"
	"anshoryihsan/simple_rest/middleware"
	"anshoryihsan/simple_rest/repository"
	"anshoryihsan/simple_rest/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/categoryId", categoryController.Update)
	router.DELETE("/api/categories/categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	// Handler: router,

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
