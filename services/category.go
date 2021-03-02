package services

import (
  "github.com/gin-gonic/gin"
  "mydroid/models"
)

func GetAllCategories(c *gin.Context) {
  query := "SELECT id, name, parent_id FROM category"
  rows, err := db.Query(query)
  if err != nil {
    c.JSON(500, gin.H{
      "result": "fail",
      "message": err,
    })
  }

  categories := make([]models.Category, 0)

  for rows.Next() {
    var category models.Category
    rows.Scan(&category.ID, &category.Name, &category.ParentID)
    categories = append(categories, category)
  }

  c.JSON(200, gin.H{
    "result": "services",
    "message": categories,
  })
}
