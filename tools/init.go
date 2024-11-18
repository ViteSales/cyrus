package tools

import "github.com/vitesales/cyrus/tools/logging"

var log logging.Logger

func init() {
	log = logging.GetLogger("tools")
}
