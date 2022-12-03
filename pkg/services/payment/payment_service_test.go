# funcmodule PaymentServiceTest {
#   use DataCase

#   require Logger

#   test "deposit/4" {
#     # given a customer
#     pub_id = TestUtils.create_client()
#     assert DBTestUtils.get_count("payment") == 0

#     # when amount is deposited
#     PaymentService.deposit(pub_id, 10.00, "EUR", "BANK", "REF123")

#     # then amount should increase and payment should be created
#     assert 10.00 == PaymentAccount.get_balance(pub_id, "EUR")
#     assert 10.00 == PaymentAccount.get_available_balance(pub_id, "EUR")
#     assert 0.00 == PaymentAccount.get_reserved_balance(pub_id, "EUR")
#     assert DBTestUtils.get_count("payment") == 1

#     # when another payment is deposited
#     PaymentService.deposit(pub_id, 10.00, "EUR", "BANK", "REF1234")

#     # then amount should increase and another payment should be created
#     assert 20.00 == PaymentAccount.get_balance(pub_id, "EUR")
#     assert 20.00 == PaymentAccount.get_available_balance(pub_id, "EUR")
#     assert 0.00 == PaymentAccount.get_reserved_balance(pub_id, "EUR")
#     assert DBTestUtils.get_count("payment") == 2
#   }

#   test "transfer/4" {
#     # given two customers
#     pub_id = TestUtils.create_client()
#     pub_id2 = TestUtils.create_client2()
#     PaymentService.deposit(pub_id, 10.00, "EUR", "BANK", "REF1234")
#     assert DBTestUtils.get_count("payment") == 1
#     assert 10.00 == PaymentAccount.get_balance(pub_id, "EUR")
#     assert 0.00 == PaymentAccount.get_balance(pub_id2, "EUR")

#     # when one user makes a payment to another
#     PaymentService.transfer(pub_id, pub_id2, 10.00, "EUR", "Hey hey!")

#     # then amount should change and a payment should be created
#     assert DBTestUtils.get_count("payment") == 2

#     assert 0.00 == PaymentAccount.get_balance(pub_id, "EUR")
#     assert 10.00 == PaymentAccount.get_balance(pub_id2, "EUR")

#     # when insuffient funds
#     # then should raise an error
#     assert_raise Postgrex.Error, fn ->
#       PaymentService.transfer(pub_id, pub_id2, 1000.00, "EUR", "Hey hey!")
#     }
#   }
# }
