package repository

type ChatRepository interface {
	AnalyzeData()
	ChatWithAI()
}

type chatRepoImpl struct {
	HuggingfaceToken string
}

func NewChatRepo(HuggingfaceToken string) *chatRepoImpl {
	return &chatRepoImpl{HuggingfaceToken: HuggingfaceToken}
}

func (c *chatRepoImpl) AnalyzeData() {

}

func ChatWithAI() {

}
