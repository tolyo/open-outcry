![Build status](https://github.com/tolyo/open-outcry/actions/workflows/elixir.yml/badge.svg)

# Open Outcry

Open Outcry is a multi-asset [matching](https://en.wikipedia.org/wiki/Order_matching_system) and trading engine for market places of all sizes. As a
web-application, written in [Elixir](https://elixir-lang.org/) and [Phoenix](https://www.phoenixframework.org/),
it can be used and integrated in any context requiring an exchange of assets between two or more parties. Its
potential use cases range from small electronic exchanges and cryptoexchages to currency brokers and trading
simulation.

## Rationale

There are plenty of matching engines that [can](https://github.com/Laffini/Java-Matching-Engine) 
[be](https://github.com/enewhuis/liquibook) [found](https://www.opensourceagenda.com/projects/exchange-core) in open-source that are based around 
a variation of [same data structure](https://link.springer.com/chapter/10.1007/978-1-4302-0147-2_2), consisting of a [TreeMap](https://docs.oracle.com/javase/8/docs/api/java/util/TreeMap.html) with keys for prices and a [Queue](https://docs.oracle.com/javase/7/docs/api/java/util/Queue.html) of orders for values. These
solutions put microsecond performance at the forefront of their productivity, leaving open [the non-trivial management](https://martinfowler.com/articles/lmax.html#KeepingItAllInMemory) of this in-memory data structure up to greater application. This approach may make sense in the context of a large securities exchange, where
trading and settlement are separated into different contexts and some market participants are given priority
access to the order book through DMAs. In the context of a small crypto-exchange, however, where every order must be validated against an account balance held in a traditional database, this achitecture makes no sense 
as your order processing capacity will never exceed that of your database.  

Futhermore, these solutions fail to consider the needs of an active trader for release of funds after a match for subsequent trading. Closing a LONG position implies a trader's assumption that a market has 
reversed its bullish trend and potential desire to open a SHORT position in a 
financial instrument. A matching engine, burdened by a multi-step process of syncing and validating all of its moving parts, will only penalize such traders by freezing their funds during settlement as the market moves away from the price of the executed trade. In other words, a true thoroughput of a matching engine must reflect the flow rate of instruments and funds between users' accounts - the core of trading itself.

These problems are fundamental, to say nothing of technical ones like: How do we ensure ACID properties of trading transactions? How do we ensure zero-downtime? Hot-code upgrades? How do we scale for unknown number of clients, connected to our trading system?


## Development

To start your Phoenix server:

  * Install dependencies with `mix deps.get`
  * Create and migrate your database with `mix ecto.setup`
  * Start Phoenix endpoint with `mix phx.server` or inside IEx with `iex -S mix phx.server`

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.

Ready to run in production? Please [check our deployment guides](https://hexdocs.pm/phoenix/deployment.html).