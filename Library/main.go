package main

import (
	// "errors"
	"fmt"
)

// Структура представляет книгу
type Book struct{
	ID int
	Name string
	Access bool
}
// Структура представляет читателя
type Reader struct{
	ID int
	Name string
	Books map[int]string
}
// Структура представляет библиотеку
type Libary struct{
	Books map[int]Book
	Readers map[int]Reader
}

// Интерфейс представляет библиотеку и его функционал
type Libaryint interface{
	AddReader(reader Reader) error
	AddBook(book Book) error
	DistributionBook(book Book) error
	BookReturn(book Book) error
	AvailabilityReader(reader Reader)
}

// Функция добавления книги
func (li *Libary) AddBook(book Book) error{
	if _,ok := li.Books[book.ID]; ok {
		return fmt.Errorf("книга с ID %d уже существует в библиотеке", book.ID)
	}
	li.Books[book.ID] = book

	return nil
}

// Функция добавления читателя
func (li *Libary) AddReader(reader Reader) error{
	if _, ok := li.Readers[reader.ID]; ok {
		return fmt.Errorf("читатель с ID %d уже существует в библиотеке", reader.ID)
	}
	li.Readers[reader.ID] = reader
	return nil
}

//Функция выдачи книги читателю 
func (li *Libary) DistributionBook(bookID int, readerID int) error{
	book, exists := li.Books[bookID]
	if !exists{
		return fmt.Errorf("книги с ID %d не существует", bookID)
	} 
	if !book.Access{
		return fmt.Errorf("книга с ID %d уже выдана", bookID)
	}
	// Добавляем книгу читателю
	reader := li.Readers[readerID]

	if reader.Books == nil{
		reader.Books =make(map[int]string)
	}
	reader.Books[bookID] = book.Name
	li.Readers[readerID] = reader

	// Пометка книги, что выдали
	book.Access = false
	li.Books[bookID] = book
	return nil	
}

func (li *Libary) BookReturn(bookID int, readerID int) error {
	book, exist := li.Books[bookID]
	if !exist{
		return fmt.Errorf("книги с ID %d не существует", bookID)
	}

	if book.Access{
		return fmt.Errorf("книга с ID %d уже у нас", bookID)
	}

	reader, exist := li.Readers[readerID]
	if !exist{
		return fmt.Errorf("читателя с ID %d не существует", readerID)
	}


	if _, hasbook := reader.Books[bookID]; !hasbook{
		return fmt.Errorf("книги с ID %d у читателя нет", bookID)
	}

	delete(reader.Books, bookID)
	li.Readers[readerID] = reader 

	book.Access = true
	li.Books[bookID] = book

	return nil
}

func (li *Libary) AvailabilityReader(readerID int) error {

	reader, exist := li.Readers[readerID]
	if !exist{
		return fmt.Errorf("читателя с ID %d не существует", readerID)
	}
	fmt.Println(reader.Books)

	return nil
}


func main(){
	library := &Libary{
		Books: make(map[int]Book),
		Readers: make(map[int]Reader),
	}

	err := library.AddBook(Book{1, "Ведьмак", true})
	if err != nil {
		fmt.Print("Ошибка при добавлении книги: ", err)
		return
	}

	user := library.AddReader(Reader{2, "Igor", make(map[int]string)})
	if user != nil {
		fmt.Print("Ошибка при добавлении читателя: ", err)
		return
	}

	library.DistributionBook(1, 2,)

	library.AvailabilityReader(2)

}





