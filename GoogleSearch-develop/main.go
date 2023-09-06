package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

// maps variable variants location in the file mapVarCode.go

var answerCode = 200

var ginLambda *ginadapter.GinLambda

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func runSearch(c *gin.Context) {

	defer returnError(c)

	googleQuery := c.Query("query")

	numReturnRes, _ := strconv.Atoi(c.Query("returnResult"))

	startReturnPage, _ := strconv.Atoi(c.Query("returnPage"))

	country := c.Query("countryCode")

	langCode := c.Query("languageCode")

	filterRes, _ := strconv.Atoi(c.Query("filterResult"))

	safeOn := c.Query("safe")

	sortBy := c.Query("sortBy")

	googleQuery = strings.Trim(googleQuery, " ")
	googleQuery = strings.Replace(googleQuery, " ", "+", -1)

	if len(googleQuery) > 2048 {
		panicWrapper("Too many symbols", http.StatusBadRequest)
	} else if googleQuery == "" {
		panicWrapper("Empty query", http.StatusBadRequest)
	}

	if (numReturnRes < 0) || (numReturnRes > 10) {
		panicWrapper("Enter a number from 1 to 10", http.StatusBadRequest)
	}

	if (startReturnPage < 0) || (startReturnPage > 90) {
		panicWrapper("Enter a number from 1 to 90", http.StatusBadRequest)
	}

	_, keyIsInMap := googleCountry[country]

	_, keyIsInMapLang := language[langCode]

	if !keyIsInMap && country != "" {
		err := fmt.Sprintf("country (%s) is currently not supported", country)
		panicWrapper(err, http.StatusBadRequest)
	} else if (keyIsInMap == true) && (country != "") {
		country = string(country)
	} else {
		country = "us"
	}

	if !keyIsInMapLang && langCode != "" {
		err := fmt.Sprintf("language (%s) is currently not supported", langCode)
		panicWrapper(err, http.StatusBadRequest)
	} else if (keyIsInMapLang == true) && (langCode != "") {
		langCode = string(langCode)
	} else {
		langCode = "en"
	}

	switch filterRes {
	case 1:
		filterRes = 1
	default:
		filterRes = 0
	}

	switch safeOn {
	case "active":
		safeOn = "active"
	default:
		safeOn = "off"
	}

	switch sortBy {
	case "date":
		sortBy = "date"
	default:
		sortBy = ""
	}

	API_KEY := os.Getenv("KEY")

	CX := "e6b3c16aa3c714755"

	url := fmt.Sprintf("https://customsearch.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s&gl=%s&num=%d&hl=%s&sort=%s&safe=%s&start=%d&filter=%d", API_KEY, CX, googleQuery, country, numReturnRes, langCode, sortBy, safeOn, startReturnPage, filterRes)

	//fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("Access-Control-Allow-Origin", "*")

	req.Header.Set("User-Agent", randomUserAgent())

	res, _ := http.DefaultClient.Do(req)

	dec := json.NewDecoder(res.Body)

	defer res.Body.Close()

	var rsp struct {
		Queries struct {
			Request []struct {
				SearchTerms    string `json:"searchTerms"`
				Count          int    `json:"count"`
				StartIndex     int    `json:"startIndex"`
				InputEncoding  string `json:"inputEncoding"`
				OutputEncoding string `json:"outputEncoding"`
				Safe           string `json:"safe"`
				Filter         string `json:"filter"`
				Gl             string `json:"gl"`
				Hl             string `json:"hl"`
				Sort           string `json:"sort"`
				TotalResults   string `json:"totalResults"`
			} `json:"request"`
		} `json:"queries"`

		Item []struct {
			Title        string      `json:"title"`
			Link         string      `json:"link"`
			DisplayLink  string      `json:"displayLink"`
			Snippet      string      `json:"snippet"`
			FormattedUrl string      `json:"formattedUrl"`
			Pagemap      interface{} `json:"pagemap"`
		} `json:"items"`
	}

	err := dec.Decode(&rsp)
	if err != nil {
		panicWrapper("No content", http.StatusNoContent)
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"queries": rsp.Queries,
			"item":    rsp.Item,
		})
	}
}

func LoadRoutes(r *gin.Engine) {

	r.GET("/", runSearch)

}

func init() {
	err := godotenv.Load()
	if err != nil {
		panicWrapper("Error loading .env file", http.StatusNotFound)
	}

	r := gin.Default()
	LoadRoutes(r)

	ginLambda = ginadapter.New(r)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return ginLambda.ProxyWithContext(ctx, req)

}

func main() {

	lambda.Start(Handler)

}

func returnError(c *gin.Context) {
	if err := recover(); err != nil {
		c.JSON(answerCode, gin.H{
			"status":  "error",
			"message": err,
		})
	}
}

func panicWrapper(err string, code int) {
	answerCode = code
	panic(err)
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}
