package services

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "strconv"
  "io/ioutil"
  "mydroid/models"
  "encoding/json"
)

func GetCommentsForApp(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
    return
  }

  query := fmt.Sprintf("select app.name as app, account.name as user, content from comment JOIN app ON app_id=app.id JOIN account ON user_id=account.id where app.id=%d", id)

  rows, err := db.Query(query)
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
    return
  }


  comments := make([]models.Comment, 0)

  for rows.Next() {
    var comment models.Comment

    rows.Scan(&comment.Sender, &comment.App, &comment.Content)
    comments = append(comments, comment)
  }


  c.JSON(200, gin.H{
    "result": "success",
    "message": comments,
  })
}


func PostComment(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
  c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

  var root map[string]interface{}

  body, err := ioutil.ReadAll(c.Request.Body)
  if err != nil {
    c.JSON(200, gin.H{
      "result": "fail",
      "message": err,
    })
    return
  }

  token, _ := c.Cookie("X-AUTH")
  userid := GetUserByToken(token)

  if userid == -1 {
      c.JSON(403, gin.H{
        "result": "fail",
        "message": "Not authorized to do this operation",
      })
      return
  }

  json.Unmarshal(body, &root)
  query := fmt.Sprintf("INSERT INTO comment values (%d, %d, '%s')", int(root["app"].(float64)), userid, root["content"].(string))
  println(query)
  db.Exec(query)

  c.JSON(200, gin.H{
    "result": "success",
    "message": nil,
  })
}
