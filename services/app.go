package services

import (
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "fmt"
  "encoding/json"
)

func UploadApp(c *gin.Context) {
  content, err := ioutil.ReadAll(c.Request.Body)
  if err != nil {
      c.JSON(400, gin.H{
        "result": "failed",
        "message": err,
      })
    }

    var data map[string]interface{}
    err = json.Unmarshal(content, &data)
    if err != nil {
      c.JSON(400, gin.H{
        "result": "failed",
        "message": err,
      })
    }

    if data["name"] == nil || data["description"] == nil || data["publisher_id"] == nil {
      c.JSON(400, gin.H{
        "result": "failed",
        "message": "provided data is not sufficent",
      })
    }

    _, err = db.Exec(fmt.Sprintf("INSERT INTO app(name, description, publisher_id) VALUES('%s', '%s', %d)", data["name"].(string), data["description"].(string), int(data["publisher_id"].(float64))))
    if err != nil {
      c.JSON(400, gin.H{
        "result": "failed",
        "message": err, 
      })
    }
}
