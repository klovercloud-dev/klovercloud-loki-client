package v1

import (
	"github.com/klovercloud-dev/klovercloud-loki-client/pkg/apis/v1/loki/query"
	_range "github.com/klovercloud-dev/klovercloud-loki-client/pkg/apis/v1/loki/query/range"
)

func queryRangeUrl() query.QueryResponse {
	return _range.NewBuilder().Init().Get().Label("app","csi-cephfsplugin").Build().Fire()
}