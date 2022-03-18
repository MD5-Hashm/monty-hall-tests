# monty-hall-tests

### The problem
"Suppose you're on a game show, and you're given the choice of three doors: Behind one door is a car; behind the others, goats. You pick a door, say No. 1, and the host, who knows what's behind the doors, opens another door, say No. 3, which has a goat. He then says to you, "Do you want to pick door No. 2?" Is it to your advantage to switch your choice?"

### My Explonation
When you pick your door there is a 33.3% chance that each door is the car. The group of the other 2 doors that you didn't pick have a 66.6% chance of conataing the car. When the host reveals door 3 showing that it is a goat, door 2 is still a 66.6% chance of being a car. Obviously door 2 in this case deems you a better chance of winning and this program demonstates that.

### Whats this program shows
This program runs through the problem with your guess being completely random and the car door being completely random. The default run through amount is 100000 times but you can specify a custom one by passing it in as an arg. This program then prints out the result of the simulation.
