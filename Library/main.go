package main

import (
	// "errors"
	"bufio"
	"fmt"
	"os"
	"strings"
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
	reader, ok := li.Readers[readerID]
	if !ok{
		return fmt.Errorf("читателя с ID %d не существует", readerID)
	}
	book, ok := li.Books[bookID]
	if !ok{
		return fmt.Errorf("книги с ID %d не существует", bookID)
	} 
	if !book.Access{
		return fmt.Errorf("книга с ID %d уже выдана", bookID)
	}
	// Добавляем книгу читателю
	reader = li.Readers[readerID]

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
// Функция возврата книги
func (li *Libary) BookReturn(bookID int, readerID int) error {
	book, exists := li.Books[bookID]
	if !exists{
		return fmt.Errorf("книги с ID %d не существует", bookID)
	}

	if book.Access{
		return fmt.Errorf("книга с ID %d уже у нас", bookID)
	}

	reader, ok := li.Readers[readerID]
	if !ok{
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
// Функция, которая показывает книги у читателя
func (li *Libary) AvailabilityReader(readerID int) error {

	reader, ok := li.Readers[readerID]
	if !ok{
		return fmt.Errorf("читателя с ID %d не существует", readerID)
	}

	fmt.Printf("Книги у читателя %s:\n", reader.Name)
	for id, name := range reader.Books {
        fmt.Printf("- ID: %d, Название: %s\n", id, name)
    }

	return nil
}


func main(){
	library := &Libary{
		Books: make(map[int]Book),
		Readers: make(map[int]Reader),
	}

	ans := 1
	scanner := bufio.NewScanner(os.Stdin)

	for ans != 0{
		fmt.Print("Выберите действие: \n 1-Добавить книгу \n 2-Добавить читателя \n 3-Выдать книгу читателю \n 4-Показать книги у читателя \n 5-Вернуть книгу \n 0-Выйти \n")
		fmt.Scan(&ans)

		switch ans{
		case 1: 
			var bookID int
			// var bookName string
			fmt.Println(library.Books)
			fmt.Println("Введите ID: ")
			fmt.Scan(&bookID)
			scanner.Scan()
			fmt.Println("Введите название книги")
			scanner.Scan()
			bookName := strings.TrimSpace(scanner.Text())
			
			err := library.AddBook(Book{bookID, bookName, true})
			if err != nil{
				fmt.Println("Ошибка: ", err)
			}else{fmt.Println("Книга успешно добавлена")}

		case 2:
			var readerID int
			
			fmt.Println(library.Readers)
			fmt.Println("Введите ID: ")
			fmt.Scan(&readerID)
			scanner.Scan()
			fmt.Println("Введите имя читателя: ")
			scanner.Scan()
			readerName := strings.TrimSpace(scanner.Text())

			err := library.AddReader(Reader{readerID, readerName, make(map[int]string)})
			if err != nil{
				fmt.Println("Ошибка: ", err)
			}else{
				fmt.Println("Читатель успешно добавлен")
			}

		case 3:
			var readerID int
			var bookID int
			fmt.Println(library.Readers)
			fmt.Println("Введите ID читателя: ")
			fmt.Scan(&readerID)
			fmt.Println(library.Books)
			fmt.Println("Введите ID книги: ")
			fmt.Scan(&bookID)

			err := library.DistributionBook(bookID, readerID)
			if err != nil{
				fmt.Println("Ошибка: ", err)
			}else{fmt.Println("Успешная выдача")}

		case 4:
			var readerID int
			fmt.Println("Введите ID читателя: ")
			fmt.Scan(&readerID)

			err := library.AvailabilityReader(readerID)
			if err != nil{
				fmt.Println("Ошибка: ", err)
			}
		
		case 5:
			var readerID, bookID int
			fmt.Println(library.Readers)
			fmt.Println("Введите ID читателя: ")
			fmt.Scan(&readerID)
			library.AvailabilityReader(readerID)
			fmt.Println("Введите ID книги: ")
			fmt.Scan(&bookID)

			err := library.BookReturn(bookID, readerID)
			if err != nil{
				fmt.Println("Ошибка: ", err)
			}
		case 6:
			fmt.Println(library.Books)	
		case 7:
			fmt.Println(library.Readers)
		}


	}

}





