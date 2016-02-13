
# Cartogopher [![Build Status](https://travis-ci.org/LiterallyElvis/cartogopher.svg?branch=master)](https://travis-ci.org/LiterallyElvis/cartogopher) [![Go Report Card](https://goreportcard.com/badge/github.com/LiterallyElvis/cartogopher)](https://goreportcard.com/report/github.com/LiterallyElvis/cartogopher)

A super simple library to read CSVs as maps instead of arrays in golang

## Usage

Cartogopher can be used as an augmentor of Go's built in CSV reader.

The 'Hello World' for cartogopher looks something like this:

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

        myCSV, err := cartogopher.NewReader(reader)
        if err != nil{
            // handle error as you please
        }

        theRestOfTheFile, err := myCSV.Reader.ReadAll()
        if err != nil{
            // handle error as you please
        }
        // do things with theRestOfTheFile
    }
