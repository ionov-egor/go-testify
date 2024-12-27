package main

import (
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, errWrite := w.Write([]byte("count missing"))
		if errWrite != nil {
			http.Error(w, errWrite.Error(), http.StatusBadRequest)
			return
		}

		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, errWrite := w.Write([]byte("wrong count value"))
		if errWrite != nil {
			http.Error(w, errWrite.Error(), http.StatusBadRequest)
			return
		}

		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, errWrite := w.Write([]byte("wrong city value"))
		if errWrite != nil {
			http.Error(w, errWrite.Error(), http.StatusBadRequest)
			return
		}
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write([]byte(answer))
	if errWrite != nil {
		http.Error(w, errWrite.Error(), http.StatusBadRequest)
		return
	}

}

func main() {
	http.HandleFunc(`/cafe`, mainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
