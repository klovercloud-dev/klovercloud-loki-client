package v1

import "klovercloud-loki-client/pkg/apis/v1/loki/query"

func queryByLabel() query.QueryResponse {
	return query.NewBuilder().Init().Get().Label("app","csi-cephfsplugin").Build().Fire()
}

func queryBySurmAndRate() query.QueryResponse {
	return query.NewBuilder().Init().Get().Label("app","csi-cephfsplugin").Rate(10).Sum().Build().Fire()
}

