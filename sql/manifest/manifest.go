package manifest

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// migrationSources is the single source of truth for migration execution order.
// Keep entries in dependency order, regardless of which pkg directory owns the SQL file.
var migrationSources = []string{
	// Models: foundational types.
	"pkg/models/transfer/transfer_type.sql",
	"pkg/models/trade_order/order_fill.sql",
	"pkg/models/trade_order/order_side.sql",
	"pkg/models/trade_order/order_type.sql",
	"pkg/models/trade_order/trade_order_status.sql",
	"pkg/models/trade_order/order_time_in_force.sql",
	"pkg/models/transfer/ledger_entry_type.sql",

	// Models: core tables.
	"pkg/models/app_entity/app_entity.sql",
	"pkg/models/currency/currency.sql",
	"pkg/models/currency_account/currency_account.sql",
	"pkg/models/transfer/transfer.sql",
	"pkg/models/fee/fee.sql",
	"pkg/models/instrument/instrument.sql",
	"pkg/models/order_book/price_level.sql",
	"pkg/models/instrument_account/instrument_account.sql",
	"pkg/models/instrument_account/instrument_account_holding.sql",
	"pkg/models/instrument_account/instrument_account_transfer.sql",
	"pkg/models/trade_order/trade_order.sql",
	"pkg/models/trade/trade.sql",
	"pkg/models/order_book/book_order.sql",
	"pkg/models/trade_order/stop_order.sql",

	// Services: shared helpers and workflows.
	"pkg/services/order_book/banker_round.sql",
	"pkg/services/transfer/create_currency_account.sql",
	"pkg/services/transfer/create_transfer.sql",
	"pkg/services/transfer/process_transfer.sql",
	"pkg/services/registration/create_client.sql",
	"pkg/services/order_book/get_crossing_limit_orders.sql",
	"pkg/services/order_book/get_available_limit_volume.sql",
	"pkg/services/order_book/get_best_limit_price.sql",
	"pkg/services/matching/get_fill_type.sql",
	"pkg/services/matching/create_trade.sql",
	"pkg/services/matching/create_book_order.sql",
	"pkg/services/trade_order/cancel_trade_order.sql",
	"pkg/services/matching/activate_crossing_stop_orders.sql",
	"pkg/services/matching/process_crossing_stop_orders.sql",
	"pkg/services/order_book/get_trade_price.sql",
	"pkg/services/trade_order/process_trade_order.sql",
	"pkg/services/order_book/update_price_level.sql",
	"pkg/services/order_book/get_potential_self_trade_volume.sql",
	"pkg/services/transfer/create_instrument_account_transfer.sql",
	"pkg/services/order_book/get_available_market_volume.sql",
	"pkg/services/order_book/get_market_orders.sql",

	// FIX persistence.
	"pkg/fix/sessions.sql",
	"pkg/fix/messages.sql",
	"pkg/fix/messages_log.sql",
	"pkg/fix/event_log_table.sql",
}

var envSeedSources = map[string][]string{
	"DEV": {
		"pkg/conf/seeds_dev.sql",
	},
}

func MigrationSources() []string {
	sources := make([]string, 0, len(migrationSources)+len(envSeedSources[currentEnv()]))
	sources = append(sources, migrationSources...)
	sources = append(sources, envSeedSources[currentEnv()]...)
	return sources
}

func GeneratedMigrationName(index int, source string) string {
	name := strings.TrimSuffix(filepath.ToSlash(source), filepath.Ext(source))
	name = strings.ReplaceAll(name, "/", "_")
	return fmt.Sprintf("%04d_%s.sql", index+1, name)
}

func currentEnv() string {
	return strings.ToUpper(strings.TrimSpace(os.Getenv("ENV")))
}
