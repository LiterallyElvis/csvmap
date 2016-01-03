# Cartogopher
A super small library to read CSVs as maps instead of arrays in golang

## Usage
Cartogopher can be used one of two ways: Either as a wrapper around or an augmentor of Go's built in CSV reader.

For instance, you could open your files like this:

    package main
    
    import "github.com/literallyelvis/cartogopher"
    
    func main(){
        myCSV, err := cartogopher.NewFromFilePath("whatever.csv")
        if err != nil{
            // handle error as you please
        }
        
        theRestOfTheFile, err := myCSV.Reader.ReadAll()
        if err != nil{
            // handle error as you please
        }
        
        // do things with theRestOfTheFile
    }

Alternatively, you could use your own reader if you so choose, like so: 

    package main
    
    import (
        "encoding/csv"
        "github.com/literallyelvis/cartogopher"
        "os"
    )
    
    func main(){
        file, err := os.Open("whatever.csv")
        if err != nil{
            // handle error as you please
        }
        
        reader, err := csv.NewReader(file)
        if err != nil{
            // handle error as you please
        }
        
        myCSV, err := cartogopher.NewFromReader(reader)
        if err != nil{
            // handle error as you please
        }
        
        theRestOfTheFile, err := myCSV.Reader.ReadAll()
        if err != nil{
            // handle error as you please
        }
        // do things with theRestOfTheFile
    }
