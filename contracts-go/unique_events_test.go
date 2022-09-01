package contracts_go

import (
	"fmt"
	"sort"
	"testing"

	v001 "github.com/eluv-io/contracts-evm-builds/contracts-go/v0.0.1/contracts"

	"github.com/eluv-io/contracts-evm-builds/contracts-go/events"
)

// TestDuplicateEvents does not look at tradable
func TestDuplicateEvents(t *testing.T) {
	t.Skip("not meaningful ")
	evLast := map[string]*events.EventInfo{} // should be the actual last
	_ = evLast
	ev001 := v001.UniqueEvents

	//trLast := elv_tradable.UniqueEvents
	//fmt.Println("",
	//	"201903", len(ev201903),
	//	"202002", len(ev202002),
	//	"202008", len(ev202008),
	//	"latest", len(evLast),
	//	"tradable", len(trLast),
	//	"all", len(UniqueEvents))
	/*
		201903: 26 only in this version
		202002:  3 duplicates of other versions
		202008:  1 only in this version (CreateExtUserWallet)
		latest: 81
		trad:   15
		all:   126
	*/

	getId := func(version, eventName string) (string, bool) {
		switch version {
		case "ev001":
			if e, ok := ev001[eventName]; ok {
				return e.ID.String(), true
			}
		}
		return "", false
	}
	_ = getId

	//for name, ev := range evLast {
	//for _, version := range []string{"v100"} {
	//	if id, ok := getId(version, name); ok && id != ev.ID.String() {
	//		switch name {
	//		case "RunAccessCharge", "VersionConfirm", "AccessRequest":
	//			// RunAccessCharge and VersionConfirm are known to have changed
	//		default:
	//			t.Errorf("id mismatch - version %s, event %s", version, name)
	//		}
	//	}
	//}
	//}
}

func TestUniqueEvents(t *testing.T) {
	var names []string
	for name := range UniqueEvents {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println(fmt.Sprintf("all events (%d)", len(UniqueEvents)))
	for _, name := range names {
		event := UniqueEvents[name]
		for _, eventType := range event.Types {
			fmt.Println(fmt.Sprintf("%-35s", name),
				fmt.Sprintf("%-50v", eventType.Type),
				event.ID.String())

			byName := EventsByName[name]
			if byName != event {
				t.Errorf("event %s by name differs - \n\texpected: %+v\n\tactual:   %+v", name, event, byName)
			}

			byID := EventsByID[event.ID]
			if byID != event {
				t.Errorf("event %s by ID differs - \n\texpected: %+v\n\tactual:   %+v", name, event, byID)
			}

			byType := EventsByType[eventType.Type]
			if byType != event {
				t.Errorf("event %s by type differs - \n\texpected: %+v\n\tactual:   %+v", name, event, byType)
			}

			if EventNamesByID[event.ID] != event.Name {
				t.Errorf("event %s name by ID differs - ID %s name %s", name, event.ID.String(), event.Name)
			}
		}
	}
}

func TestEventsByType(t *testing.T) {
	for evt, event := range EventsByType {
		fmt.Println(
			fmt.Sprintf("%-50v", evt),
			fmt.Sprintf("%+v", event))
	}
}
