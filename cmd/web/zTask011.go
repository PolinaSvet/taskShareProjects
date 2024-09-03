package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"tracker/cmd/task/templTask1"
	"tracker/cmd/task/templTask2"
)

type TypeTask011 struct {
	PageName     string
	PageDescribe string
	PageTask     string
	TaskCode     string
}

func processHandlerTask011Start(w http.ResponseWriter, r *http.Request) {
	channelNum := r.URL.Query().Get("channel")
	par1Str := r.URL.Query().Get("par1")
	par2Str := r.URL.Query().Get("par2")
	var par1, par2 int

	if i, err := strconv.Atoi(par1Str); err == nil {
		par1 = i
	}
	if i, err := strconv.Atoi(par2Str); err == nil {
		par2 = i
	}

	switch channelNum {
	case "1":
		go templTask1.StartTask(channelTask11, par1, par2) //100,5
	case "2":
		go templTask2.StartTask(channelTask11, par1, par2) //100,5
	}

	fmt.Printf("Start %v %v %v: \n", channelNum, par1, par2)
}

func processHandlerTask011RefreshData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == Task011URLData {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for {
			select {
			case <-r.Context().Done():
				fmt.Println("Client closed the connection")
				return
			case msg := <-channelTask11:
				fmt.Fprintf(w, "data: %s\n\n", msg)
				w.(http.Flusher).Flush()
			}
		}
	}
}

func ServeHTTPTask011(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == Task011URL {

		tmpl, err := template.ParseFiles(Task011Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask011{PageName: ReplaceTxtBr(ListTasksAll[14].Name, "; "),
			PageDescribe: ReplaceTxtBr(ListTasksAll[14].Describe, " "),
			PageTask:     ListTasksAll[14].Task,
			TaskCode:     ""}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

}
