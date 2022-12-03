![Build status](https://github.com/tolyo/open-outcry/actions/workflows/elixir.yml/badge.svg)

<p align="center">
  <a target="_blank" rel="noopener noreferrer">
    <img src="https://raw.githubusercontent.com/tolyo/open-outcry/main/assets/market.jpg" width="400">
  </a>

# Open Outcry

</p>

Open Outcry is a multi-asset [matching and trading engine](https://en.wikipedia.org/wiki/Order_matching_system) for market places of all sizes. Written in [Golang](https://go.dev/) and [PL/pgSQL](https://www.postgresql.org/docs/14/plpgsql.html)
it can be used any context requiring an exchange of assets between two or more parties, including small electronic exchanges, crypto-exchages, currency brokers or trading simulators.

## Rationale

There are plenty of [available](https://github.com/Laffini/Java-Matching-Engine)
[open](https://github.com/enewhuis/liquibook) [source](https://www.opensourceagenda.com/projects/exchange-core) matching engines that are based around a variation of [same data structure](https://link.springer.com/chapter/10.1007/978-1-4302-0147-2_2), consisting of a [TreeMap](https://docs.oracle.com/javase/8/docs/api/java/util/TreeMap.html) with keys for prices and a [Queue](https://docs.oracle.com/javase/7/docs/api/java/util/Queue.html) of orders for values. These solutions put microsecond performance at the forefront of their productivity, leaving open [the non-trivial management](https://martinfowler.com/articles/lmax.html#KeepingItAllInMemory) of this in-memory data structure up to greater application. This approach may make sense in the context of liquidity pools and large securities exchanges where participants are given priority access to the order book through [DMAs](https://www.investopedia.com/terms/d/directmarketaccess.asp). In the context of a small crypto-exchange, however, where every order must be validated against an account balance held in a traditional RDBMS, this achitecture provides questionable benefits as the order processing capacity will never exceed that of the database.
Futhermore, this architecture can actually harm an active trader because a matching engine, burdened by a multi-step process of syncing and validating all of its moving parts, must necessarily freeze funds during settlement while the market moves away from the price of the executed trade. A true performance of a matching engine must measure the entire trading cycle of funds allocation between users' accounts.
These problems are fundamental, to say nothing of technical ones like: How do we ensure ACID properties of trading transactions? How do we ensure zero-downtime? Hot-code upgrades? How do we scale for unknown number of clients, connected to our trading system?

## Solution

Open Outcry puts performance and correctness of the entire trading cycle as its priority. It minimizes the number
of moving parts by putting all the trading logic into optimized PostreSQL procedures. Clients are ensured stable,
scalable and fault-tolerant access to the database through a cloud-native server, which listens to events from the database propagated to an Kubernetes cluster. This approach allows trading and settlement to be processed by a single transactional database call with event notifications delivered directly to a client without resorting to routing via a message broker.

Open Outcry's reliance on SQL also means that it can focus on business logic to provide the most feature-complete, tested and accurate trading engine, capable of evolving along with future developments in financial technology. These include marging trading, short orders, futures and options, pro-rata amongst many allocation algorithms, and hop-trades (cross-matching) where more than two parties accross several instruments are involved.

## Current features

- Time/price priority allocation
- Regular and fiat instruments
- Market and limit orders
- Stop loss and stop limit orders
- GTC, FOK, IOK, GTD, GTT orders
- Trading and payment accounts
- Self-trade prevention

## Planned features

- Peg orders
- Configurable fees
- REST API
- Websocket and FIX client connection
- Margin trading accounts
- Short orders`
- Futures and options
- Pro-rata allocation
- Multi-instrument matching

## Contributions welcome
