package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	ListDir      = 0x0001
	UPLOAD_DIR   = "./upload"
	TEMPLATE_DIR = "./views"
)

var _templates = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	checkError(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != "html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template", templatePath)
		t := template.Must(template.ParseFiles(templateName))
		_templates[templateName] = t
	}
	// for _, tmpl := range []string{"upload", "list"} {
	// 	t := template.Must(template.ParseFiles(tmpl + "html"))
	// 	_templates[tmpl] = t
	// }
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	err := _templates[tmpl].Execute(w, locals)
	checkError(err)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "upload", nil)
		// if err := renderHtml(w, "upload", nil); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// t.Execute(w, nil)
		// io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
		// 	" enctype=\"multipart/format-data\">"+
		// 	"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
		// 	"<input type=\"submut\" value=\"Upload\" />"+
		// 	"</format>")

		// return
	}
	if r.Method == "POST" {
		// type(f), type(h), type(err)
		// multipart.File, *multipart.FileHeader, error
		f, h, err1 := r.FormFile("image")
		checkError(err1)
		filename := h.Filename
		defer f.Close()
		t, err2 := ioutil.TempFile(UPLOAD_DIR, filename)
		checkError(err2)
		defer t.Close()
		_, err3 := io.Copy(t, f)
		checkError(err3)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageID := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageID
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	checkError(err)

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "list", locals)
	// t, err := template.ParseFiles("list.html")
	// if err := renderHtml(w, "list", locals);err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// t.Execute(w, locals)
	// var listHtml string
	// for _, fileInfo := range fileInfoArr {
	// 	imgid := fileInfo.Name()
	// 	listHtml += "<li><a href=\"/view?id=" + imgid + "\">imgid</a></li>"
	// }

	// io.WriteString(w, "<ol>"+listHtml+"</ol>")
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	http.HandleFunc("/", safeHandler(listHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
