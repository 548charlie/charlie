package main

import(  
    "strings" 
        "fmt" 
      ) 

func main(){
exts:= [] string{  "pdf","jpg","bmp","tiff","tif"}
    for _, ex := range exts {
        viewer := get_viewer(ex)
        fmt.Println("viewer", viewer ) 
    } 
}  


func get_viewer(ext string) string {
    exts := []string{"tiff","bmp","jpg" }
    ext = strings.ToLower(ext) 
    contain := strings.Contains(ext, "pdf")
    var viewer string = "junk" 
    if contain == true {
        viewer =  "Adobe Acrobat" 
    }  else {
        for _, ex := range exts {
            contain = strings.Contains(ext, ex) 
            if contain == true {
                viewer= "image" 
            } 
        } 
    }
    return viewer
} 

