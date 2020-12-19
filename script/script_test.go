package script

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	Execute("pwd")
	Execute("echo trying to fail\nkdkdk\npwd")
}
