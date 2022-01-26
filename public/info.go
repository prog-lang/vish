package public

import (
	"fmt"
)

const (
	AuthorName     = "Viktor A. Rozenko Voitenko"
	AuthorEmail    = "sharp.vik@gmail.com"
	ProgramName    = "vish"
	ProgramVersion = "0.1.0"
)

func ShortInfo() string {
	return fmt.Sprintf("%s v%s by %s <%s>",
		ProgramName, ProgramVersion, AuthorName, AuthorEmail)
}
