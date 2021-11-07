![Build status](https://github.com/tolyo/open-outcry/actions/workflows/elixir.yml/badge.svg)

# Open Outcry

Open Outcry is a multi-asset [matching](https://en.wikipedia.org/wiki/Order_matching_system) and trading engine for market places of all sizes. As a
web-application, written in [Elixir](https://elixir-lang.org/) and [Phoenix](https://www.phoenixframework.org/),
it can be used and integrated in any context requiring an exchange of assets between two or more parties. Its
potential use cases range from small electronic exchanges and cryptoexchages to currency brokers and trading
simulation.

## Rationale

There are plenty of matching engines engines that can be found in open-source and all of them are based around 
a variation of [same data structure](https://link.springer.com/chapter/10.1007/978-1-4302-0147-2_2), consisting of a [TreeMap](https://docs.oracle.com/javase/8/docs/api/java/util/TreeMap.html) with keys for prices and a [Queue](https://docs.oracle.com/javase/7/docs/api/java/util/Queue.html) of orders for values. These
solutions put microsecond performance at the forefront of their productivity and leave the developer on their
own to embed this in-memory data structure into their greater application.

## Development

To start your Phoenix server:

  * Install dependencies with `mix deps.get`
  * Create and migrate your database with `mix ecto.setup`
  * Start Phoenix endpoint with `mix phx.server` or inside IEx with `iex -S mix phx.server`

Now you can visit [`localhost:4000`](http://localhost:4000) from your browser.

Ready to run in production? Please [check our deployment guides](https://hexdocs.pm/phoenix/deployment.html).

## Learn more

  * Official website: https://www.phoenixframework.org/
  * Guides: https://hexdocs.pm/phoenix/overview.html
  * Docs: https://hexdocs.pm/phoenix
  * Forum: https://elixirforum.com/c/phoenix-forum
  * Source: https://github.com/phoenixframework/phoenix
