package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

func ItermSaver() chan interface{}{
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for  {
			item := <- out
			itemCount++

			//log.Printf("item saver: got #%d:%v", itemCount,item)

			_, err := save(item)

			if err != nil {
				log.Printf("item error: %v", err)
			}

			//get()
		}


	}()


	return  out
}

func get()  {
	c, _ := elastic.NewClient(elastic.SetSniff(false))
	res, err := c.Get().
		Index("data_es_test").
		Type("user").
		Do(context.Background())

	if err != nil{
		log.Printf("err %s", err)
	} else {
		log.Printf("es source %v", res.Source)
		log.Printf("es id %+v", res.Id)
	}
}

func save(item interface{}) (id string,err error) {
	c, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil{
		return "", err
	}
	res, err := c.Index().Index("data_es_test").
		Type("user").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return res.Id, nil
}
