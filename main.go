package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type quote struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var quotes = []quote{
	{ID: "1", Title: "Нащо я маю це слухати", Author: "Dima"},
	{ID: "2", Title: "Хіба я на таке здатний", Author: "Mykola"},
	{ID: "3", Title: "Я тільки по-любві", Author: "Vlad"},
	{ID: "4", Title: "В одній мірі", Author: "Dima"},
	{ID: "5", Title: "А ти пробував?", Author: "Dima"},
	{ID: "6", Title: "Як та машина, постав з ключами заведену посеред дороги і ніхто не візьме", Author: "Mykola"},
	{ID: "7", Title: "Хоч може льоха якого перепишуть", Author: "Mykola"},
	{ID: "8", Title: "Перепрошую, шановний пане, чи можна будь ласка до вас звернутись", Author: "Vlad"},
	{ID: "9", Title: "Додаткова мотивація", Author: "Dima"},
	{ID: "10", Title: "Я вже на собі хрест поставив", Author: "Vlad"},
	{ID: "11", Title: "Якщо мене ще можна назвати людиною", Author: "Vlad"},
	{ID: "12", Title: "Найрозумніший в межах сто метрів", Author: "Vlad"},
	{ID: "13", Title: "Люди - це просто шматки біологічного м'яса", Author: "Dima"},
}

func main() {
	router := gin.Default()
	router.GET("/quotes", getQuotes)
	router.GET("/quotes/:author", getQuoteByAuthor)
	router.POST("/quotes", postQuotes)
	router.Run("localhost:8080")
}

func getQuotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quotes)
}

func postQuotes(c *gin.Context) {
	var newQuote quote

	// Call BindJSON to bind the received JSON to
	// newQuote.
	if err := c.BindJSON(&newQuote); err != nil {
		return
	}

	// Add the new quote to the slice.
	quotes = append(quotes, newQuote)
	c.IndentedJSON(http.StatusCreated, newQuote)
}

func getQuoteByAuthor(c *gin.Context) {
	author := c.Param("author")

	var filteredQuotes []quote

	// Loop over the list of quotes, looking for
	// an author whose name matches the parameter.
	// add matched author to the filteredArray to filter by author
	for _, q := range quotes {
		if q.Author == author {
			filteredQuotes = append(filteredQuotes, q)
		}
	}
	if len(filteredQuotes) != 0 {
		c.IndentedJSON(http.StatusOK, filteredQuotes[rand.Intn(len(filteredQuotes))])
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Author not found"})
}
