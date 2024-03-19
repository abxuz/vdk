package format

import (
	"github.com/abxuz/vdk/av/avutil"
	"github.com/abxuz/vdk/format/aac"
	"github.com/abxuz/vdk/format/flv"
	"github.com/abxuz/vdk/format/mp4"
	"github.com/abxuz/vdk/format/rtmp"
	"github.com/abxuz/vdk/format/rtsp"
	"github.com/abxuz/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
