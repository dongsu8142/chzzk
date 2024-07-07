# CHZZK

Go로 작성된 네이버의 스트리밍 서비스 CHZZK 라이브러리입니다.

기능들

- [x] 로그인 (쿠키)
- [x] 검색 (채널, 영상, 생방송)
- [ ] 방송 상세 정보 조회
- [ ] 채팅

## 설치
```bash
go get github.com/dongsu8142/chzzk
```

## 예시
`chzzk.naver.com`에 로그인하신 후, 개발자 도구를 열어 `Application > Cookies > https://chzzk.naver.com` 에서 `NID_AUT`, `NID_SES` 쿠키를 복사해서 로그인할 수 있습니다.

```go

package chzzk_test

import (
	"fmt"
	"testing"

	"github.com/dongsu8142/chzzk"
)

func TestMain(m *testing.M) {
	client := chzzk.New("NID_AUT", "NID_SES")
	channels, err := client.SearchChannels("괴물쥐")
	if err != nil {
		panic(err)
	}
	fmt.Println(channels[0].ChannelName)
	lives, err := client.SearchLives("괴물쥐")
	if err != nil {
		panic(err)
	}
	fmt.Println(lives[0].Live.ChannelID)
	videos, err := client.SearchVideos("괴물쥐")
	if err != nil {
		panic(err)
	}
	fmt.Println(videos[0].Video.VideoTitle)
}
