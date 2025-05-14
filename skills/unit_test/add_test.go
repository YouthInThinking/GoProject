package unittest_test

import (
	"os"
	"testing"

	unittest "github.com/YouthInThinking/GoProject/skills/unit_test"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	os.Getenv("CONFIG_PATH")
	//判断逻辑的第一种方式：通过程序逻辑判断
	// result := unittest.Add(2, 3)
	// if result != 5 {
	// 	t.Errorf("Expected 5, got %d", result)
	// } else {
	// 	t.Logf("Test passed: Add(2, 3) = %d", result)
	// }

	//判断第二种方式：专门的断言库
	should := assert.New(t)

	should.Equal(unittest.Add(2, 3), 5)

}
