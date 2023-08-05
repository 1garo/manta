package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Err struct {
    Message string `json:"message"`
}

func main() {
    app := fiber.New()

    app.Get("/players/:steam32ID", func(c *fiber.Ctx) error {
        log.Info("[GET /players/:steam32ID]")

        id := c.Params("steam32ID")
        playerID, err := strconv.Atoi(id)
        if err != nil {
            e := Err{
                Message: fmt.Sprintf("expect a int, received: %s", id),
            }
            return c.Status(fiber.StatusBadRequest).JSON(e)
        }

        resp, err := http.Get(fmt.Sprintf("https://api.opendota.com/api/players/%d", playerID))
        if err != nil {
            return c.Status(fiber.StatusBadRequest).SendString(err.Error())
        }

        if resp.StatusCode != 200 {
            err := Err {
                Message: fmt.Sprintf("could not find player with steam32ID: %d", playerID),
            }
            return c.Status(resp.StatusCode).JSON(err)
        }

        // TODO: chech if resp.body.profile exists, if not throw a error

        var p interface{}

        err = json.NewDecoder(resp.Body).Decode(&p)
        if err != nil {
            log.Fatal(err)
        }

        return c.JSON(p)
    })

    app.Get("/playersByRank", func(c *fiber.Ctx) error {
        log.Info("[GET /playersByRank]")
        resp, err := http.Get("https://api.opendota.com/api/playersByRank")
        if err != nil {
            log.Fatal(err)
        }
        var p interface{}

        err = json.NewDecoder(resp.Body).Decode(&p)
        if err != nil {
            log.Fatal(err)
        }

        return c.JSON(p)
    })

    app.Listen(":3000")
}
