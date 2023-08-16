package faq

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"vector_demo/internal/logic/embedding"
	"vector_demo/internal/logic/vector"
	"vector_demo/sdk/zilliz"
)

type Faq struct {
	LiveId   string
	Question string
	Answer   string
}

func NewFaq(liveId, question string) *Faq {
	return &Faq{
		LiveId:   liveId,
		Question: question,
	}
}

func (f *Faq) Ask() (answer string) {
	ctx := context.Background()
	url, _ := g.Cfg().Get(ctx, "zilliz.url")
	key, _ := g.Cfg().Get(ctx, "zilliz.key")

	vectorDB := vector.NewVectorDB(url.String(), key.String())

	// 1. 生成向量
	vector, err := embedding.GenerateVector(f.Question)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return
	}

	// 2. 查询
	filter := make(map[string]string)
	filter["liveId"] = f.LiveId

	resp, err := vectorDB.SearchVector(vector, filter)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
		return
	}

	f.filterAnswer(resp)

	fmt.Printf("answer: %s", f.Answer)

	return
}

func (f *Faq) filterAnswer(resp *zilliz.QueryResp) (answer string) {
	if resp.Code != 0 {
		fmt.Printf("resp code: %d", resp.Code)
		return
	}

	if len(resp.Data) == 0 {
		fmt.Printf("resp data is null")
		return
	}

	answer = resp.Data[0].Text

	return
}
