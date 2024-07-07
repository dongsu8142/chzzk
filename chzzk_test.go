package chzzk_test

import (
	"fmt"
	"testing"

	"github.com/dongsu8142/chzzk"
)

func TestMain(m *testing.M) {
	client := chzzk.New("ammkAu7jt9ZmLXR8YIwTb0XqD59SzqfdxROm1IPtRKAWghPj0Pwct9jZdTAUQbMi", "AAABkvxVBeDlGL+192+nrmBrKkklSBF3NnG3l4gOGt3RKsE3zM0WXPkygoRb8ao2ir8Z0clztTjechq2PbPC1XzqDgkN8hrXsGiK9xWipymJEwQBCW+CvzhOSxaML4zsTWXuze+s8NL+Hcevtut/l1SLgJzs6EahkhJyjceWxIJ9/Ahqq6Jk+tPpYat6vrOALp6GWzzvHIZiUZ7oWmfj2+8HVVUarEjNANbfiijmdgiScvBaeKfNQUQiHdIG6NZjeVONv4tC2+3TSETsJwc2rxpp1XkOK9YWbsIm7lfhAKQrTQyUSUBwdmtotufGm2TbNWbDp82XMXsXZotUzHSOCrwJurJL62iYV6DKBHMArK6OuMGuKCAxNkd18Y/oHF4JMze3M0DegOArLvNWSaYpoorVtflT6bTaBD+iyxga2yBOmVGdloMk6S4LhZtXz4ewepXSlY21Fmt9NIyyywN7dkZJ7yVzVZ6GPZj21d3CHcws3DRzUEH9aWuSawnEp1xYzRXSXRA650718e9R6HJMsDJBHusZXkXd3KXPkCNARFJ1d8Gy")
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
