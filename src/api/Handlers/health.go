package handlers

import "github.com/gin-gonic/gin"

type HealthHandler struct {
}

func (h *HealthHandler) NewHealthHandle() *HealthHandler {
	return &HealthHandler{}
}

var persons []Person

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
