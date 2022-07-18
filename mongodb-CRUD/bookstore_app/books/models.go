package books

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/markojudas/web-development/mongodb-CRUD/bookstore_app/config"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {

	// add ID and tags if you need them
	// ID     bson.ObjectId // `json:"id" bson:"_id"`
	Isbn   string  // `json:"isbn" bson:"isbn"`
	Title  string  // `json:"title" bson:"title"`
	Author string  // `json:"author" bson:"author"`
	Price  float64 // `json:"price" bson:"price,truncate"`
}

func AllBooks() ([]Book, error) {
	bks := []Book{}

	curr, err := config.Books.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer curr.Close(context.Background())

	if err = curr.All(context.TODO(), &bks); err != nil {
		return nil, err
	}

	if err := curr.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}

func OneBook(r *http.Request) (Book, error) {
	bk := Book{}

	isbn := r.FormValue("isbn")

	if isbn == "" {
		return bk, errors.New("400. Bad request")
	}

	curr, err := config.Books.Find(context.Background(), bson.M{"isbn": isbn})
	if err != nil {
		return bk, err
	}
	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		err := curr.Decode(&bk)
		if err != nil {
			return bk, err
		}
	}

	if err := curr.Err(); err != nil {
		return Book{}, err
	}

	return bk, nil
}

func PutBook(r *http.Request) (Book, error) {

	//get form values
	bk := Book{}

	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. All fields must be completed")
	}

	//convert from values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}
	bk.Price = f64

	//insert values
	_, err = config.Books.InsertOne(context.Background(), bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}

	return bk, nil
}

func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Enter number for price")
	}
	bk.Price = f64

	// update values
	//err = config.Books.Update(bson.M{"isbn": bk.Isbn}, &bk)

	filter := bson.M{"isbn": bk.Isbn}
	_, err = config.Books.UpdateOne(context.TODO(), filter, bk)

	if err != nil {
		return Book{}, err
	}

	return bk, nil
}

func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	//err := config.Books.Remove(bson.M{"isbn": isbn})
	_, err := config.Books.DeleteOne(context.TODO(), bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
