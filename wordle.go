package wordle

const (
	maxGuesses = 6
	wordSize   = 5
)

type wordleState struct {
	// word is the word that the user is trying to guess
	word [wordSize]byte
	// guesses holds the guesses that the user has made
	guesses [maxGuesses]guess
	// currGuess is the index of the available slot in guesses
	currGuess int
}

type letter struct {
	// char is the letter that this struct represents
	char byte
	// status is the state of the letter (absent, present, correct)
	status letterStatus
}

// guess is an attempt to guess the word
type guess [wordSize]letter

// letterStatus can be none, correct, present, or absent
type letterStatus int

const (
	// none = no status, not guessed yet
	none letterStatus = iota
	// absent = not in the word
	absent
	// present = in the word, but not in the correct position
	present
	// correct = in the correct position
	correct
)

// newWordleState builds a new wordleState from a string.
// Pass in the word you want the user to guess.
func newWordleState(word string) wordleState {
	w := wordleState{}
	copy(w.word[:], word)
	return w
}

func newLetter(oneLetter byte) letter {
	l := letter{}
	l.char = oneLetter
	l.status = none
	return l
}

func newGuess(guessWord string) guess {
	g := guess{}
	for i := 0; i < len(guessWord); i++ {
		l := letter{}
		l.char = guessWord[i]
		g[i] = l
	}
	return g
}
