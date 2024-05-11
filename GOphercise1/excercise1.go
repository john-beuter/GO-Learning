package main


//Need to read in values from a CSV and make it a game.

import(
  "encoding/csv" 
  "fmt"
  "io"
  "log"
  "os"
)

func main() {

  f, err := os.Open("problems.csv")
  if err != nil{
    log.Fatal(err)
  }


  defer f.Close()

  csvReader := csv.NewReader(f)
  for{
    rec, err := csvReader.Read()
    if err == io.EOF{
      break
    }

    if err != nil{
       log.Fatal(err)
    }
    
    fmt.Println("%v", rec[0])
    fmt.Println("Enter you answer: ")

    //Need to add a reading to get the users input
    //then compare that input ot rec[1] if it matches have points


  }


}