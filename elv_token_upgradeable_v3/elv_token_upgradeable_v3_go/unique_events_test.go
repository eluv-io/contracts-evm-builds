package elv_token_upgradeable_v3_go

import (
	"fmt"
	"sort"
	"testing"
    
    "github.com/eluv-io/contracts-evm-builds/elv_token_upgradeable_v3/elv_token_upgradeable_v3_go/v0.0.0-dev"
)

// TestDuplicateEvents does not look at tradable
func TestDuplicateEvents(t *testing.T) {
    
    ev000_dev:= elv_token_upgradeable_v3_v0_0_0_dev.UniqueEvents
	evLast := elv_token_upgradeable_v3_v0_0_0_dev.UniqueEvents

	tagsList :=  []string{ 
	    "ev000_dev",
	}

	getId := func(version, eventName string) (string, bool) {
        switch version { 
         case "ev000_dev":
           if e, ok := ev000_dev[eventName]; ok {
               return e.ID.String(), true
           }
        }
        return "", false
    }

	for name, ev := range evLast {
        for _, version := range tagsList {
            if id, ok := getId(version, name); ok && id != ev.ID.String() {
                switch name {
                case "RunAccessCharge", "VersionConfirm", "AccessRequest":
                    // RunAccessCharge and VersionConfirm are known to have changed
                default:
                    t.Errorf("id mismatch - version %s, event %s", version, name)
                }
            }
        }
    }
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
