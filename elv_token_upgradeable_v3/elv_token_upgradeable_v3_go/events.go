package elv_token_upgradeable_v3_go

import (
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/common"

	"github.com/eluv-io/contracts-evm-builds/elv_token_upgradeable_v3/elv_token_upgradeable_v3_go/events"
    "github.com/eluv-io/contracts-evm-builds/elv_token_upgradeable_v3/elv_token_upgradeable_v3_go/v0.0.0-dev"
)

var (
	// UniqueEvents contains all events keyed by their name (e.g. "AccessRequest"). Events that have different
	// signatures in different contract versions (sub-packages) are treated as separate events. Their most recent
	// version will get the event's original name, while older versions are suffixed with the version date of the
	// contract package, e.g. VersionConfirm_20200206.
	UniqueEvents = make(map[string]*events.EventInfo)
	// EventsByName is an alias of UniqueEvents
	EventsByName = UniqueEvents
	// EventsByType contains all events keyed by their go type.
	EventsByType = make(map[reflect.Type]*events.EventInfo)
	// EventsByID contains all events keyed by their id (topic hash)
	EventsByID = make(map[common.Hash]*events.EventInfo)
	// EventNamesByID contains all event names keyed by their id (topic hash)
	EventNamesByID = make(map[common.Hash]string)

	allEventInfos = []packageEvent{ 
		{ elv_token_upgradeable_v3_v0_0_0_dev.UniqueEvents, "v000_dev" },
	}
)

type packageEvent struct {
	events map[string]*events.EventInfo
	suffix string
}

func init() {
	addEvents := func(
		events map[string]*events.EventInfo,
		eventSuffix string) {

		for name, event := range events {
			ev, ok := UniqueEvents[name]
			if !ok {
				UniqueEvents[name] = event
				EventsByID[event.ID] = event
				continue
			}
			if ev.ID == event.ID || EventsByID[event.ID] != nil {
				continue
			}
			UniqueEvents[fmt.Sprintf("%s_%s", name, eventSuffix)] = event
			EventsByID[event.ID] = event
		}
	}

	for _, packageEvents := range allEventInfos {
		addEvents(packageEvents.events, packageEvents.suffix)
	}

	for name, event := range UniqueEvents {
		for _, eventType := range event.Types {
			evt := eventType.Type
			ev, _ := EventsByType[evt]
			if ev != nil {
				panic(fmt.Sprintf("duplicate events for type %v: %s <-> %s",
					evt, ev.Name, name))
			}
			EventsByType[evt] = event
		}
		EventNamesByID[event.ID] = event.Name
	}

	// above might have left out some types - ensure that all event types are set
	for _, pes := range allEventInfos {
		ues := pes.events
		for _, event := range ues {
			for _, eventType := range event.Types {
				evt := eventType.Type
				ev, _ := EventsByType[evt]
				if ev == nil {
					byId, _ := EventsByID[event.ID]
					if byId == nil {
						panic(fmt.Sprintf("no event in EventsByID for ID %s",
							event.ID.String()))
					}
					if !byId.HasType(eventType) {
						byId.Types = append(byId.Types, eventType)
					}
					EventsByType[evt] = byId
				}
			}
		}
	}
}
