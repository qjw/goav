// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/log.h>
	#include <stdlib.h>
*/
import "C"

type AvLogLevel int

// AV_LOG_xxx
const (
	AvLogQuiet   AvLogLevel = -8
	AvLogPanic   AvLogLevel = 0
	AvLogFatal   AvLogLevel = 8
	AvLogError   AvLogLevel = 16
	AvLogWarning AvLogLevel = 24
	AvLogInfo    AvLogLevel = 32
	AvLogVerbose AvLogLevel = 40
	AvLogDebug   AvLogLevel = 48
	AvLogTrace   AvLogLevel = 56
)

// AvLogSetLevel Set the log level
func AvLogSetLevel(level AvLogLevel) {
	C.av_log_set_level(C.int(level))
}

// AvLogGetLevel Get the current log level
func AvLogGetLevel() AvLogLevel {
	return AvLogLevel(C.av_log_get_level())
}
