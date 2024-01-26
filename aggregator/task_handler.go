package aggregator

import "net/http"

type taskHandler struct {
	aggregator *Aggregator
}

func (h *taskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s1 := r.URL.Query().Get("system1Value")
	s2 := r.URL.Query().Get("system2Value")
	dkg := r.URL.Query().Get("dkgValue")

	task := MfssiaTask{
		system1Value: s1,
		system2Value: s2,
		dkgValue:     dkg,
	}

	h.aggregator.sendNewTask(task)
}
