package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadFile(t *testing.T) {
	// Non existant file
	_, err := LoadFile("test_data/non-existant")
	if err == nil {
		t.Fatal("non-existant config file failed without error")
	}

	// Invalid yaml file
	_, err = LoadFile("test_data/invalid_yaml")
	if err == nil {
		t.Fatal("invalid_yaml config file failed without error")
	}

	// TODO test non existant menu in action menu entries, and duplicate, and that anonymous and logged_in exist
	// TODO test non existant game in play actions, and duplicate
	//menuEntry = MenuEntry{
	//Key:    "p",
	//Label:  "play non existant game",
	//Action: "play nonexistant",
	//}
	//if err := menuEntry.validate(); err == nil {
	//t.Fatal("An inexistant game cannot be played")
	//}

	t.Cleanup(func() { os.RemoveAll("var/") })
	// Invalid App example
	if _, err := LoadFile("test_data/invalid_app.yaml"); err == nil {
		t.Fatal("Invalid App entry should fail to load")
	}
	// Not enough menus example
	if _, err := LoadFile("test_data/not_enough_menus.yaml"); err == nil {
		t.Fatal("not enough menu entries should fail to load")
	}
	// Invalid Menus example
	if _, err := LoadFile("test_data/invalid_menus.yaml"); err == nil {
		t.Fatal("Invalid menu entry should fail to load")
	}
	// no anonymous Menu example
	if _, err := LoadFile("test_data/no_anonymous_menu.yaml"); err == nil {
		t.Fatal("Invalid menu entry should fail to load")
	}
	// no logged_in Menu example
	if _, err := LoadFile("test_data/no_logged_in_menu.yaml"); err == nil {
		t.Fatal("Invalid menu entry should fail to load")
	}

	// Complexe example
	config, err := LoadFile("../example/complete.yaml")
	want := Config{
		App: App{
			WorkingDirectory:  "var/",
			MaxUsers:          512,
			AllowRegistration: true,
			MaxNickLen:        15,
			MenuMaxIdleTime:   600,
			PostLoginCommands: []string{
				"mkdir %w/userdata/%u",
				"mkdir %w/userdata/%u/dumplog",
				"mkdir %w/userdata/%u/ttyrec",
			},
		},
		Menus: map[string]Menu{
			"anonymous": Menu{
				Banner:  "Shell Game Launcher - Anonymous access%n======================================",
				XOffset: 5,
				YOffset: 2,
				MenuEntries: []MenuEntry{
					MenuEntry{
						Key:    "l",
						Label:  "login",
						Action: "login",
					},
					MenuEntry{
						Key:    "r",
						Label:  "register",
						Action: "register",
					},
					MenuEntry{
						Key:    "w",
						Label:  "watch",
						Action: "watch_menu",
					},
					MenuEntry{
						Key:    "q",
						Label:  "quit",
						Action: "quit",
					},
				},
			},
			"logged_in": Menu{
				Banner:  "Shell Game Launcher%n===================",
				XOffset: 5,
				YOffset: 2,
				MenuEntries: []MenuEntry{
					MenuEntry{
						Key:    "p",
						Label:  "play Nethack 3.7",
						Action: "play nethack3.7",
					},
					MenuEntry{
						Key:    "o",
						Label:  "edit game options",
						Action: "menu options",
					},
					MenuEntry{
						Key:    "w",
						Label:  "watch",
						Action: "watch",
					},
					MenuEntry{
						Key:    "r",
						Label:  "replay",
						Action: "replay",
					},
					MenuEntry{
						Key:    "c",
						Label:  "change password",
						Action: "passwd",
					},
					MenuEntry{
						Key:    "m",
						Label:  "change email",
						Action: "chmail",
					},
					MenuEntry{
						Key:    "q",
						Label:  "quit",
						Action: "quit",
					},
				},
			},
		},
		Games: map[string]Game{
			"nethack3.7": Game{
				ChrootPath: "/opt/nethack",
				FileMode:   "0666",
				ScoreCommands: []string{
					"exec /games/nethack -s all",
					"wait",
				},
				Commands: []string{
					"cp /games/var/save/%u%n.gz /games/var/save/%u%n.gz.bak",
					"exec /games/nethack -u %n",
				},
				Env: map[string]string{
					"NETHACKOPTIONS": "@%ruserdata/%n/%n.nhrc",
				},
			},
		},
	}
	if err != nil || !reflect.DeepEqual(want, config) {
		t.Fatalf("complete example failed:\nerror %v\nwant:%+v\ngot: %+v", err, want, config)
	}
}
