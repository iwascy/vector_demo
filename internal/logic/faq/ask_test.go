package faq

import (
	"fmt"
	"testing"
)

func TestFaq_Ask(t *testing.T) {
	var liveId = "liveId123123"
	var question = "我想了解一下北京"
	faq := NewFaq(liveId, question)
	faq.Ask()

	fmt.Printf("answer: %s", faq.Answer)
}
