package libavutil

//#cgo pkg-config: libavutil
//#include <libavutil/channel_layout.h>
import "C"

type ChannelLayout uint64

const (
	AV_CH_FRONT_LEFT            ChannelLayout = 0x00000001
	AV_CH_FRONT_RIGHT           ChannelLayout = 0x00000002
	AV_CH_FRONT_CENTER          ChannelLayout = 0x00000004
	AV_CH_LOW_FREQUENCY         ChannelLayout = 0x00000008
	AV_CH_BACK_LEFT             ChannelLayout = 0x00000010
	AV_CH_BACK_RIGHT            ChannelLayout = 0x00000020
	AV_CH_FRONT_LEFT_OF_CENTER  ChannelLayout = 0x00000040
	AV_CH_FRONT_RIGHT_OF_CENTER ChannelLayout = 0x00000080
	AV_CH_BACK_CENTER           ChannelLayout = 0x00000100
	AV_CH_SIDE_LEFT             ChannelLayout = 0x00000200
	AV_CH_SIDE_RIGHT            ChannelLayout = 0x00000400
	AV_CH_TOP_CENTER            ChannelLayout = 0x00000800
	AV_CH_TOP_FRONT_LEFT        ChannelLayout = 0x00001000
	AV_CH_TOP_FRONT_CENTER      ChannelLayout = 0x00002000
	AV_CH_TOP_FRONT_RIGHT       ChannelLayout = 0x00004000
	AV_CH_TOP_BACK_LEFT         ChannelLayout = 0x00008000
	AV_CH_TOP_BACK_CENTER       ChannelLayout = 0x00010000
	AV_CH_TOP_BACK_RIGHT        ChannelLayout = 0x00020000
	AV_CH_STEREO_LEFT           ChannelLayout = 0x20000000 ///< Stereo downmix.
	AV_CH_STEREO_RIGHT          ChannelLayout = 0x40000000 ///< See AV_CH_STEREO_LEFT.
// AV_CH_WIDE_LEFT      ChannelLayout =        0x0000000080000000ULL
// AV_CH_WIDE_RIGHT      ChannelLayout =       0x0000000100000000ULL
// AV_CH_SURROUND_DIRECT_LEFT  ChannelLayout = 0x0000000200000000ULL
// AV_CH_SURROUND_DIRECT_RIGHT ChannelLayout = 0x0000000400000000ULL
// AV_CH_LOW_FREQUENCY_2     ChannelLayout =   0x0000000800000000ULL
// /** Channel mask value used for AVCodecContext.request_channel_layout
//     to indicate that the user requests the channel order of the decoder output
//     to be the native codec channel order. */
// AV_CH_LAYOUT_NATIVE      ChannelLayout =    0x8000000000000000ULL
)

