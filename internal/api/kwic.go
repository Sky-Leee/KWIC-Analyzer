package api

type KWICRequest struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type KWICResult struct {
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
