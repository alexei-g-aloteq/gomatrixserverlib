package gomatrixserverlib

import (
	"bytes"
	"testing"
)

func TestRedactionAlgorithmV4(t *testing.T) {
	// Specifically, the version 4 redaction algorithm used in room
	// version 9 is ensuring that the `join_authorised_via_users_server`
	// key doesn't get redacted.

	input := []byte(`{"content":{"avatar_url":"mxc://something/somewhere","displayname":"Someone","join_authorised_via_users_server":"@someoneelse:somewhere.org","membership":"join"},"origin_server_ts":1633108629915,"sender":"@someone:somewhere.org","state_key":"@someone:somewhere.org","type":"m.room.member","unsigned":{"age":539338},"room_id":"!someroom:matrix.org"}`)
	expectedv8 := []byte(`{"sender":"@someone:somewhere.org","room_id":"!someroom:matrix.org","content":{"membership":"join"},"type":"m.room.member","state_key":"@someone:somewhere.org","origin_server_ts":1633108629915}`)
	expectedv9 := []byte(`{"sender":"@someone:somewhere.org","room_id":"!someroom:matrix.org","content":{"membership":"join","join_authorised_via_users_server":"@someoneelse:somewhere.org"},"type":"m.room.member","state_key":"@someone:somewhere.org","origin_server_ts":1633108629915}`)

	redactedv8, err := redactEvent(input, RoomVersionV8)
	if err != nil {
		t.Fatal(err)
	}

	redactedv9, err := redactEvent(input, RoomVersionV9)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(redactedv8, expectedv8) {
		t.Fatalf("room version 8 redaction produced unexpected result\nexpected: %s\ngot: %s", string(expectedv8), string(redactedv8))
	}

	if !bytes.Equal(redactedv9, expectedv9) {
		t.Fatalf("room version 9 redaction produced unexpected result\nexpected: %s\ngot: %s", string(expectedv9), string(redactedv9))
	}
}
