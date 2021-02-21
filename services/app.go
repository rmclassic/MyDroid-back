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

    apps = append(apps, tempapp)
  }


  tempapp.Publisher = publisher.Name
  payload, _ := json.Marshal(tempapp)
  c.JSON(200, gin.H{
    "result": "fail",
    "data": string(payload),
  })
}

func GetBestApps(c *gin.Context) {
  pagestr := c.DefaultQuery("page", "0")
  per_pagestr := c.DefaultQuery("per_page", "10")
  page, _ := strconv.Atoi(pagestr)
  per_page, _ := strconv.Atoi(per_pagestr)

  if err != nil {
    err, _ := db.Query(fmt.Sprintf("SELECT app_id, COUNT(app_id) AS cnt from (app JOIN download ON app.id=download.app_id) GROUP BY app_id ORDER BY cnt DESC LIMIT %d, %d", page, per_page))
    c.JSON(200, gin.H{
      "result": "fail",
      "data": string(payload),
    })
  }

  var tempapp models.App
  rows.Scan(&tempapp.ID, &tempapp.Name, &tempapp.Description, &pid)
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
}
