package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func TestRedis() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	fmt.Println("hello from redis")
}

type RedisAdapter struct {
	C redis.Conn
}

func (this *RedisAdapter) initRedis() {
	var err error
	this.C, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	// defer func() {
	// 	fmt.Println("redis closing")
	// 	this.c.Close()
	// 	fmt.Println("redis closed")
	// }()
}

func (this *RedisAdapter) Set(name string, value string) {
	_, err := this.C.Do("SET", name, value)
	if err != nil {
		fmt.Println("redis set failed", err)
	}
}

func (this *RedisAdapter) Get(name string) {
	value, err := redis.String(this.C.Do("GET", name))
	if err != nil {
		fmt.Println("redis get failed", err)
	} else {
		fmt.Printf("Got name %v \n", value)
	}
}

func (this *RedisAdapter) Lpush(name string, value interface{}) {
	_, err := this.C.Do("LPUSH", name, value)
	if err != nil {
		fmt.Println("redis lpush failed", err)
	}
}

func (this *RedisAdapter) Lrange(name string, start int64, end int64) ([]string, error) {
	result, err := this.C.Do("LRANGE", name, start, end)
	if err != nil {
		fmt.Println("redis lrange failed", err)
	}
	var result_list []string
	result_list, err = redis.Strings(result, err)
	if err != nil {
		fmt.Println("redis strings failed", err)
	}

	for i, v := range result_list {
		fmt.Println(i)
		fmt.Println(v)
	}
	return result_list, err
}

func (this *RedisAdapter) LrangeInt(name string, start int64, end int64) ([]int, error) {
	result_list, err := redis.Ints(this.C.Do("LRANGE", name, start, end))
	return result_list, err
}

func (this *RedisAdapter) IterateHash() {
	this.C.Send("HMSET", "balbum:1", "title", "Red", "rating", 5)
	this.C.Send("HMSET", "balbum:2", "title", "Earthbound", "rating", 10)
	this.C.Send("HMSET", "balbum:3", "title", "Beat")
	this.C.Send("LPUSH", "albums", "1")
	this.C.Send("LPUSH", "albums", "2")
	this.C.Send("LPUSH", "albums", "3")
	values, err := redis.Values(this.C.Do("SORT", "balbums",
		"BY", "balbum:*->rating",
		"GET", "balbum:*->title",
		"GET", "balbum:*->rating"))
	if err != nil {
		panic(err)
	}

	for len(values) > 0 {
		var title string
		rating := -1 // initialize to illegal value to detect nil.
		values, err = redis.Scan(values, &title, &rating)
		if err != nil {
			panic(err)
		}
		if rating == -1 {
			fmt.Println(title, "not-rated")
		} else {
			fmt.Println(title, rating)
		}
	}
}

var _redis_instance *RedisAdapter

func RedisInstance() *RedisAdapter {
	if _redis_instance == nil {
		_redis_instance = new(RedisAdapter)
		fmt.Println("redis instance initialzing")
		_redis_instance.initRedis()
		fmt.Println("redis instance initialzed")
	}
	return _redis_instance
}

func init() {

}
