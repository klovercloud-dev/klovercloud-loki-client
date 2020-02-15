package v1

import (
	"klovercloud-loki-client/pkg/apis/v1/loki/query"
	_range "klovercloud-loki-client/pkg/apis/v1/loki/query/range"
)

func queryRangeUrl() query.QueryResponse {
	return _range.NewUrlBuilder().Init().Get().Label("app","csi-cephfsplugin").Build().Fire()
}