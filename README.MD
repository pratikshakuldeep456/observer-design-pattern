
Observer Design Pattern Example
This program demonstrates the Observer Design Pattern, which is used to establish a one-to-many dependency between objects. When the state of the subject changes, all its dependent observers are notified automatically.

Overview
In this example:

Subject: Represents an Item (e.g., "Shirt") that notifies its observers when it becomes available.
Observers: Represent Customer objects who subscribe to updates about the item's availability.
When the item's state changes (e.g., it becomes available), all registered customers are notified.