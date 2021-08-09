package sink

import (
	"errors"
	"strings"

	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/state/indexer"
	"github.com/tendermint/tendermint/state/indexer/sink/kv"
	"github.com/tendermint/tendermint/state/indexer/sink/null"
	"github.com/tendermint/tendermint/state/indexer/sink/psql"
)

func EventSinksFromConfig(cfg *config.Config, dbProvider config.DBProvider, chainID string) ([]indexer.EventSink, error) {
	if len(cfg.TxIndex.Indexer) == 0 {
		return []indexer.EventSink{null.NewEventSink()}, nil
	}

	// check for duplicated sinks
	sinks := map[string]struct{}{}
	for _, s := range cfg.TxIndex.Indexer {
		sl := strings.ToLower(s)
		if _, ok := sinks[sl]; ok {
			return nil, errors.New("found duplicated sinks, please check the tx-index section in the config.toml")
		}
		sinks[sl] = struct{}{}
	}
	eventSinks := []indexer.EventSink{}
	for k := range sinks {
		switch k {
		case string(indexer.NULL):
			// When we see null in the config, the eventsinks will be reset with the
			// nullEventSink.
			return []indexer.EventSink{null.NewEventSink()}, nil

		case string(indexer.KV):
			store, err := dbProvider(&config.DBContext{ID: "tx_index", Config: cfg})
			if err != nil {
				return nil, err
			}

			eventSinks = append(eventSinks, kv.NewEventSink(store))

		case string(indexer.PSQL):
			conn := cfg.TxIndex.PsqlConn
			if conn == "" {
				return nil, errors.New("the psql connection settings cannot be empty")
			}

			es, _, err := psql.NewEventSink(conn, chainID)
			if err != nil {
				return nil, err
			}
			eventSinks = append(eventSinks, es)
		default:
			return nil, errors.New("unsupported event sink type")
		}
	}
	return eventSinks, nil

}
