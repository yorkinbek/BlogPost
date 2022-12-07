package handlers

import (
	"net/http"
	"strconv"

	"yorqinbek/microservices/Blogpost/api_gateway/models"
	"yorqinbek/microservices/Blogpost/api_gateway/protogen/blogpost"

	"github.com/gin-gonic/gin"
)

// CreateArticle godoc
// @Summary     Create article
// @Description create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article       body     models.CreateArticleModel true  "article body"
// @Param       Authorization header   string                    false "Authorization"
// @Success     201           {object} models.JSONResponse{data=models.Article}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/article [post]
func (h handler) CreateArticle(c *gin.Context) {
	var body models.CreateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	article, err := h.grpcClients.Article.CreateArticle(c.Request.Context(), &blogpost.CreateArticleRequest{
		Content: &blogpost.Content{
			Title: body.Title,
			Body:  body.Body,
		},
		AuthorId: body.AuthorID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Article | GetList",
		Data:    article,
	})
}

// GetArticleByID godoc
// @Summary     get article by id
// @Description get an article by id
// @Tags        articles
// @Accept      json
// @Param       id            path   string true  "Article ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/article/{id} [get]
func (h handler) GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	article, err := h.grpcClients.Article.GetArticleByID(c.Request.Context(), &blogpost.GetArticleByIDRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}

// GetArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "smth"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Article}
// @Router      /v1/article [get]
func (h handler) GetArticleList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.grpcClients.Article.GetArticleList(c.Request.Context(), &blogpost.GetArticleListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    articleList,
	})
}

// SearchArticleByMyUsername godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Article}
// @Router      /v1/my-articles [get]
func (h handler) SearchArticleByMyUsername(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)

	usernameRaw, ok := c.Get("auth_username")
	username, ok := usernameRaw.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, "Something is worng")
		return
	}

	searchStr := username

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.grpcClients.Article.GetArticleList(c.Request.Context(), &blogpost.GetArticleListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    articleList,
	})
}

// UpdateArticle godoc
// @Summary     Update article
// @Description update a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article       body     models.UpdateArticleModel true  "article body"
// @Param       Authorization header   string                    false "Authorization"
// @Success     200           {object} models.JSONResponse{data=models.Article}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/article [put]
func (h handler) UpdateArticle(c *gin.Context) {
	var body models.UpdateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.grpcClients.Article.UpdateArticle(c.Request.Context(), &blogpost.UpdateArticleRequest{
		Content: &blogpost.Content{
			Title: body.Title,
			Body:  body.Body,
		},
		Id: body.ID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}

// DeleteArticle godoc
// @Summary     delete article by id
// @Description delete an article by id
// @Tags        articles
// @Accept      json
// @Param       id            path   string true  "Article ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.PackedArticleModel}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/article/{id} [delete]
func (h handler) DeleteArticle(c *gin.Context) {
	idStr := c.Param("id")

	article, err := h.grpcClients.Article.DeleteArticle(c.Request.Context(), &blogpost.DeleteArticleRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    article,
	})
}
