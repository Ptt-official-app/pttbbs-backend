package mand

import (
	context "context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type MockClientConn struct{}

func NewMockClientConn() *MockClientConn {
	return &MockClientConn{}
}

func (c *MockClientConn) Invoke(ctx context.Context, method string, args any, reply any, opts ...grpc.CallOption) (err error) {
	logrus.Infof("mand.MockClient: method: %v", method)
	switch method {
	case "/pttbbs.man.ManService/List":
		r := reply.(*ListReply)
		r.Entries = []*Entry{
			testArticle0, testArticle1, testArticle2, testArticle3, testArticle4, testArticle5, testArticle6, testArticle7, testArticle8,
		}
	case "/pttbbs.man.ManService/Article":
		params := args.(*ArticleRequest)
		r := reply.(*ArticleReply)

		switch params.Path {
		case "M.1608386280.A.BC9":
			r.Content = []byte("作者: SYSOP () 看板: WhoAmI\n標題: [心得] 測試一下特殊字～\n時間: Sat Dec 19 21:57:58 2020\n\n※這樣子有綠色嗎？～\n※ 發信站\n\n--\n※ 發信站: 批踢踢 docker(pttdocker.test), 來自: 172.22.0.1")
		}
	}

	return nil
}

func (c *MockClientConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (cs grpc.ClientStream, err error) {
	return nil, nil
}

func (c *MockClientConn) Close() {
}
