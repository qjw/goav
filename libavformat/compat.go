//go:build !windows

package libavformat

import "C"

// x86_64-w64-mingw32 检查编译类型不对
type MallocType = C.ulong
