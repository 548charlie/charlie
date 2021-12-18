package main

import (
    "bytes"
    "encoding/csv"
    "encoding/xml"
    "os"
)

var XML = `<?xml version="1.0"?>
<resultset statement="select * from mysql.user
" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <row>
        <field name="Host">localhost</field>
        <field name="User">root</field>
        <field name="Password"></field>
        <field name="Select_priv">Y</field>
        <field name="Insert_priv">Y</field>
        <field name="Update_priv">Y</field>
        <field name="Delete_priv">Y</field>
        <field name="Create_priv">Y</field>
        <field name="Drop_priv">Y</field>
        <field name="Reload_priv">Y</field>       
  </row>
  <row>
        <field name="Host">localhost2</field>
        <field name="User">root2</field>
        <field name="Password"></field>
        <field name="Select_priv">Y</field>
        <field name="Insert_priv">N</field>
        <field name="Update_priv">Y</field>
        <field name="Delete_priv">Y</field>
        <field name="Create_priv">Y</field>
        <field name="Drop_priv">Y</field>
        <field name="Reload_priv">Y</field>       
</row>
</resultset>`

type Row struct {
    Fields []struct {
        Name    string    `xml:"name,attr"`
        Value    string    `xml:",chardata"`
    } `xml:"field"`
}

func main() {
    b := bytes.NewBufferString(XML)
    decoder := xml.NewDecoder(b)
    out := csv.NewWriter(os.Stdout)
    var err error

    row := Row{}
    var record []string
    writeHeader := true
    for token, _ := decoder.Token(); err == nil; token, err = decoder.Token() {
        if start, ok := token.(xml.StartElement); ok && start.Name.Local == "row" {
            decoder.DecodeElement(&row, &start)
            if cap(record) < len(row.Fields) {
                record = make([]string, len(row.Fields))
            }
            record = record[:len(row.Fields)]
            if writeHeader {
                for i := range row.Fields {
                    record[i] = row.Fields[i].Name
                }
                out.Write(record)
                writeHeader = false
            }
            for i := range row.Fields {
                record[i] = row.Fields[i].Value
            }
            out.Write(record)
            row.Fields = row.Fields[:0]
        }
    }
    out.Flush()
}

