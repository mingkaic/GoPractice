package redisAPI

import (
	"testing"
)

func testWrap(test func(*RedisCli)) {
	// before each
	cli := CreateLocal()
	cli.Connect()
	defer cli.Close()

	test(cli)
}

func TestConn(t *testing.T) {
	testWrap(func(cli *RedisCli) {
		expected := "+PONG"
		resp, err := cli.Ping()
		if err != nil {
			t.Error(err)
		}
		if expected != resp {
			t.Errorf("expected response %s, got %s", expected, resp)
		}
	})
}

func TestSetGet(t *testing.T) {
	testWrap(func(cli *RedisCli) {
		input := []string{"key1234", "value3456"}
		cli.Set(input[0], input[1])
		result := cli.Get(input[0])
		if input[1] != result {
			t.Errorf("expected value %s, got %s", input[1], result)
		}
		cli.Del(input[0])
	})
}

func TestDel(t *testing.T) {
	testWrap(func(cli *RedisCli) {
		input := "key1234"
		ignoreState := cli.Del(input)
		cli.Set(input, "value3456")
		absentState := cli.Del(input)
		if ignored != ignoreState {
			t.Error("expected to be ignored, got absent")
		}
		if absent != absentState {
			t.Error("expected to be absent, got ignored")
		}
	})
}
