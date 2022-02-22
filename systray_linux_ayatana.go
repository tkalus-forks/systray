//go:build linux && !legacy_appindicator

package systray

/*
#cgo linux pkg-config: ayatana-appindicator3-0.1
#cgo linux CFLAGS: -DUSE_AYATANA

#include "systray.h"
*/
import "C"
