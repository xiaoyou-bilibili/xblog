package module

import (
	"fmt"
	"testing"
)

func TestNetPost(t *testing.T) {
	NetPost("https://music.163.com/weapi/v6/playlist/detail?csrf_token=e67273f96711aa6835d5f902d40e90f7",
		map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.73",
		},
		map[string]string{
			"params":    "hz7p7IkxvAhEyC9j+z6aNA3vMlvSg2WwzGPiGmQT8UhlSV7s/Nw2uR7BAeC3hLxY2QwP2SWrb55w8ahtp0ds7zvUb0WCv+BFpcJ2bEwYmg56HxVMPkNUQrwv+HoOZipalTROJofFzITkwFpAPqVfCLJWh6sGS0zIuOvvK6qEfBdR91/hO1aoFBhFSiamCI9/yxP+la4xymtOZOl401WJLXYx4awwy5j+Q45x1lBt3EU=",
			"encSecKey": "33894d186e9634107c54158c3c733cb7d7983848408ab950619773d9be5e2b0514b401d21ed68c46a90e08bd856f552131ae693deebf26deb53cd4831374826f243ec95c691f543157a01ca899ae3d83e8477ef222f797b11fe4f9e793a6fcc5c7486c6730eddb8bb718ddd90d9e269b44f25268ee4a207a944bbeb47d1a878f",
		},
		func(i interface{}, i2 interface{}) {
			fmt.Println(i, i2)
		},
	)
}
