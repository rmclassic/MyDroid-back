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

    tempapp.Publisher = publisher.Name
    apps = append(apps, tempapp)
  }


  payload, _ := json.Marshal(apps)
  c.JSON(200, gin.H{
    "result": "success",
    "data": string(payload),
  })
}

func GetBestApps(c *gin.Context) {
  pagestr := c.DefaultQuery("page", "0")
  category := c.Param("category")
  per_pagestr := c.DefaultQuery("per_page", "10")

  page, _ := strconv.Atoi(pagestr)
  per_page, _ := strconv.Atoi(per_pagestr)

  rows, err := db.Query(fmt.Sprintf("SELECT *, COUNT(app_id) AS cnt FROM ((app JOIN download ON app.id=download.app_id) JOIN (select name as category, app_id as capp_id from app_category JOIN category ON category_id=id) cat ON cat.capp_id=app_id) WHERE category='%s' GROUP BY app_id ORDER BY cnt DESC LIMIT %d, %d", category, page, per_page))
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
      var wow string
      rows.Scan(&tempapp.ID, &tempapp.Name, &tempapp.Description, &pid, &wow, &wow, &wow, &tempapp.Category, nil, nil)
      publisher, err := GetAccountByID(pid)
      if err != nil {
        c.JSON(200, gin.H{
          "result": "fail",
          "message": err,
        })
      }

      tempapp.Publisher = publisher.Name
      apps = append(apps, tempapp)
    }

    payload, _ := json.Marshal(apps)
    c.JSON(200, gin.H{
      "result": "success",
      "data": string(payload),
    })
}
