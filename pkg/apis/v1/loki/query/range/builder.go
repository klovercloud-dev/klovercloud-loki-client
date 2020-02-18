package _range

import (
	"encoding/json"
	"io/ioutil"
	"github.com/klovercloud-dev/klovercloud-loki-client/config"
	"github.com/klovercloud-dev/klovercloud-loki-client/pkg/apis/common"
	"github.com/klovercloud-dev/klovercloud-loki-client/pkg/apis/v1/loki/query"
	"log"
	"net/http"
	"strconv"
)


const (
	URL = "loki/api/v1/query_range?"
)




type Builder interface {
	Get() Builder
	Post() Builder
	Query() Builder
	Init() Builder
	Label(labelName string, value string) Builder
	Limit(int) Builder
	Sum() Builder
	Rate(int) Builder
	Start(int64) Builder
	Step(int) Builder
	Contains(string) Builder
	NotContains(string) Builder
	Matches(expression string) Builder
	NotMatches(expression string) Builder
	End(int64) Builder
	CountOverTime(minutes int) Builder
	TopK(k int64) Builder
	Avg() Builder
	Build() Builder
	Fire() query.QueryResponse
}

type builder struct {
	method string
	url    string
	labels map[string]string
	start  int64
	end    int64
	limit  int
	step   int
	pipe   string
	body   interface{}
	query  string
}
func (qb *builder) Fire() query.QueryResponse {
	log.Println(qb.url)
	client := &http.Client{}
	req, err := http.NewRequest(qb.method, qb.url, nil)
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

func (qb *builder) TopK(k int64) Builder {
	temp:=""
	if(qb.query==""){
		temp=createKeyValuePairs(qb.labels)+qb.pipe
	}else{
		temp=qb.query
	}
	qb.query="topk("+strconv.Itoa(int(k))+","+temp+")"
	return qb
}


func (qb *builder) CountOverTime(minutes int) Builder {
	temp:=""
	if(qb.query==""){
		temp=createKeyValuePairs(qb.labels)+qb.pipe
	}else{
		temp=qb.query
	}
	qb.query="count_over_time("+temp+"["+strconv.Itoa(int(minutes))+"m])"
	return qb
}

func (qb *builder) Contains(str string) Builder {
	qb.pipe=qb.pipe+"|="+"\""+str+"\""
	return qb
}

func (qb *builder) NotContains(str string) Builder {
	qb.pipe=qb.pipe+"!="+"\""+str+"\""
	return qb
}

func (qb *builder) Matches(expression string) Builder {
	qb.pipe=qb.pipe+"|~"+"\""+expression+"\""
	return qb
}

func (qb *builder) NotMatches(expression string) Builder {
	qb.pipe=qb.pipe+"!~"+"\""+expression+"\""
	return qb
}


func (qb *builder) Start(start int64) Builder {
	qb.start=start
	return qb
}

func (qb *builder) Step(step int) Builder {
	qb.step=step
	return qb
}

func (qb *builder) End(end int64) Builder {
	qb.end=end
	return qb
}


func (qb *builder) Sum() Builder {
	temp:=""
	if(qb.query==""){
		temp=createKeyValuePairs(qb.labels)+qb.pipe
	}else{
		temp=qb.query
	}
	qb.query="sum("+temp+")"
	return qb
}

func (qb *builder) Avg() Builder {
	temp:=""
	if(qb.query==""){
		temp=createKeyValuePairs(qb.labels)+qb.pipe
	}else{
		temp=qb.query
	}
	qb.query="avg("+temp+")"
	return qb
}


func (qb *builder) Rate(minutes int) Builder {
	temp:=""
	if(qb.query==""){
		temp=createKeyValuePairs(qb.labels)+qb.pipe
	}else{
		temp=qb.query
	}
	qb.query="rate("+temp+"["+strconv.Itoa(int(minutes))+"m])"
	return qb
}

func (qb *builder) Init() Builder {
	qb.labels = make(map[string]string)
	return qb
}


func (qb *builder) Limit(limit int) Builder {
	qb.limit = limit
	return qb
}

func (qb *builder) Label(levelName string, value string) Builder {
	qb.labels[levelName] = value
	return qb
}

func (qb *builder) Query() Builder {
	qb.url = qb.url + "query="
	return qb
}

func (qb *builder) Get() Builder {
	qb.method = common.GET
	qb.url = config.LokiUrl + URL
	return qb
}

func (qb *builder) Post() Builder {
	qb.method = common.POST
	return qb
}

func (qb *builder) Build() Builder{
	str:=""
	if(qb.query==""){
		str= qb.url + "query=" + createKeyValuePairs(qb.labels)+qb.pipe
	}else{
		str=qb.url + "query=" +qb.query
	}
	if qb.limit != 0 {
		str = str + "&limit=" + strconv.Itoa(int(qb.limit))
	}

	if qb.start != 0 {
		str = str + "&start=" + strconv.FormatInt(qb.start, 10)
	}

	if qb.end != 0 {
		str = str + "&end=" + strconv.FormatInt(qb.end, 10)
	}

	if qb.step != 0 {
		str = str + "&step=" + strconv.Itoa(int(qb.step))
	}
	qb.url = str
	return qb
}



func NewBuilder() Builder {
	return &builder{}
}

func createKeyValuePairs(m map[string]string) string {
	count := 0
	str := "{"
	for key, value := range m {
		count++
		if count < len(m) {
			str = str + key + "=" + "\"" + value + "\"" + ","
		} else {
			str = str + key + "=" + "\"" + value + "\"" + "}"
		}
	}
	return str
}

