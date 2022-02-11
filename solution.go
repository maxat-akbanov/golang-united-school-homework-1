package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	greetingMessage := emoji.Sprint("Hello :world_map:!")
	return greetingMessage
}
