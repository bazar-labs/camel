package api

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (h *handler) Upload(c *fiber.Ctx) error {
	// Get the file
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Generate temporary file
	tempFile := utils.UUIDv4() + "_" + file.Filename
	if err := c.SaveFile(file, tempFile); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Prepare the request to Pinata
	client := resty.New()
	resp, err := client.R().
		SetFile("file", tempFile).
		SetHeader("Authorization", "Bearer "+h.cfg.Pinata.JWT).
		Post("https://api.pinata.cloud/pinning/pinFileToIPFS")

	// Delete temporary file
	if err := os.Remove(tempFile); err != nil {
		log.Println("Warning: couldn't delete temporary file,", err)
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if resp.StatusCode() != 200 {
		return c.Status(resp.StatusCode()).SendString(resp.String())
	}

	return c.JSON(fiber.Map{
		"success": true,
		"ipfs":    resp.String(),
	})
}
