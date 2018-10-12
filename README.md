# Cards
_go run main.go deck.go_ -> Runs GO project.
_go test_ -> Runs all tests.

# Useful Functions
_newDeck()_ -> Generates a standard deck, containing 52 cards.

_{deckname}.print()_ -> Logs the contents of the deck.

_{deckname}.toString()_ -> Converts a slice of type deck to a single string, where each element is separated by a comma.

_{deckname}.saveToFile(filename)_ -> Converts a string to a byte slice before writing the byte slice to the specified filename.

_removeFile(filename)_ -> Permanently removes a local file that matches the filename that was passed as a parameter.

_newDeckFromFile(filename)_ -> Reads in the contents of a file and attempts to convert the contents to a deck.

_deal(deck, handsize)_ -> Returns two decks. One deck will contain the amount of cards that was passed as the handsize, while the other deck will have the remaining cards.

_{deckname}.shuffle()_ -> Randomly shuffles the order of every card in the deck.

# Testing
All test functions follow GO Lang Standards, where each test file will end with "\_test" and every function name will start with "Test....".
