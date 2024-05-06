package models

type Req struct {
	UserInput string `json:"user_input"`
}

type Res struct {
	UserInput         string `json:"user_input"`
	BotResponse       string `json:"bot_response"`
	ResponseTimestamp string `json:"response_timestamp"`
}
