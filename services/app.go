package services

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "strconv"
  "mydroid/models"
  "encoding/json"
)

func GetAllApps(c *gin.Context) {
  pagestr := c.DefaultQuery("page", "0")
  per_pagestr := c.DefaultQuery("per_page", "10")
  page, _ := strconv.Atoi(pagestr)
  per_page, _ := strconv.Atoi(per_pagestr)
  _, _ = page, per_page
  rows, err := db.Query(fmt.Sprintf("SELECT id,name,description,publisher_id FROM app LIMIT %d, %d", page, per_page))
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
  }

  apps := make([]models.App, 0)

  for rows.Next() {
    var pid int

    var tempapp models.App
    rows.Scan(&tempapp.ID, &tempapp.Name, &tempapp.Description, &pid)

    //fmt.Println(string(tempapp.Name))
    //rows.Scan(&tempapp.ID, &tempapp.Name, &tempapp.Description, &tempapp.Publisher)
    publisher, err := GetAccountByID(pid)
    if err != nil {
      c.JSON(200, gin.H{
        "result": "fail",
        "message": err,
      })
    }

    tempapp.Publisher = publisher.Name
    payload, _ := json.Marshal(tempapp)
    c.JSON(200, gin.H{
      "result": "fail",
      "data": string(payload),
    })

    apps = append(apps, tempapp)
  }


}
