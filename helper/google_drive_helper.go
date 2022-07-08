package helper

import (
	"log"

	"google.golang.org/api/drive/v3"
)

type googleDriveHelper struct {
	googleDriveService *drive.Service
}

func NewGoogleDriveHelper(googleDriveService *drive.Service) *googleDriveHelper {
	return &googleDriveHelper{googleDriveService}
}

func (h *googleDriveHelper) GoogleDriveCreateFolder(ID string) (*drive.File, error) {
	data := &drive.File{
		Name:     ID,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{"1hGeCEuCcY2IlRvxasSIADZohuIo8gLY8"},
	}

	folder, err := h.googleDriveService.Files.Create(data).Do()
	if err != nil {
		log.Println("Could not create dir: " + err.Error())
		return nil, err
	}
	return folder, nil
}
