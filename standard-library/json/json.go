package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//json解析


type (
	gResult struct{
		GSearchResultClass string `json:"GSearchResultClass"`
		UnescapedURL string `json:"unescapedUrl"`
		URL string `json:"url"`
		VisibleUrl string `json:"visibleUrl"`
		CacheUrl string `json:"cacheUrl"`
		Title string `json:"title"`
		TitleNoFormatting string `json:"titleNoFormatting"`
		Content string `json:"content"`
	}

	gResponse struct{
		ResponseData struct{
			Results []gResult `json:"result"`
		} `json:"responseData"`
	}
)


func main(){
	url:="http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	r,error:=http.Get(url)
	if error!=nil {
		log.Println("error",error)
		return
	}

	defer r.Body.Close()

	var gResponse gResponse
	error=json.NewDecoder(r.Body).Decode(&gResponse)
	if error!=nil {
		log.Println("error",error)
		return
	}

	fmt.Println(gResponse)



}