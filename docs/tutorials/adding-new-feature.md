# 新增功能教學

本教學將引導您如何在 Go Template 專案中新增一個新的功能模組。我們將以新增一個 "產品" (Product) 功能為例。

## 步驟

1.  **定義資料模型 (`internal/models/product.go`)**：

    ```go
    package models

    import (
        "gorm.io/gorm"
    )

    type Product struct {
        gorm.Model
        Name        string  `json:"name" binding:"required"`
        Description string  `json:"description"`
        Price       float64 `json:"price" binding:"required"`
    }
    ```

2.  **定義資料庫操作 (`internal/repository/product.go`)**：

    ```go
    package repository

    import (
        "go-template/internal/models"

        "gorm.io/gorm"
    )

    type Product interface {
        Create(product *models.Product) error
        GetByID(id uint) (*models.Product, error)
        Update(product *models.Product) error
        Delete(id uint) error
    }

    type Product struct {
        db *gorm.DB
    }

    func NewProduct(db *gorm.DB) Product {
        return &Product{db: db}
    }

    func (r *Product) Create(product *models.Product) error {
        return r.db.Create(product).Error
    }

    func (r *Product) GetByID(id uint) (*models.Product, error) {
        var product models.Product
        err := r.db.First(&product, id).Error
        return &product, err
    }

    func (r *Product) Update(product *models.Product) error {
        return r.db.Save(product).Error
    }

    func (r *Product) Delete(id uint) error {
        return r.db.Delete(&models.Product{}, id).Error
    }
    ```

3.  **定義業務邏輯 (`internal/services/product.go` 和 `internal/services/IProduct.go`)**：

    `internal/services/product.go`:

    ```go
    package services

    import "go-template/internal/models"

    type Product interface {
        CreateProduct(product *models.Product) error
        GetProductByID(id uint) (*models.Product, error)
        UpdateProduct(product *models.Product) error
        DeleteProduct(id uint) error
    }
    ```

    `internal/services/IProduct.go`:

    ```go
    package services

    import (
        "go-template/internal/models"
        "go-template/internal/repositories"
        "go-template/internal/services"
    )

    type product struct {
        repo repositories.Product
    }

    func NewProduct(repo repositories.Product) services.Product {
        return &Product{repo: repo}
    }

    func (s *Product) CreateProduct(product *models.Product) error {
        return s.repo.Create(product)
    }

    func (s *Product) GetProductByID(id uint) (*models.Product, error) {
        return s.repo.GetByID(id)
    }

    func (s *Product) UpdateProduct(product *models.Product) error {
        return s.repo.Update(product)
    }

    func (s *Product) DeleteProduct(id uint) error {
        return s.repo.Delete(id)
    }
    ```

4.  **編寫處理函數 (`internal/api/handlers/product.go`)**：

    ```go
    package handlers

    import (
        "net/http"
        "strconv"

        "go-template/internal/api/handlers/response"
        "go-template/internal/models"
        "go-template/internal/services"

        "github.com/gin-gonic/gin"
    )

    type Product struct {
        service services.Product
    }

    func NewProduct(service services.Product) *Product {
        return &Product{service: service}
    }

    func (h *Product) CreateProduct(c *gin.Context) {
        var product models.Product
        if err := c.ShouldBindJSON(&product); err != nil {
            response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
            return
        }

        if err := h.service.CreateProduct(&product); err != nil {
            response.ErrorResponse(c, http.StatusInternalServerError, "Failed to create product", err)
            return
        }

        response.SuccessResponse(c, http.StatusCreated, "Product created successfully", product)
    }

     func (h *Product) GetProductByID(c *gin.Context) {
        id, err := strconv.ParseUint(c.Param("id"), 10, 64)
        if err != nil {
            response.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err)
            return
        }

        product, err := h.service.GetProductByID(uint(id))
        if err != nil {
            response.ErrorResponse(c, http.StatusNotFound, "Product not found", err)
            return
        }

        response.SuccessResponse(c, http.StatusOK, "Product found", product)
    }

    func (h *Product) UpdateProduct(c *gin.Context) {
        id, err := strconv.ParseUint(c.Param("id"), 10, 64)
           if err != nil {
               response.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err)
               return
           }
        var product models.Product
        if err := c.ShouldBindJSON(&product); err != nil {
            response.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err)
            return
        }
        product.ID = uint(id)
        if err := h.service.UpdateProduct(&product); err != nil {
            response.ErrorResponse(c, http.StatusInternalServerError, "Failed to update product", err)
            return
        }

        response.SuccessResponse(c, http.StatusOK, "Product updated successfully", product)
    }

    func (h *Product) DeleteProduct(c *gin.Context) {
        id, err := strconv.ParseUint(c.Param("id"), 10, 64)
        if err != nil {
            response.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err)
            return
        }

        if err := h.service.DeleteProduct(uint(id)); err != nil {
            response.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete product", err)
            return
        }

        response.SuccessResponse(c, http.StatusOK, "Product deleted successfully", nil)
    }
    ```

5.  **定義路由 (`internal/api/routes/product.go`)**：

    ```go
    package routes

    import (
        "go-template/internal/api/handlers"
        "go-template/internal/middleware"

        "github.com/gin-gonic/gin"
    )

    func SetupProduct(r *gin.Engine, Product *handlers.Product, authMiddleware *middleware.Auth) {
        products := r.Group("/products")
        {
            products.POST("", Product.CreateProduct)
            products.GET("/:id", Product.GetProductByID)
            products.PUT("/:id", authMiddleware.Handle(), Product.UpdateProduct) // 需要身份驗證
            products.DELETE("/:id", authMiddleware.Handle(), Product.DeleteProduct) // 需要身份驗證
        }
    }
    ```

6.  **在 `cmd/go-template/wire.go` 中加入新的依賴項**：

    ```go
    //+build wireinject

    package main

    import (
        "go-template/internal/api/handlers"
        "go-template/internal/api/handlers/routes"
        "go-template/internal/configs"
        "go-template/internal/middleware"
        "go-template/internal/repository"
        "go-template/internal/server"
        "go-template/internal/services"
    	impl "go-template/internal/services/impl"
        "go-template/internal/utils/database"
        "go-template/internal/utils/jwt"
        "go-template/internal/utils/logger"

        "github.com/google/wire"
    )

    func InitializeApp() (*server.Server, error) {
        wire.Build(
            configs.NewConfig,
            logger.NewLogger,
            database.NewDB,
            repository.NewUser,
            jwt.NewJWTService,
            services.NewUserImpl,
            handlers.NewUserHandler,
    		repository.NewProduct, // 新增
    		impl.NewProduct, //新增
    		handlers.NewProduct, // 新增
            middleware.NewAuth,
            routes.SetupUserRoutes,
    		routes.SetupProductRoutes, // 新增
            server.NewServer,
        )
        return &server.Server{}, nil
    }
    ```

7.  **執行 `make generate` 或在 `cmd/go-template/` 目錄下執行 `go generate ./...` 重新產生 `cmd/go-template/wire_gen.go`**。

8. **執行資料庫遷移：** 執行 `go run migrations/migrate.go`。 (需要先在 `migrations/migrate.go` 中加入 `db.AutoMigrate(&models.Product{})`)

9. **啟動伺服器並測試新的 API。**
