package cmd

import (
	"fmt"
	"log"
	"net/http"
	"tracker/cmd/task/errorhandler"
)

const (
	HomePage     string = "ui/html/index.html"
	ErrorPage    string = "ui/html/error.html"
	TemplatePage string = "ui/html/templates.html"
	AboutPage    string = "ui/html/about.html"
	ContactsPage string = "ui/html/contacts.html"
	ListPage     string = "ui/html/list.html"

	ErrorLogPage      string = "ui/html/errorLog.html"
	Task_001_01_Page  string = "ui/html/z_task_001_01.html"
	Task_001_02_Page  string = "ui/html/z_task_001_02.html"
	Task_001_03_Page  string = "ui/html/z_task_001_03.html"
	Task_002_01_Page  string = "ui/html/z_task_002_01.html"
	Task_002_02_Page  string = "ui/html/z_task_002_02.html"
	Task_002_03_Page  string = "ui/html/z_task_002_03.html"
	Task_003_01_Page  string = "ui/html/z_task_003_01.html"
	Task_003_01_PList string = "ui/html/z_task_003_01_list.html"
	Task004Page       string = "ui/html/zTask004.html"
	Task004PList      string = "ui/html/zTask004list.html"
	Task005Page       string = "ui/html/zTask005.html"
	Task005PList      string = "ui/html/zTask005list.html"
	Task006Page       string = "ui/html/zTask006.html"
	Task006PList      string = "ui/html/zTask006list.html"
	Task006PTest      string = "ui/html/zTask006test.html"
	Task007Page       string = "ui/html/zTask007.html"
	Task007Tree       string = "ui/html/zTask007tree.html"
	Task008Page       string = "ui/html/zTask008.html"
	Task009Page       string = "ui/html/zTask009.html"
	Task009PList      string = "ui/html/zTask009list.html"
	Task010Page       string = "ui/html/zTask010.html"
	Task011Page       string = "ui/html/zTask011.html"

	ErrorLogURL     string = "/errorLog/"
	Task_001_01_URL string = "/z_task_001_01/"
	Task_001_02_URL string = "/z_task_001_02/"
	Task_001_03_URL string = "/z_task_001_03/"
	Task_002_01_URL string = "/z_task_002_01/"
	Task_002_02_URL string = "/z_task_002_02/"
	Task_002_03_URL string = "/z_task_002_03/"
	Task_003_01_URL string = "/z_task_003_01/"
	Task004URL      string = "/zTask004/"
	Task005URL      string = "/zTask005/"
	Task006URL      string = "/zTask006/"
	Task007URL      string = "/zTask007/"
	Task008URL      string = "/zTask008/"
	Task008URLData  string = "/data"
	Task009URL      string = "/zTask009/"
	Task010URL      string = "/zTask010/"
	Task010URLData  string = "/dataChannel"
	Task011URL      string = "/zTask011/"
	Task011URLData  string = "/dataChannelSema"

	Task_001_01_Code string = "ui/static/code/task_001.txt"
	Task_002_01_Code string = "ui/static/code/task_002.txt"

	MonthsInfo_json  string = "ui/static/json/months.json"
	PhoneInfo_json   string = "ui/static/json/phone.json"
	AutoInfoJson     string = "ui/static/json/auto.json"
	CriminalInfoJson string = "ui/static/json/criminal.json"
	SortCustomJson   string = "ui/static/json/sortCustomType.json"

	MonthsInfo_image string = "/ui/static/img/months/m_00.png"
)

var channelTask10 chan string
var channelTask11 chan string
var stopChannelTask10 chan bool
var flagXTask10 bool = true
var flagTask10 *bool = &flagXTask10

