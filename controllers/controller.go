package controllers

import (
	"app/config"
	"app/models"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TestEndPoint(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func GetQuotes(context *gin.Context) {
	pageStr := context.DefaultQuery("page", "1")
	limitStr := context.DefaultQuery("limit", "7")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Id ma boi"})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Id ma boi"})
		return

	}

	offset := (page - 1) * limit
	row, err := config.DB.Query("SELECT id, message, author FROM quotes LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		log.Fatal(err)
	}

	var quotes []models.Quote

	for row.Next() {
		var q models.Quote
		err := row.Scan(&q.ID, &q.Message, &q.Author)
		if err != nil {
			log.Fatal(err)
		}

		quotes = append(quotes, q)

	}
	context.IndentedJSON(http.StatusOK, quotes)

}

func GetQuoteById(context *gin.Context) {
	idStr := context.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {

		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Id ma boi"})
		return

	}

	var quote models.Quote

	err = config.DB.QueryRow("SELECT id, message, author FROM quotes WHERE id = $1", id).Scan(&quote.ID, &quote.Message, &quote.Author)
	if err != nil {
		if err == sql.ErrNoRows {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Quote Not Found"})
			return
		}
	}
	context.IndentedJSON(http.StatusOK, quote)
}

func GetRandomQuote(context *gin.Context) {
	var quote models.Quote
	err := config.DB.QueryRow("SELECT id, message, author FROM quotes ORDER BY RANDOM() LIMIT 1").Scan(&quote.ID, &quote.Message, &quote.Author)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Id ma boi"})
		return
	}
	context.IndentedJSON(http.StatusOK, quote)

}

func AddQuote(context *gin.Context) {

	var newQuote models.Quote
	err := context.BindJSON(&newQuote)
	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO quotes (message, author) VALUES ($1, $2) RETURNING id, message, author"
	var createQuotes models.Quote
	result := config.DB.QueryRow(query, newQuote.Message, newQuote.Author)
	result.Scan(&createQuotes.ID, &createQuotes.Message, &createQuotes.Author)
	context.IndentedJSON(http.StatusCreated, createQuotes)

}

func UpdateQuote(context *gin.Context) {
	var updatedQuote models.Quote
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Id ma boi"})
		return
	}

	err = context.BindJSON(&updatedQuote)
	if err != nil {
		log.Fatal(err)

	}

	query := `UPDATE quotes SET message=$1, author=$2 WHERE id=$3`

	_, err = config.DB.Exec(query, updatedQuote.Message, updatedQuote.Author, id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Id ma boi"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Quote updated successfully"})
}

func DeleteQuote(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request ma boi"})
		return
	}

	query := `DELETE FROM quotes WHERE id = $1`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Quote Not Found!"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Quote Deleted Successfully!"})

}

func SearchQuote(context *gin.Context) {
	searchQuery := context.Query("q")
	searchTerm := "%" + searchQuery + "%"

	query := `SELECT * FROM quotes WHERE message LIKE $1`
	row, err := config.DB.Query(query, searchTerm)
	if err != nil {
		log.Fatal(err)
	}

	var quotes []models.Quote
	for row.Next() {
		var quote models.Quote
		err := row.Scan(&quote.ID, &quote.Message, &quote.Author)
		if err != nil {
			log.Fatal(err)
		}
		quotes = append(quotes, quote)
	}

	if len(quotes) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No quotes found"})
		return
	}
	context.IndentedJSON(http.StatusOK, quotes)
}

func GetCommentsByQuoteId(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Id ma boi"})
		return
	}

	query := `SELECT id, user_id, comment FROM comments WHERE user_id = $1`
	row, err := config.DB.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}

	var comments []models.Comment
	for row.Next() {
		var comment models.Comment
		err := row.Scan(&comment.ID, &comment.User_id, &comment.Comment)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}

	if len(comments) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No comments yet, be the first to post one"})
		return
	}
	context.IndentedJSON(http.StatusOK, comments)
}

func AddCommentsById(context *gin.Context) {
	quoteId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Id ma boi"})
		return
	}
	var newComment models.Comment
	err = context.BindJSON(&newComment)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Request ma boi"})
		return
	}

	query := `INSERT INTO comments(user_id, comment) VALUES($1, $2) RETURNING id, user_id, comment`

	var createdComment models.Comment

	err = config.DB.QueryRow(query, quoteId, newComment.Comment).Scan(&createdComment.ID, &createdComment.User_id, &createdComment.Comment)
	if err != nil {
		log.Fatal(err)
	}
	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Comment added with success"})
}
