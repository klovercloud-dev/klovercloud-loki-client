package v1

import "github.com/klovercloud-dev/klovercloud-loki-client/pkg/apis/v1/loki/tail"

func WSUrl()  interface{} {
	return tail.NewUrlBuilder().Init().Get().Label("app","klovercloud-agent").Build().Fire()
}
