package x

import (
	

	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"strings"

	_ "github.com/lib/pq"
)


type Data struct {
	
	Text string `gorm:"column:text"`
}

func ReadSQl() ([]string, error) {
	dsn := "host=golang-db-1 user=postgres password=postgresql dbname=CRUD port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
  var data []string // Slice to store retrieved text values
	tx := db.Raw("select text from data").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return data, nil


	// rows, err := db.Raw("SELECT * FROM data").Rows() // Replace "your_table" with your actual table name
  // if err != nil {
  //   return nil, err
  // }
  // defer rows.Close() // Close the rows after iterating

  // for rows.Next() {
  //   var text string
  //   err := rows.Scan(&text) // Scan the "texto" column into the string variable
  //   if err != nil {
  //     return nil, err
  //   }
  //   data = append(data, text) // Append the retrieved text to the slice
  // }

  // return data, nil
 


}







func Read() ([]string, error) {
	file, err := os.Open("/app/prueba.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	content := string(buffer)
	lines := strings.Split(content, "\n")

	var result []string
	result = append(result, lines...)

	for _, line := range result {
		fmt.Println(line)
	}

	return result, nil
}

func Write(lineNum string) []string {
	dsn := "host=golang-db-1 user=postgres password=postgresql dbname=CRUD port=5432 sslmode=disable"
	db,_ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	data := Data{Text: lineNum} // Replace "text" with the actual text to insert

	result := db.Create(&data) // Create a new record using the provided data
	if result.Error != nil {
		fmt.Println("NO INSERTO")
	}
	Content, err := ReadSQl()
	if err != nil {
		log.Fatal(err)
	}

	return Content

}

func DeleteAll() string {
	dsn := "host=golang-db-1 user=postgres password=postgresql dbname=CRUD port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
         fmt.Errorf("failed to connect to database: %w", err)
    }
    

    // Delete all records from the "data" table
    result := db.Model(&Data{}).Where("TRUE").Delete(&Data{})
    if result.Error != nil {
         fmt.Errorf("failed to delete all data: %w", result.Error)
    }
	text:="borrado"
 
    return text

	
}

func DeleteSingle(id int) error {
	dsn := "host=golang-db-1 user=postgres password=postgresql dbname=CRUD port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return fmt.Errorf("failed to connect to database: %w", err)
    }
    

    // Delete the record with the specified ID from the "data" table
    result := db.Model(&Data{}).Where("id = ?", id).Delete(&Data{}) // Pass the ID as an argument
    if result.Error != nil {			
        return fmt.Errorf("failed to delete record: %w", result.Error)
    }

    return nil
	
}

func UpdateSingle(id int, newText string) ([]string, error){
	dsn := "host=golang-db-1 user=postgres password=postgresql dbname=CRUD port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
         fmt.Errorf("failed to connect to database: %w", err)
    }

	result:= db.Model(&Data{}).Where("id = ?",id ).Update("text",newText)
	if result.Error != nil {			
         fmt.Errorf("failed to delete record: %w", result.Error)
    }


	Content, err := ReadSQl()
	if err != nil {
		log.Fatal(err)
	}

	return Content, err
}
