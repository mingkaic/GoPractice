package redisAPI

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const (
	absent  = iota
	ignored = iota
)

type ConnInfo struct {
	host string
	port uint16
}

type RedisCli struct {
	opts ConnInfo
	conn net.Conn
}

// ======== Public ========

func CreateLocal() *RedisCli {
	return Create("127.0.0.1", 6379)
}

func Create(host string, port uint16) *RedisCli {
	return &RedisCli{opts: ConnInfo{
		host: host,
		port: port,
	}}
}

func (cli *RedisCli) Connect() error {
	uri := fmt.Sprintf("%s:%d", cli.opts.host, cli.opts.port)
	var err error
	cli.conn, err = net.DialTimeout("tcp", uri, time.Minute)
	return err
}

func (cli *RedisCli) Close() {
	cli.conn.Close()
}

func (cli *RedisCli) Ping() (resp string, err error) {
	command := "PING"
	return cli.process(command)
}

func (cli *RedisCli) Set(key, value string) {
	command := fmt.Sprintf("SET %s %s", key, value)
	_, err := cli.process(command)
	logIfErr(err)
}

func (cli *RedisCli) Get(key string) string {
	command := fmt.Sprintf("GET %s", key)
	resp, err := cli.process(command)
	logIfErr(err)
	return resp
}

func (cli *RedisCli) Del(key string) int {
	state := absent
	command := fmt.Sprintf("DEL %s", key)
	resp, err := cli.process(command)
	logIfErr(err)

	if resp != ":1" {
		state = ignored
	}

	return state
}

func (redis *RedisCli) Publish() {

}

func (redis *RedisCli) Subscribe() {

}

// ======== Private ========

func (cli *RedisCli) process(command string) (resp string, err error) {
	fmt.Fprintf(cli.conn, command+"\r\n")
	b := bufio.NewReader(cli.conn)
	resp, err = b.ReadString('\n')
	if err != nil {
		return
	}
	if resp[0] == '$' {
		resp, err = b.ReadString('\n')
	}
	if err != nil {
		return
	}
	resp = strings.TrimSpace(resp)
	return
}

func logIfErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
