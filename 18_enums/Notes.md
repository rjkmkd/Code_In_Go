Explanation of Concepts in the Code
Custom Type Definition:

The orderStatus type is defined as string. This approach is helpful for readability and ensures type safety. Only values of type orderStatus can be assigned to variables or constants of this type, reducing errors.
Constant Grouping with Custom Type:

The constants Receved, Prepaired, and Delevered represent different statuses an order can have. Grouping constants improves organization and code readability.
Defining these statuses as constants also protects the code from invalid values, as only predefined statuses are allowed.
Alternative Integer-Based Enumeration:

By defining orderStatus as an int and using iota, each constant is automatically assigned a unique integer value. This approach is useful when using statuses as numeric codes.
Function with Custom Type Parameter:

The changeOrderStatus function demonstrates how custom types enforce type checking, as the parameter must be of type orderStatus. This improves code safety and helps ensure the function receives only valid values.
Function Call and Output:

In main, we call changeOrderStatus with Delevered as the status. This outputs a formatted message showing the status change.