package stncupload

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"stncCms/app/domain/helpers/stnccollection"
	"strings"

	"github.com/minio/minio-go/v6"
)

/*
	filename, uploadError = stncupload.NewFileUpload().UploadFile(filenameForm, c.PostForm("Resim2"))

	_, filetype := stncupload.NewFileUpload().RealImage("public/upl/kiosk/" + filename)
*/

/*

kullanım 2
	upl := stncupload.FileUpload{UploadPath: "public/upl/kiosk2/"}
	filename, uploadError = upl.UploadFile(filenameForm, c.PostForm("Resim2"))
*/
/*
//media diye bi tablo olsun oraya gitsin her resim silincekse oraya bi flag atsın
// TODO: bu kısım veritabanına gitsin daha sonra silsin gibi bişey olacak
// yada bunun yerine veritabanına flag atabilir
// errFile := os.Remove(uploadFilePath + deleteFilename)
// if errFile != nil {
// 	fmt.Println("errFile.Error()")
// 	fmt.Println(errFile.Error())
// 	errorReturn = errFile.Error()
// 	return filename, errorReturn
// }
*/

//buradan eşiebilmesi için func (fu FileUpload) Uplo bolyle olmalı fonksiyon
//NewFileUpload  construct s
func NewFileUpload() *FileUpload {
	return &FileUpload{}
}

type FileUpload struct {
	UploadPath string
	UploadSize int
	Types      []string
	MaxFiles   int
}

//UploadFileInterface interface
type UploadFileInterface interface {
	UploadFileForMinio(file *multipart.FileHeader) (string, error)
	UploadFile(filest *multipart.FileHeader) (string, string)
	MultipleUploadFile(filest []*multipart.FileHeader, originalName string)
	RealFileType(fileName string) (bool, string)
}

//So what is exposed is Uploader
var _ UploadFileInterface = &FileUpload{}

//TODO: https://github.com/gin-gonic/examples/tree/master/upload-file upload ornekleri var
//TODO: gerçek resim dosayasını tespit eden fonksiyon başka yere alınablir
//TODO: boyutlandırma https://github.com/disintegration/imaging
//https://socketloop.com/tutorials/golang-how-to-verify-uploaded-file-is-image-or-allowed-file-types
//https://www.golangprograms.com/how-to-get-dimensions-of-an-image-jpg-jpeg-png-or-gif.html
//UploadFile standart upload
func (fu FileUpload) UploadFile(filest *multipart.FileHeader) (filename string, errorReturn string) {
	var uploadFilePath string = fu.UploadPath

	// var deleteFilename string
	// var filename string
	// var errorReturn string

	if filest != nil {
		f, err := filest.Open()
		defer f.Close()
		if err != nil {
			errorReturn = err.Error()
		}

		if filest.Header != nil {

			size := filest.Size
			// var size2 = strconv.FormatUint(uint64(size), 10)
			if size > int64(1024000*fu.UploadSize) { // 1 MB
				uploadSizeStr := stnccollection.IntToString(fu.UploadSize)

				errorReturn = "HATA: Resim boyutu çok yüksek maximum " + uploadSizeStr + " MB olmalıdır" //+ size2
				filename = "false"
			}

			filename = newFileNameFunc(filest.Filename)
			// deleteFilename = filename

			fmt.Println(filename)

			out, err := os.Create(uploadFilePath + filename)

			defer out.Close()

			if err != nil {
				log.Fatal(err)
				errorReturn = err.Error()
				filename = "false"
			}

			_, err = io.Copy(out, f)

			if err != nil {
				log.Fatal(err)
				errorReturn = err.Error()
				filename = "false"
			}

		}
	}
	return filename, errorReturn
}
func (fu FileUpload) RealFileType(fileName string) (bool, string) {
	var uploadFilePath string = fu.UploadPath
	var typeControlError bool = false
	// open the uploaded file
	file, err := os.Open(uploadFilePath + fileName)
	defer file.Close()
	if err != nil {
		//TODO: buraya log koymak gerekiyor
		fmt.Println(err)
		err.Error()
		// os.Exit(1)
	}

	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)

	if err != nil {
		fmt.Println(err)
		err.Error()
		// os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	fmt.Println(fu.Types)
	_, typeControlError = stnccollection.FindSlice(fu.Types, filetype)

	fmt.Println(filetype)
	fmt.Println("typeControlError")
	fmt.Println(typeControlError)
	return typeControlError, filetype
}

func (fu FileUpload) MultipleUploadFile(files []*multipart.FileHeader, originalName string) {
	var uploadFilePath string = "public/upl/kiosk/"

	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		fmt.Println(files[i].Filename)
		defer file.Close()
		if err != nil {
			// fmt.Fprintln(w, err)
			return
		}

		out, err := os.Create(uploadFilePath + files[i].Filename)

		defer out.Close()
		if err != nil {
			// fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			// fmt.Fprintln(w, err)
			return
		}

		fmt.Println("Files uploaded successfully : ")
		fmt.Println(files[i].Filename + "\n")

	}

}

//	"github.com/minio/minio-go/v6"
func (fu FileUpload) UploadFileForMinio(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", errors.New("cannot open file")
	}
	defer f.Close()

	size := file.Size
	//The image should not be more than 500KB
	fmt.Println("the size: ", size)
	if size > int64(512000) {
		return "", errors.New("sorry, please upload an Image of 500KB or less")
	}
	//only the first 512 bytes are used to sniff the content type of a file,
	//so, so no need to read the entire bytes of a file.
	buffer := make([]byte, size)
	f.Read(buffer)
	fileType := http.DetectContentType(buffer)
	//if the image is valid
	if !strings.HasPrefix(fileType, "image") {
		return "", errors.New("please upload a valid image")
	}
	filePath := FormatFile(file.Filename)

	accessKey := os.Getenv("DO_SPACES_KEY")
	secKey := os.Getenv("DO_SPACES_SECRET")
	endpoint := os.Getenv("DO_SPACES_ENDPOINT")
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	if err != nil {
		log.Fatal(err)
	}
	fileBytes := bytes.NewReader(buffer)
	cacheControl := "max-age=31536000"
	// make it public
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	n, err := client.PutObject("chodapi", filePath, fileBytes, size, minio.PutObjectOptions{ContentType: fileType, CacheControl: cacheControl, UserMetadata: userMetaData})
	if err != nil {
		fmt.Println("the error", err)
		return "", errors.New("something went wrong")
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return filePath, nil
}
