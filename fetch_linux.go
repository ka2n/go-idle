package idle

import (
	"time"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/screensaver"
	"github.com/jezek/xgb/xproto"
)

func parseIdleFromXCB() (time.Duration, error) {
	conn, err := xgb.NewConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	info := xproto.Setup(conn)
	screen := info.DefaultScreen(conn)

	if err := screensaver.Init(conn); err != nil {
		return 0, err
	}

	rep, err := screensaver.QueryInfo(conn, xproto.Drawable(screen.Root)).Reply()
	if err != nil {
		return 0, err
	}

	return time.Duration(rep.MsSinceUserInput) * time.Millisecond, nil
}

// Get idle time for Linux Xorg
func Get() (time.Duration, error) {
	return parseIdleFromXCB()
}
