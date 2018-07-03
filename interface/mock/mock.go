package mock

type Retriever struct{
    Contents string
}

func ( this Retriever )Get( url string )string{
    return this.Contents
}
