package handlers

import "github.com/gin-gonic/gin"

type testhandler struct {
}

func Newtesthandler() *testhandler {
	return &testhandler{}
}

func (h *testhandler) HandleTestRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Test request successful",
	})
}

func (h *testhandler) Createuser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "Create user successful",
	})
}

func (h *testhandler) GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all users successful",
	})
}
func (h *testhandler) GetuserbyId(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Get user by ID successful",
		"id":      id,
	})
}
func (h *testhandler) Headerbinder(c *gin.Context) {
	header := c.GetHeader("X-Custom-Header")
	if header == "12345678" {
		c.JSON(200, gin.H{
			"message": "Header bound successfully",
			"header":  header,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "error",
		})
	}

}