const (
	AV_CH_LAYOUT_MONO              ChannelLayout = AV_CH_FRONT_CENTER
	AV_CH_LAYOUT_STEREO            ChannelLayout = AV_CH_FRONT_LEFT | AV_CH_FRONT_RIGHT
	AV_CH_LAYOUT_2POINT1           ChannelLayout = AV_CH_LAYOUT_STEREO | AV_CH_LOW_FREQUENCY
	AV_CH_LAYOUT_2_1               ChannelLayout = AV_CH_LAYOUT_STEREO | AV_CH_BACK_CENTER
	AV_CH_LAYOUT_SURROUND          ChannelLayout = AV_CH_LAYOUT_STEREO | AV_CH_FRONT_CENTER
	AV_CH_LAYOUT_3POINT1           ChannelLayout = AV_CH_LAYOUT_SURROUND | AV_CH_LOW_FREQUENCY
	AV_CH_LAYOUT_4POINT0           ChannelLayout = AV_CH_LAYOUT_SURROUND | AV_CH_BACK_CENTER
	AV_CH_LAYOUT_4POINT1           ChannelLayout = AV_CH_LAYOUT_4POINT0 | AV_CH_LOW_FREQUENCY
	AV_CH_LAYOUT_2_2               ChannelLayout = AV_CH_LAYOUT_STEREO | AV_CH_SIDE_LEFT | AV_CH_SIDE_RIGHT
	AV_CH_LAYOUT_QUAD              ChannelLayout = AV_CH_LAYOUT_STEREO | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT
	AV_CH_LAYOUT_5POINT0           ChannelLayout = AV_CH_LAYOUT_SURROUND | AV_CH_SIDE_LEFT | AV_CH_SIDE_RIGHT
	AV_CH_LAYOUT_5POINT1           ChannelLayout = AV_CH_LAYOUT_5POINT0 | AV_CH_LOW_FREQUENCY
	AV_CH_LAYOUT_5POINT0_BACK      ChannelLayout = AV_CH_LAYOUT_SURROUND | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT
	AV_CH_LAYOUT_5POINT1_BACK      ChannelLayout = AV_CH_LAYOUT_5POINT0_BACK | AV_CH_LOW_FREQUENCY
	AV_CH_LAYOUT_6POINT0           ChannelLayout = AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_CENTER
	AV_CH_LAYOUT_6POINT0_FRONT     ChannelLayout = AV_CH_LAYOUT_2_2 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER
	AV_CH_LAYOUT_HEXAGONAL         ChannelLayout = AV_CH_LAYOUT_5POINT0_BACK | AV_CH_BACK_CENTER
	AV_CH_LAYOUT_6POINT1           ChannelLayout = AV_CH_LAYOUT_5POINT1 | AV_CH_BACK_CENTER
	AV_CH_LAYOUT_6POINT1_BACK      ChannelLayout = AV_CH_LAYOUT_5POINT1_BACK | AV_CH_BACK_CENTER
	AV_CH_LAYOUT_6POINT1_FRONT     ChannelLayout = AV_CH_LAYOUT_6POINT0_FRONT | AV_CH_LOW_FREQUENCY
	AV_CH_LAYOUT_7POINT0           ChannelLayout = AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT
	AV_CH_LAYOUT_7POINT0_FRONT     ChannelLayout = AV_CH_LAYOUT_5POINT0 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER
	AV_CH_LAYOUT_7POINT1           ChannelLayout = AV_CH_LAYOUT_5POINT1 | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT
	AV_CH_LAYOUT_7POINT1_WIDE      ChannelLayout = AV_CH_LAYOUT_5POINT1 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER
	AV_CH_LAYOUT_7POINT1_WIDE_BACK ChannelLayout = AV_CH_LAYOUT_5POINT1_BACK | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER
	AV_CH_LAYOUT_OCTAGONAL         ChannelLayout = AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_LEFT | AV_CH_BACK_CENTER | AV_CH_BACK_RIGHT
	AV_CH_LAYOUT_STEREO_DOWNMIX    ChannelLayout = AV_CH_STEREO_LEFT | AV_CH_STEREO_RIGHT
	// AV_CH_LAYOUT_HEXADECAGONAL     ChannelLayout = AV_CH_LAYOUT_OCTAGONAL | AV_CH_WIDE_LEFT | AV_CH_WIDE_RIGHT | AV_CH_TOP_BACK_LEFT | AV_CH_TOP_BACK_RIGHT | AV_CH_TOP_BACK_CENTER | AV_CH_TOP_FRONT_CENTER | AV_CH_TOP_FRONT_LEFT | AV_CH_TOP_FRONT_RIGHT
)

/**
 * Return the number of channels in the channel layout.
 */
func (cl ChannelLayout) GetNbChannels() int {
	return int(C.av_get_channel_layout_nb_channels((C.uint64_t)(cl)))
}

/**
 * Get the name of a given channel.
 *
 * @return channel name on success, NULL on error.
 */
func (cl ChannelLayout) Name() string {
	return C.GoString(C.av_get_channel_name((C.uint64_t)(cl)))
}

/**
 * Get the description of a given channel.
 *
 * @param channel  a channel layout with a single channel
 * @return  channel description on success, NULL on error
 */
func (cl ChannelLayout) Description() string {
	return C.GoString(C.av_get_channel_description((C.uint64_t)(cl)))
}

/**
 * Get the index of a channel in channel_layout.
 *
 * @param channel a channel layout describing exactly one channel which must be
 *                present in channel_layout.
 *
 * @return index of channel in channel_layout on success, a negative AVERROR
 *         on error.
 */
func (cl ChannelLayout) GetChannelIndex(channel uint64) int {
	return int(C.av_get_channel_layout_channel_index((C.uint64_t)(cl), (C.uint64_t)(channel)))
}

/**
 * Return the number of channels in the channel layout.
 */
func GetDefaultChannelLayout(nbChannels int) int64 {
	return int64(C.av_get_default_channel_layout((C.int)(nbChannels)))
}
