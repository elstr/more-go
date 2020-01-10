package main

import (
	"fmt"
	"net/http"
	"image"
	"image/jpeg" 
)

// Where we'll save the images uploaded by the user
var images = make(map[string]image.Image)


// HandleRoot handles requests to / and renders upload button
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<html>
		<body>
			<form method="post"enctype="multipart/form-data" action="/upload" name="upload">
			<labe for="file">Choose img:</label>
			<input type="file" name="image"/> <br/>
			<input type="submit" name="submit" value="Upload"/>
		</form>
		</body>
		</html>
	`)
}

// HandleUpload saves the image uploaded in the images map
func HandleUpload(w http.ResponseWriter, r *http.Request) {
	file, header, _ := r.FormFile("image")
	image, _, _ := image.Decode(file)
	
	images[header.Filename] = image
	http.Redirect(w, r, "/image?name=" + header.Filename, 303)
}

// HandleRenderImage renders the image the user uploaded
func HandleRenderImage( w http.ResponseWriter, r *http.Request){
	imageName := r.FormValue("name")
	image := images[imageName]
	jpeg.Encode(w, image, &jpeg.Options{ Quality: jpeg.DefaultQuality })
}

func main() {
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/upload", HandleUpload)
	http.HandleFunc("/image", HandleRenderImage)
	http.ListenAndServe(":8000", nil)
}
