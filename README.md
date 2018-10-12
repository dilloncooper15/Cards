# Cards

# Useful Functions
newDeck() -> Generates a standard deck, containing 52 cards.

{deckname}.print() -> Logs the contents of the deck.

{deckname}.toString() -> Converts a slice of type deck to a single string, where each element is separated by a comma.

{deckname}.saveToFile(filename) -> Converts a string to a byte slice before writing the byte slice to the specified filename.

removeFile(filename) -> Permanently removes a local file that matches the filename that was passed as a parameter.

newDeckFromFile(filename) -> Reads in the contents of a file and attempts to convert the contents to a deck.

deal(deck, handsize) -> Returns two decks. One deck will contain the amount of cards that was passed as the handsize, while the other deck will have the remaining cards.

{deckname}.shuffle() -> Randomly shuffles the order of every card in the deck.

# Testing
All test functions follow GO Lang Standards, where each test file will end with "\_test" and every function name will start with "Test....".
