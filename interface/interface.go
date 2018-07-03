package main

import "fmt"
import "./mock"

type Retiriever interface {
    Get(url string)string
}

func download( r Retiriever )string{
    return r.Get("www.imooc.com")
}

func main(){
    var r Retiriever
    r = mock.Retriever{"this is a fake imooc.com"}
    fmt.Println(download(r))
}


