package go_demo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type PromJob struct {
	App     string   `json:"app"`
	Targets []string `json:"targets"`
}

// curl -XGET localhost:8080/register -d '{"app":"test-app", "targets":["127.0.0.1:8081","127.0.0.1:8082"]}'
func RegisterServer(port int) {
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		s, _ := ioutil.ReadAll(request.Body)
		fmt.Println(string(s))
		n := PromJob{}
		json.Unmarshal(s, &n)
		fmt.Println(n)
		f, _ := os.OpenFile("/etc/prometheus//yaml/test-app.yml", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		defer f.Close()
		w := bufio.NewWriter(f)
		yml := `
- targets: ['%s']
  labels:
    job: %s`
		w.WriteString(fmt.Sprintf(yml, strings.Join(n.Targets, "', '"), n.App))
		w.Flush()
	})
	http.ListenAndServe(":"+cast.ToString(port), nil)
}