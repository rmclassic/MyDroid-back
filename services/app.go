package services

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "strconv"
  "mydroid/models"
  "encoding/json"
  "net/http"
)

func GetAllApps(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
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
        return
      }

    tempapp.ImageUrl = fmt.Sprintf("/assets/thumbs/%d.jpg", tempapp.ID)
    tempapp.Publisher = publisher.Name
    apps = append(apps, tempapp)
  }


  payload, _ := json.Marshal(apps)
  c.JSON(200, gin.H{
    "result": "success",
    "data": string(payload),
  })
}

func GetBest(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
  pagestr := c.DefaultQuery("page", "0")
  category := c.Param("category")
  per_pagestr := c.DefaultQuery("per_page", "10")

  page, _ := strconv.Atoi(pagestr)
  per_page, _ := strconv.Atoi(per_pagestr)
  query := fmt.Sprintf("SELECT *, COUNT(app_id) AS cnt FROM ((app JOIN download ON app.id=download.app_id) JOIN (select name as category, app_id as capp_id from app_category JOIN category ON category_id=id) cat ON cat.capp_id=app_id) WHERE category='%s' GROUP BY app_id ORDER BY cnt DESC LIMIT %d, %d", category, page, per_page)
  println(query)
  rows, err := db.Query(query)
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
      rows.Scan(&tempapp.ID, &tempapp.Name, &tempapp.Description, &pid, &wow, &wow, &wow, &wow, &wow, &tempapp.Category, nil, nil)
      publisher, err := GetAccountByID(pid)
      if err != nil {
        c.JSON(200, gin.H{
          "result": "fail",
          "message": err,
        })
      }

      tempapp.ImageUrl =  fmt.Sprintf("/assets/thumbs/%d.jpg", tempapp.ID)
      tempapp.Publisher = publisher.Name
      apps = append(apps, tempapp)
    }

    c.JSON(200, gin.H{
      "result": "success",
      "data": apps,
    })
}

func GetLatest(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
  category := c.Param("category")


  query := fmt.Sprintf("select app.id, app.name, description, date_modified, publisher_id, account.name, app_category.category_id, category.name FROM app JOIN account ON publisher_id=account.id JOIN app_category ON app.id=app_category.app_id JOIN category ON category_id=category.id where category.name='%s' ORDER BY date_modified DESC;", category)
  rows, err := db.Query(query)
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
  }

  apps := make([]models.App, 0)

  for rows.Next() {

      var tempapp models.App
      var wow string
      rows.Scan(&tempapp.ID, &tempapp.Name, &tempapp.Description, &wow, &wow, &tempapp.Publisher, &wow, &tempapp.Category)

      tempapp.ImageUrl =  fmt.Sprintf("/assets/thumbs/%d.jpg", tempapp.ID)
      apps = append(apps, tempapp)
    }

    c.JSON(200, gin.H{
      "result": "success",
      "data": apps,
    })
}



func GetAppById(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
  }

  query := fmt.Sprintf("select description, date_modified, app_name, account.name as publisher , category.name as category, app_id from ((select *, name AS app_name from app JOIN app_category ON id=app_category.app_id) d JOIN category ON d.category_id=category.id JOIN account ON publisher_id=account.id) WHERE app_id=%d", id)
  println(query)
  rows, err := db.Query(query)
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
  }
  var app models.App
  s := ""
  if rows.Next() {
    rows.Scan(&app.Description, &s, &app.Name, &app.Publisher, &app.Category, nil)
    app.ImageUrl = fmt.Sprintf("/assets/thumbs/%d.jpg", id)
    app.DownloadUrl = fmt.Sprintf("/app/download/%d", id)
    app.ID = id
  }
  c.JSON(200, gin.H{
    "result": "success",
    "message": app,
  })
}

func DownloadApp(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
  }

  query := fmt.Sprintf("INSERT INTO download VALUES (%d, %d)", id, 3) //reconsider user after auth
  db.Exec(query)
  c.Redirect(http.StatusFound, fmt.Sprintf("http://localhost:8080/assets/apps/%d.apk", id))

}
