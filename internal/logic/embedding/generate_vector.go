package embedding

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"vector_demo/sdk/azure"
	"vector_demo/sdk/zilliz"
)

type Embedding struct {
}

func GenerateVector(text string) (vector []float64, err error) {
	ctx := context.Background()
	key, _ := g.Cfg().Get(ctx, "azure.key")
	deploymentName, _ := g.Cfg().Get(ctx, "azure.deploymentName")
	resourceName, _ := g.Cfg().Get(ctx, "azure.resourceName")

	configuration := azure.NewDefaultConfiguration(resourceName.String(), deploymentName.String(), key.String())
	input := azure.DocumentInput{}
	input.Input = text
	embedding, err := configuration.GetEmbedding(input)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	if embedding == nil {
		err = gerror.New("embedding resp is nil")
		return
	}

	vector = embedding.Data[0].Embedding

	return
}

func InsertVector() {
	var text = "我想了解一下北京"
	var liveId = "liveId123123"
	vector, _ := GenerateVector(text)
	if len(vector) == 0 {
		fmt.Printf("vector is null")
		return
	}

	ctx := context.Background()
	url, _ := g.Cfg().Get(ctx, "zilliz.url")
	key, _ := g.Cfg().Get(ctx, "zilliz.key")

	data := zilliz.VectorData{
		CollectionName: "faq",
		Data: zilliz.Data{
			Vector: vector,
			Text:   text,
			LiveId: liveId,
		},
	}

	zillizSDK := zilliz.New(url.String(), key.String())
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
