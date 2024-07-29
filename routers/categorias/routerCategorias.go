package routers



import(

	"github.com/gin-gonic/gin"
    "net/http"

)


func ConfigRoutersCategory(r *gin.Engine){

	r.GET("/all", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":"OK",
		})
	})

	r.POST("/categorias", func(c *gin.Context) {

		name:= c.PostForm("name")
		c.JSON(http.StatusOK,gin.H{

			"message": "bao " + name,
		} )
	})

	r.GET("/categoria/:name", func(c *gin.Context) {

		name:= c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "bao " + name,
		})
	})
}