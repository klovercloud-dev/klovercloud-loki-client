package labels

import (
	"encoding/json"
	"io/ioutil"
	"klovercloud-loki-client/config"
	"klovercloud-loki-client/pkg/apis/common"
	"log"
	"net/http"
)

const (
	URL = "loki/api/v1/"
)


type Builder interface {
	Get() Builder
	Init() Builder
	Values(string) Builder
	Build() Builder
	Fire() interface{}
}

type builder struct {
	method  string
	url string
	values string
}


func (qb *builder) Fire() interface{} {
	log.Println(string(qb.url))
	client := &http.Client{}
	req, err := http.NewRequest(qb.method, qb.url, nil)
	req.SetBasicAuth(config.Username, config.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)

	return data

}

func (qb *builder) Values(level string) Builder {
	qb.values=level
	return qb
}

func (qb *builder) Init() Builder {
	qb.url= config.LokiUrl
	return qb
}


func (qb *builder) Get() Builder {
	qb.method = common.GET
	qb.url = qb.url + URL
	return qb
}



func (qb *builder) Build() Builder {
	if(qb.values==""){
		qb.url=qb.url+"labels"
	}else{
		qb.url=qb.url+"label/"+qb.values+"/values"
	}
	return qb
}

func NewUrlBuilder() Builder {
	return &builder{}
}


