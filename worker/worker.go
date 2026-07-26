package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	app := fiber.New()
	api := app.Group("/api/v1")

	api.Post("/comment", createComment)

	app.Listen(":3000")
}

func createComment(c *fiber.Ctx) error {

	cmt := new(comment)
	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})

		return err
	}

	cmtInBytes, err := json.Marshal(cmt)
	PushCommentToQueue("comments", cmtInBytes)

	c.JSON(&fiber.Map{
		"success": true,
		"message": "Comment Pushed Successfully",
		"comment": cmt,
	})

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return err
	}

	return err
}

func PushCommentToQueue(topic string, message []byte) {
	brokerUrl := []string{"localhost:29092"}
	producer, err := ConnectProducer(brokerUrl)

}

func ConnectProducer(brokerUrl []string) (srama.SyncProducer, error) {
	config := saram.NewConfig()
	config.Producer.Return.Successes = true

	config.Producer.RequiredAcks = sarama.WaitForAll

	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokerUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
