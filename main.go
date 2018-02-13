package main

import (
	"fmt"
	"image"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	ascii "github.com/joyrex2001/aws-lambda-ascii/asciiart"
)

// Request describes the input json structure of our lamda function.
type Request struct {
	Url      string `json:"url"`
	Width    int    `json:"width"`
	Gradient string `json:"gradient"`
}

// Response describes the returned json structure of our lambda function.
type Response struct {
	AsciiArt []string `json:"asciiart"`
}

// downloadAsImage will download the resource from the given url, and will
// convert this to an image.Image. It will return an error if failed.
func downloadAsImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// HandleRequest will take json messages as specified by the 'Request' struct.
// It will validate the request, and if valid, it will download the image
// specified by the url and resize convert this image to an ascii art image
// for the requested width.
func HandleRequest(r Request) (*Response, error) {
	if r.Width < 0 || r.Width > 160 {
		return nil, fmt.Errorf("invalid width: %d", r.Width)
	}
	if r.Url == "" {
		return nil, fmt.Errorf("missing url")
	}
	if r.Gradient == "" {
		r.Gradient = "@8QOECLIeoc+:-. "
	}

	img, err := downloadAsImage(r.Url)
	if err != nil {
		return nil, err
	}

	res := &Response{
		AsciiArt: ascii.Convert(img, r.Width, r.Gradient),
	}

	return res, nil
}

// main is the entrypoint for this lambda function and will start the aws lambda
// listener. Requests will be dispatched to the HandleRequest method.
func main() {
	lambda.Start(HandleRequest)
}
