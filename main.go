package main

import (
	"fmt"

	homework03 "github.com/jheader/go-homework1-template/homeworkgorm03"
)

func main() {

	user := homework03.User{}

	mp := make(map[string]any)
	// 键：字段名（name）；值：切片[表达式, 参数]
	mp["name"] = []string{"Alice", "Alice1"}
	users, err := (&user).SelectPageByMap(1, 3, mp)
	if err != nil {
		// 错误处理（生产环境建议用日志框架，而非fmt）
		fmt.Printf("分页查询用户失败：%v\n", err)
		return // 或根据业务逻辑返回错误、重试等
	}

	// 3. 优化遍历：直接使用元素e，无需索引i
	for _, user := range users {
		fmt.Println(user)
	}

}
