package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"git.sr.ht/~jackmordaunt/go-libwebp/webp"

	"github.com/gin-gonic/gin"
	gim "github.com/ozankasikci/go-image-merge"
)

func gridHandler(c *gin.Context, mediaImages []ProductImage, ID string) (string, error) {
	gridFileName := filepath.Join("static", ID+".webp")

	if _, err := os.Stat(gridFileName); err == nil {
		return gridFileName, nil
	}

	images := make([]string, len(mediaImages))

	dirname, err := os.MkdirTemp("static", ID+"*")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(dirname)

	for i, media := range mediaImages {
		req, err := http.NewRequest("GET", media.Image.URI, http.NoBody)
		if err != nil {
			return "", err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()

		fileName := filepath.Join(dirname, strconv.Itoa(i)+".jpg")
		file, err := os.Create(fileName)
		if err != nil {
			return "", err
		}

		_, err = file.ReadFrom(res.Body)
		if err != nil {
			return "", err
		}

		images[i] = fileName

	}

	var gridImage []*gim.Grid
	for _, image := range images {
		gridImage = append(gridImage, &gim.Grid{
			ImageFilePath: image,
		})
	}

	grid, err := gim.New(gridImage, 2, 2).Merge()
	if err != nil {
		return "", err
	}

	f, err := os.Create(gridFileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if err := webp.Encode(f, grid, nil); err != nil {
		return "", err
	}

	c.File(gridFileName)
	return gridFileName, nil
}
