package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/adminsemy/URLShorting/internal/shorten"
	"github.com/adminsemy/URLShorting/internal/storage"
)

func main() {
	svc := shorten.NewService(storage.NewInMemory())
	input := model.ShortenInput{RawURL: "https://google.com"}
	shortening, err := svc.Shorten(context.Background(), input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(shortening)
}
