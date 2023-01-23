package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadImage(image string) (string, error) {

	cld, err := cloudinary.NewFromParams("dhohircloud", "259886355277397", "KCnW5OfaEX8_3oNu3k6YVW7cpDg")

	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	var ctx = context.Background()

	uploadResult, err := cld.Upload.Upload(
		ctx,
		"https://res.cloudinary.com/dhohircloud/image/upload/music_app/" + image,
		uploader.UploadParams{PublicID: "music"})

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	fmt.Println(uploadResult.SecureURL)

	return uploadResult.SecureURL, nil
}
