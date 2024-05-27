package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import "unsafe"

func (p *AvSubtitle) NumRects() int {
	return (int)(p.num_rects)
}

func (p *AvSubtitle) Format() int {
	return (int)(p.format)
}

func (p *AvSubtitle) Pts() int64 {
	return (int64)(p.pts)
}

func (p *AvSubtitle) StartDisplayTime() uint32 {
	return (uint32)(p.start_display_time)
}

func (p *AvSubtitle) EndDisplayTime() uint32 {
	return (uint32)(p.end_display_time)
}

// https://stackoverflow.com/questions/49987098/how-to-access-a-c-pointer-array-from-golang
func (p *AvSubtitle) GetAvSubtitleRect(idx int) *AvSubtitleRect {
	if idx >= 0 && idx < int(p.num_rects) {
		streams := (**AvSubtitleRect)(unsafe.Pointer(p.rects))
		i := (uintptr)(idx)
		ptrPtr := (**AvSubtitleRect)(unsafe.Pointer(uintptr(unsafe.Pointer(streams)) + i*unsafe.Sizeof(*streams)))
		return *ptrPtr
	}

	return nil
}

const (
	SubTitleNone   = C.SUBTITLE_NONE
	SubTitleBitmap = C.SUBTITLE_BITMAP
	SubTitleText   = C.SUBTITLE_TEXT
	SubTitleAss    = C.SUBTITLE_ASS
)

func (p *AvSubtitleRect) Type() int {
	return (int)(p._type)
}

func (p *AvSubtitleRect) Text() string {
	return (C.GoString)(p.text)
}

func (p *AvSubtitleRect) Ass() string {
	return (C.GoString)(p.ass)
}

func (p *AvSubtitleRect) X() int {
	return (int)(p.x)
}

func (p *AvSubtitleRect) Y() int {
	return (int)(p.y)
}

func (p *AvSubtitleRect) W() int {
	return (int)(p.w)
}

func (p *AvSubtitleRect) H() int {
	return (int)(p.h)
}

func (p *AvSubtitleRect) Flags() int {
	return (int)(p.flags)
}

func (p *AvSubtitleRect) NbColors() int {
	return (int)(p.nb_colors)
}
