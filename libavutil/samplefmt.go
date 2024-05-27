package libavutil

//#cgo pkg-config: libavutil
//#include <libavutil/samplefmt.h>
import "C"
import "unsafe"

type AVSampleFormat int32

const (
	AV_SAMPLE_FMT_NONE AVSampleFormat = iota - 1
	AV_SAMPLE_FMT_U8                  ///< unsigned 8 bits
	AV_SAMPLE_FMT_S16                 ///< signed 16 bits
	AV_SAMPLE_FMT_S32                 ///< signed 32 bits
	AV_SAMPLE_FMT_FLT                 ///< float
	AV_SAMPLE_FMT_DBL                 ///< double
	AV_SAMPLE_FMT_U8P                 ///< unsigned 8 bits, planar
	AV_SAMPLE_FMT_S16P                ///< signed 16 bits, planar
	AV_SAMPLE_FMT_S32P                ///< signed 32 bits, planar
	AV_SAMPLE_FMT_FLTP                ///< float, planar
	AV_SAMPLE_FMT_DBLP                ///< double, planar
	AV_SAMPLE_FMT_S64                 ///< signed 64 bits
	AV_SAMPLE_FMT_S64P                ///< signed 64 bits, planar
	AV_SAMPLE_FMT_NB                  ///< Number of sample formats. DO NOT USE if linking dynamically

)

/**
 * Return the name of sample_fmt, or NULL if sample_fmt is not
 * recognized.
 */
func (sf AVSampleFormat) Name() string {
	return C.GoString(C.av_get_sample_fmt_name((C.enum_AVSampleFormat)(sf)))
}

/**
 * Check if the sample format is planar.
 *
 * @param sample_fmt the sample format to inspect
 * @return 1 if the sample format is planar, 0 if it is interleaved
 */
func (sf AVSampleFormat) IsPlanar() bool {
	return C.av_sample_fmt_is_planar((C.enum_AVSampleFormat)(sf)) == 1
}

/**
 * Return number of bytes per sample.
 *
 * @param sample_fmt the sample format
 * @return number of bytes per sample or zero if unknown for the given
 * sample format
 */
func (sf AVSampleFormat) GetBytesPerSample() int {
	return int(C.av_get_bytes_per_sample((C.enum_AVSampleFormat)(sf)))
}

/**
 * Get the required buffer size for the given audio parameters.
 *
 * @param[out] linesize calculated linesize, may be NULL
 * @param nb_channels   the number of channels
 * @param nb_samples    the number of samples in a single channel
 * @param sample_fmt    the sample format
 * @param align         buffer size alignment (0 = default, 1 = no alignment)
 * @return              required buffer size, or negative error code on failure
 */
func SamplesGetBufferSize(
	nbChannels, nbSamples int, sampleFmt AVSampleFormat, align int,
) (linesize int, ret int) {
	ret = (int)(C.av_samples_get_buffer_size(
		(*C.int)(unsafe.Pointer(&linesize)),
		(C.int)(nbChannels),
		(C.int)(nbSamples),
		(int32)(sampleFmt),
		(C.int)(align),
	))
	return
}
