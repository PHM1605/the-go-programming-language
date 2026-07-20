package storage2

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// NEW: save the original "notifyUser()" to restore; after assigning it to dummy "notifyUser()" in this test
	saved := notifyUser
	defer func() {
		notifyUser = saved
	}()

	var notifiedUser, notifiedMsg string
	// NEW: attaching notifyUser() to a dummy one instead of the original
	// nothing was sent; we simply record the supposed-to-be-sent "user" and "message"
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	// ...simulate a 980MB-used condition
	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	// test for "user" and "message"
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, wantSubstring)
	}
}
