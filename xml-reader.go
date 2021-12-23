package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
    "time"
)

type Statement struct{
    XMLName xml.Name `xml:"Statement"`
    Data *struct{
        Operations []struct{
            PayDoc *struct{
                Id string `xml:"id,attr"`
                Details *struct {
                    Date docTime `xml:"DocDate"`
                    Sum float32 `xml:"Sum"`
                    Payer *struct{
                        Name string `xml:"Name"`
                        Inn string `xml:"INN"`
                    } `xml:"Payer"`
                } `xml:"PayDocRu"`
            } `xml:"PayDoc"`
            DC int `xml:"DC"`
        } `xml:"OperationInfo"`
    } `xml:"Data"`
}

type docTime struct{
    time.Time
}

func (t *docTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var v string
    err := d.DecodeElement(&v, &start)
    if err != nil {
        return err
    }
    parse, err := time.Parse("2006-01-02", v)
    if err != nil {
        return err
    }
    *t = docTime{parse}

    return nil
}

func main() {
    xmlFile, err := os.Open("./resources/payments.xml")
    if err != nil {
        fmt.Println(err)
    }
    defer xmlFile.Close()

    byteValue, err := ioutil.ReadAll(xmlFile)
    if err != nil {
        fmt.Println(err)
    }

    var statement Statement
    err = xml.Unmarshal(byteValue, &statement)
    if err != nil {
        fmt.Println(err)
    }

    for _, operation := range statement.Data.Operations {
        if operation.PayDoc.Details != nil {
            fmt.Println(fmt.Sprintf("id: %s; date: %s; sum: %f; name: %s; inn: %s; dc: %d",
                operation.PayDoc.Id,
                operation.PayDoc.Details.Date,
                operation.PayDoc.Details.Sum,
                operation.PayDoc.Details.Payer.Name,
                operation.PayDoc.Details.Payer.Inn,
                operation.DC,
            ))
        }
    }
}
