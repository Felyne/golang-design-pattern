package main

type SensitiveWordFilter interface {
	DoFilter(content string) bool
}

type SensitiveWordFilterChain struct {
	filterList []SensitiveWordFilter
}

func NewSensitiveWordFilterChain() *SensitiveWordFilterChain {
	return &SensitiveWordFilterChain{}
}

func (h *SensitiveWordFilterChain) AddFilter(filters ...SensitiveWordFilter) {
	h.filterList = append(h.filterList, filters...)
}

func (h *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range h.filterList {
		if false == filter.DoFilter(content) {
			return false
		}
	}
	return true
}

type SexyWordFilter struct {
}

func NewSexyWordFilter() SensitiveWordFilter {
	return &SexyWordFilter{}
}

func (*SexyWordFilter) DoFilter(content string) bool {
	if content == "sexy" {
		return false
	}
	return true
}

type AbusiveWordFilter struct {
}

func NewAbusiveWordFilter() SensitiveWordFilter {
	return &AbusiveWordFilter{}
}

func (*AbusiveWordFilter) DoFilter(content string) bool {
	if content == "fuck" {
		return false
	}
	return true
}
