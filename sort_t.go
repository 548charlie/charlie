package main

import "fmt"
type IntSlice []int
func (i IntSlice) len() int {
    return len(i) 
} 
func (i IntSlice) less(a, b int) bool {
    if i[a] < i[b] {
        return true
    }
    return false
}
func (i IntSlice) swap(a, b int) {
    i[a],i[b] = i[b],i[a]    
}  
type Sortable interface {
    len() int
    less(i,j int) bool
    swap(i, j int)  
} 

func GenericBubbleSort(data Sortable) {
    fmt.Println("before sort :", data) 
    for i := 0; i < data.len() ; i++  {
        for j := 0; j < data.len()  -i - 1; j++ {
            if data.less(j+1, j)  {
                data.swap(j+1, j)     
            }   
        }  
    }
    fmt.Println("after sort :", data) 
}  

func main() {
    var array  IntSlice= []int {1,25,3,7,3,2,6,4,10} 
    GenericBubbleSort(array) 
}  
