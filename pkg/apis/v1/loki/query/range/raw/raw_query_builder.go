package raw

import (
	"encoding/json"
	"github.com/klovercloud-dev/klovercloud-loki-client/config"
	"github.com/klovercloud-dev/klovercloud-loki-client/pkg/apis/common"
	"github.com/klovercloud-dev/klovercloud-loki-client/pkg/apis/v1/loki/query"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	URL = "loki/api/v1/query_range?"
)

type Builder interface {
	Init() Builder
	Get() Builder
	Query( Query string) Builder
	Build() Builder
	Fire() query.QueryResponse

}


type builder struct {
	method string
	url    string
	body   interface{}
	query  string

}
func NewBuilder() Builder {
	return &builder{}
}

func (qb *builder) Init() Builder {
	return qb
}

func (qb *builder) Fire() query.QueryResponse {
	log.Println(URL)
	client := &http.Client{}
	req, err := http.NewRequest(qb.method,qb.url, nil)
	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Add("Content-Type","application/json")
	log.Println("Requesting:",qb.url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	response :=query.QueryResponse{}
	json.Unmarshal([]byte(body), &response)
	return response

}

func (qb *builder) Get() Builder {
	qb.method = common.GET
	qb.url = config.LokiUrl + URL
	return qb
}


func (qb *builder) Query(query string) Builder {
	qb.query = query
	return qb
}

func (qb *builder) Build() Builder{
	qb.url=qb.url + "query=" +qb.query
	return qb
}
