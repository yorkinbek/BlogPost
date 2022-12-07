package handlers

// import (
// 	"net/http"
// 	"yorqinbek/microservices/Blogpost/article_service/models"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// // CreateAuthor godoc
// // @Summary     Create author
// // @Description create a new author
// // @Tags        authors
// // @Accept      json
// // @Produce     json
// // @Param       author body     models.CreateAuthorModel true "author body"
// // @Success     201    {object} models.JSONResponse{data=models.Author}
// // @Failure     400    {object} models.JSONErrorResponse
// // @Router      /v1/author [post]
// func (h handler) CreateAuthor(c *gin.Context) {
// 	var body models.CreateAuthorModel
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
// 		return
// 	}

// 	// TODO - validation should be here

// 	id := uuid.New()

// 	err := h.Stg.AddAuthor(id.String(), body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
// 			Error: err.Error(),
// 		})
// 		return
// 	}

// 	author, err := h.Stg.GetAuthorByID(id.String())
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
// 			Error: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, models.JSONResponse{
// 		Message: "Author | GetList",
// 		Data:    author,
// 	})
// }

// // GetAuthorByID godoc
// // @Summary     get author by id
// // @Description get an author by id
// // @Tags        authors
// // @Accept      json
// // @Param       id path string true "Author ID"
// // @Produce     json
// // @Success     200 {object} models.JSONResponse{data=models.Author}
// // @Failure     400 {object} models.JSONErrorResponse
// // @Router      /v1/author/{id} [get]
// func (h handler) GetAuthorByID(c *gin.Context) {
// 	idStr := c.Param("id")

// 	// TODO - validation

// 	author, err := h.Stg.GetAuthorByID(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
// 			Error: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, models.JSONResponse{
// 		Message: "OK",
// 		Data:    author,
// 	})
// }

// // GetAuthorList godoc
// // @Summary     List author
// // @Description get author
// // @Tags        authors
// // @Accept      json
// // @Produce     json
// // @Success     200 {object} models.JSONResponse{data=[]models.Author}
// // @Router      /v1/author [get]
// func (h handler) GetAuthorList(c *gin.Context) {
// 	authorList, err := h.Stg.GetAuthorList()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
// 			Error: err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, models.JSONResponse{
// 		Message: "OK",
// 		Data:    authorList,
// 	})
// }
