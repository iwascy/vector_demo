package vector

import (
	"fmt"
	"vector_demo/internal/logic/embedding"
	"vector_demo/sdk/zilliz"
)

type VectorDB struct {
	Url string
	Key string
}

func NewVectorDB(url, key string) VectorDB {
	return VectorDB{
		Url: url,
		Key: key,
	}
}

func (v *VectorDB) InsertVector() {
	var text = "我想了解一下北京"
	var liveId = "liveId123123"
	vector, _ := embedding.GenerateVector(text)
	if len(vector) == 0 {
		fmt.Printf("vector is null")
		return
	}

	data := zilliz.VectorData{
		CollectionName: "faq",
		Data: zilliz.Data{
			Vector: vector,
			Text:   text,
			LiveId: liveId,
		},
	}

	zillizSDK := zilliz.New(v.Url, v.Key)
	resp, err := zillizSDK.InsertCollection(data)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return
	}
	if resp != nil {
		fmt.Printf("resp: %+v", resp)
		return
	}

	return
}

func (v *VectorDB) SearchVector(vector []float64, filter map[string]string) (queryResp *zilliz.QueryResp, err error) {
	zillizSDK := zilliz.New(v.Url, v.Key)

	req := zilliz.QueryReq{}
	req.CollectionName = "faq"

	var filterStr string
	// (publication == 'Towards Data Science') and (claps == 1500 )
	var first = true
	for k, v := range filter {
		if !first {
			filterStr += " and "
		}
		queryStr := fmt.Sprintf("%s == '%s'", k, v)
		if !first {
			filterStr += fmt.Sprintf("(%s)", queryStr)
		} else {
			filterStr += queryStr
		}
		first = false
	}
	req.Filter = filterStr
	req.Vector = vector

	fmt.Printf("req: %+v \n", req)

	queryResp, err = zillizSDK.SearchCollection(req)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return
	}

	fmt.Printf("queryResp: %+v", queryResp)

	return
}