func Handler() {

	channelTask10 = make(chan string)
	channelTask11 = make(chan string)
	stopChannelTask10 = make(chan bool)
	*flagTask10 = true

	mux := http.NewServeMux()
	fileServer := http.FileServer(NeuteredFileSystem{http.Dir("ui")})
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")

	mux.Handle("/ui", http.NotFoundHandler())
	mux.Handle("/ui/", http.StripPrefix("/ui/", fileServer))

	mux.HandleFunc("/", Home)
	mux.Handle("/about/", httpHandler{message: "/about/", pageName: AboutPage})
	mux.Handle("/contacts/", httpHandler{message: "/contacts/", pageName: ContactsPage})

	mux.HandleFunc("/errorLog/", ServeHTTPerrorLog)

	mux.HandleFunc("/z_task_001_01/", ServeHTTP_task_001)
	mux.HandleFunc("/z_task_001_02/", ServeHTTP_task_001)
	mux.HandleFunc("/z_task_001_03/", ServeHTTP_task_001)

	mux.HandleFunc("/z_task_002_01/", ServeHTTP_task_002)
	mux.HandleFunc("/z_task_002_02/", ServeHTTP_task_002)
	mux.HandleFunc("/z_task_002_03/", ServeHTTP_task_002)
	mux.HandleFunc("/process_task_002_02_calc", processHandler_task_002_02_calc)
	mux.HandleFunc("/process_task_002_03_calc", processHandler_task_002_03_calc)

	mux.HandleFunc("/z_task_003_01/", ServeHTTP_task_003)
	mux.HandleFunc("/process_task_003_01_calc", processHandler_task_003_01_calc)
	mux.HandleFunc("/process_task_003_01_show", processHandler_task_003_01_show)

	mux.HandleFunc("/zTask004/", ServeHTTPTask004)
	mux.HandleFunc("/processHandlerTask004Show", processHandlerTask004Show)

	mux.HandleFunc("/zTask005/", ServeHTTPTask005)
	mux.HandleFunc("/processHandlerTask005Show", processHandlerTask005Show)

	mux.HandleFunc("/zTask006/", ServeHTTPTask006)
	mux.HandleFunc("/processHandlerTask006ShowSortArray", processHandlerTask006ShowSortArray)
	mux.HandleFunc("/processHandlerTask006ShowSortArrayTest", processHandlerTask006ShowSortArrayTest)

	mux.HandleFunc("/zTask007/", ServeHTTPTask007)
	mux.HandleFunc("/processHandlerTask007treeCreate", processHandlerTask007treeCreate)
	mux.HandleFunc("/processHandlerTask007treeAction", processHandlerTask007treeAction)
	mux.HandleFunc("/processHandlerTask007graphBfsAction", processHandlerTask007graphBfsAction)
	mux.HandleFunc("/processHandlerTask007graphDfsAction", processHandlerTask007graphDfsAction)

	mux.HandleFunc("/zTask008/", ServeHTTPTask008)
	mux.HandleFunc("/processHandlerTask008mapEqualAction", processHandlerTask008mapEqualAction)
	mux.HandleFunc("/data", processHandlerTask008mapRefreshData)

	mux.HandleFunc("/zTask009/", ServeHTTPTask009)
	mux.HandleFunc("/processHandlerTask009Action", processHandlerTask009Action)

	mux.HandleFunc("/zTask010/", ServeHTTPTask010)
	mux.HandleFunc("/dataChannel", processHandlerTask010RefreshData)
	mux.HandleFunc("/start", processHandlerTask010Start)
	mux.HandleFunc("/stop", processHandlerTask010Stop)

	mux.HandleFunc("/zTask011/", ServeHTTPTask011)
	mux.HandleFunc("/dataChannelSema", processHandlerTask011RefreshData)
	mux.HandleFunc("/startSema", processHandlerTask011Start)

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ERROR:Server not listening")
	}

}

func handlePanic(w http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		// Обработка паники - загрузка стартовой страницы Home
		//http.Redirect(w, r, "/", http.StatusSeeOther)
		http.Redirect(w, r, "/errorLog/", http.StatusSeeOther)
		errorhandler.PrintErrorPanic(fmt.Errorf("%v", err))
	}
}
