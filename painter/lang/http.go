package lang

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/AidXylelele/go_lab_3/painter"
)

func HttpHandler(loop *painter.Loop, p *Parser) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		var in io.Reader = r.Body
		if r.Method == http.MethodGet {
			in = strings.NewReader(r.URL.Query().Get("cmd"))
		}

		cmds, err := p.Parse(in)
		if err != nil {
			log.Printf("Bad script: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		loop.Post(painter.PainterOperationList(cmds))
		log.Printf("Operations: %s", cmds)
		rw.WriteHeader(http.StatusOK)
	})
}
