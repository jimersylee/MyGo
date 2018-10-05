package testing

import (
	"net/http"
	"testing"
)

const URL="http://blog.jimersylee.com/"
const RIGHT="\u2713"
const WRONG="\u2717"
/**
	使用go的测试框架进行测试
方法必须Test开头,参数为testing.T的指针,不能有返回值
 */
func TestDownload(t *testing.T) {
	var urls= []struct {url string
		statusCode int
	}{
		{
			"http://blog.jimersylee.com/",
			http.StatusOK,
		},
		{
			"http://blog.jimersylee.com/blog/go/go-hot-update2.html",
			http.StatusNotFound,
		},
	}

	for _,v:=range urls{
		response,err:=http.Get(v.url)
		defer response.Body.Close()
		t.Log("Ready to test download")
		if err!=nil {
			t.Fatal("http.Get error",WRONG,err)
		}

		t.Log("should be able to get url:",v.url,RIGHT)
		if response.StatusCode==v.statusCode{
			t.Log("should be right status code",v.statusCode,RIGHT)
		}else{
			t.Fatal("Checking http status error",WRONG,response.Status)
		}
	}

}