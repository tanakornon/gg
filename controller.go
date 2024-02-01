package gg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MountRoutes[T Entities](r gin.IRouter, service IService[T]) {
	r.GET("", getAll[T](service))
	r.GET("/:id", get[T](service))
	r.POST("", create[T](service))
	r.POST("/many", createMany[T](service))
	r.PUT("/:id", upsert[T](service))
	r.PATCH("/:id", update[T](service))
	r.DELETE("/:id", delete[T](service))
}

func getAll[T Entities](service IService[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query QueryParams
		c.BindQuery(&query)

		res, err := service.GetAll(query)
		if err != nil {
			c.JSON(HandleStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func get[T Entities](service IService[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := ParseUInt(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := service.Get(id)
		if err != nil {
			c.JSON(HandleStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func create[T Entities](service IService[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := service.Create(req)
		if err != nil {
			c.JSON(HandleStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}

func createMany[T Entities](service IService[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req []T
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := service.CreateMany(req)
		if err != nil {
			c.JSON(HandleStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)
	}
}

func upsert[T Entities](service IService[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := ParseUInt(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req T
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := service.Upsert(id, req)
		if err != nil {
			c.JSON(HandleStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func update[T Entities](service IService[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := ParseUInt(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var req T
		if err := c.ShouldBindWith(&req, PartialBinding{}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := service.Update(id, req)
		if err != nil {
			c.JSON(HandleStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func delete[T Entities](service IService[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := ParseUInt(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.Delete(id)
		if err != nil {
			c.JSON(HandleStatusCode(err), gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
