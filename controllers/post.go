package controllers

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"quickstart/helper"
	"quickstart/models"
	"time"
)

func init() {

}

type PostController struct {
	baseApiController
}

func (this *PostController) Prepare() {
	this.baseApiController.Prepare()

	schemaString := map[string]string{
		"POST": `
		{
			"properties": {
				"content": {"type": ["string", "null"], "default": "" },
                "pics": {
                    "type": ["array", "null"],
                    "items": {
                        "type": "object",
                        "properties": {
                            "desc": {"type": "string"},
                            "thumbnail": {"type": "string"},
                            "url": {"type": "string"}
                        }
                    }
                },
				"feed_id": { 
					"type": "number" 
				}
			},
			"type": "object"
		}`,
		"GET": `
		{
			"properties": {
				"get_key": {
					"type": "string"
				}
			},
			"required": ["get_key"],
			"type": "object"
		}`,
	}

	_, err := helper.Validate_Json(
		schemaString[this.Ctx.Request.Method], this.Data["input_map"].(map[string]interface{}))
	if err != nil {
		fmt.Println(err.Error())
		this.Abort("400")
	}
}

type Post struct {
	Content string `redis:"content"`
	Pics    string `redis:"pics"`
	Flag    int64  `redis:"flag"`
	Time    int64  `redis:"time"`
}

func (this *PostController) Post() {
	models.RedisInstance().C.Do("INCR", "post_next_id")
	post_next_id, _ := redis.Int64(models.RedisInstance().C.Do("GET", "post_next_id"))
	if this.Data["input_map"].(map[string]interface{})["pics"] == nil && this.Data["input_map"].(map[string]interface{})["content"] == nil {
		this.Abort("400")
	}

	var post Post
	post.Content, _ = this.Data["input_map"].(map[string]interface{})["content"].(string)
	post.Flag = 2323232
	post.Pics, _ = this.Data["input_map"].(map[string]interface{})["pics"].(string)
	post.Time = time.Now().Unix()

	if _, err := models.RedisInstance().C.Do("HMSET",
		redis.Args{}.Add("post:"+string(post_next_id)).AddFlat(&post)...); err != nil {
		fmt.Println(err)
	}

	v, err := redis.Values(models.RedisInstance().C.Do("HGETALL", "post:"+string(post_next_id)))
	if err != nil {
		fmt.Println(err)
	}

	var post1 Post
	if err := redis.ScanStruct(v, &post1); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", post1)

	this.Data["json"] = map[string]interface{}{
		"post": "hello world"}
	this.ServeJson()

}
