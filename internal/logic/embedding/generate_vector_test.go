package embedding

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestGenerateVector(t *testing.T) {
	convey.Convey("Testing GenerateVector function", t, func() {
		text := "你好，吃饭了吗" // 修改为你要测试的文本
		vector, err := GenerateVector(text)
		convey.So(err, convey.ShouldBeNil)
		elements := make([]string, len(vector))
		for i, v := range vector {
			elements[i] = fmt.Sprintf("%v", v)
		}
		result := strings.Join(elements, ",")

		fmt.Printf("[%s]\n", result)
	})
}

func TestInsertVector(t *testing.T) {
	// 插入
	InsertVector()
}
