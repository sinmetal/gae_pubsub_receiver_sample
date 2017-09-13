package gae_pubsub_receiver_sample

import (
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/_ah/push-handlers/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	for k, v := range r.Header {
		log.Infof(ctx, "%s = %v\n", k, v)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf(ctx, "Failed read request body : %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Infof(ctx, "\n%s", string(body))
	w.WriteHeader(http.StatusOK)
}
