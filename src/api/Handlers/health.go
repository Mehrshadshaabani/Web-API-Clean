// HealthHandler handles health-related endpoints.
//
// NewHealthHandle creates and returns a new HealthHandler.
//
// Health godoc
// @Summary      Health check and add person
// @Description  Binds JSON body to a Person, appends to persons slice, and returns the person data.
// @Tags         health
// @Accept       json
// @Produce      json
// @Param        person  body      Person  true  "Person data"
// @Success      200     {object}  Person
// @Router       /health [post]
package handlers

import "github.com/gin-gonic/gin"

type HealthHandler struct {
}

func (h *HealthHandler) NewHealthHandle() *HealthHandler {
	return &HealthHandler{}
}

var persons []Person

// HealthHandler handles health-related endpoints.
//
// NewHealthHandle creates and returns a new HealthHandler.
//
// Health godoc
// @Summary      Health check and add person
// @Description  Binds JSON body to a Person, appends to persons slice, and returns the person data.
// @Tags         health
// @Accept       json
// @Produce      json
// @Param        person  body      Person  true  "Person data"
// @Success      200     {object}  Person
// @Router       /health [post]
func (h *HealthHandler) Health(c *gin.Context) {
	var person Person
	c.ShouldBindBodyWithJSON(&person)
	persons = append(persons, person)
	c.JSON(200, gin.H{"id": person.ID, "name": person.Name, "age": person.Age})
}

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
