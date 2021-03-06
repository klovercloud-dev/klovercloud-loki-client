package v1

import "github.com/klovercloud-dev/klovercloud-loki-client/pkg/apis/v1/loki/labels"

func levelUrl() interface{}{
	return labels.NewUrlBuilder().Init().Get().Build().Fire()
}


func valuesUrl() interface{}{
	return labels.NewUrlBuilder().Init().Get().Values("app").Build().Fire()
}
