Explanation of Concepts in the Code
Interfaces as Contracts:

The paymenter interface defines a single method pay which takes an amount of type float32. Any struct implementing this interface must define the pay method with the exact signature.
This enables the use of different types (payment gateways) with a single interface, achieving polymorphism.
Struct payment as a Wrapper:

The payment struct has a field gateway of type paymenter, which can hold any type that implements paymenter.
The makePayment method of payment uses the gateway to process payments, making it flexible to use different payment methods.
Implementing the Interface with Multiple Structs:

razorPay, strype, fakepaymentGW, and paypal all implement the paymenter interface by defining the pay method. Each struct provides its specific logic for handling payments.
fakepaymentGW is particularly useful for testing without relying on real payment gateways.
Flexible and Reusable Payment Logic:

In main, we create an instance of payment with paypal as the payment gateway. The payment logic is reusable with any gateway, as long as it implements the paymenter interface.
Dependency Injection:

By passing different gateways to payment, we effectively use dependency injection, enabling payment to work independently of the specific payment logic.





