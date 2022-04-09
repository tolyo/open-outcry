defmodule TradeOrder.Status do
  @type t ::
          :OPEN
          | :PARTIALLY_FILLED
          | :CANCELLED
          | :PARTIALLY_CANCELLED
          | :PARTIALLY_REJECTED
          | :FILLED
          | :REJECTED
end
