package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"

	"github.com/go-resty/resty/v2"
)

const PINATA_URL = "https://api.pinata.cloud/pinning/pinFileToIPFS"

type PinataResponse struct {
	IPFSHash string `json:"IpfsHash"`
}

func (s *Service) UploadToIPFS(ctx context.Context, form *multipart.Form) (string, error) {
	// Create a new bytes.Buffer to hold the form data
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// iterate over the form data and add it to the writer
	for _, files := range form.File {
		files := files
		for _, file := range files {
			file := file

			// open the file
			openFile, err := file.Open()
			if err != nil {
				return "", err
			}
			defer openFile.Close()

			// create a new form file part
			// TODO use gameID + itemID as filename
			part, err := writer.CreateFormFile("file", "metadata/"+file.Filename)
			if err != nil {
				return "", err
			}

			// copy the file data into the form part
			_, err = io.Copy(part, openFile)
			if err != nil {
				return "", err
			}
		}
	}

	// close the writer
	err := writer.Close()
	if err != nil {
		return "", err
	}

	// Forward the request to Pinata
	resp, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+s.cfg.Pinata.JWT).
		SetHeader("Content-Type", writer.FormDataContentType()).
		SetBody(body.Bytes()).
		Post(PINATA_URL)

	if err != nil {
		return "", err
	}

	// unmarshal the response
	var pinataResponse PinataResponse
	err = json.Unmarshal(resp.Body(), &pinataResponse)
	if err != nil {
		return "", err
	}

	return pinataResponse.IPFSHash, nil
}
