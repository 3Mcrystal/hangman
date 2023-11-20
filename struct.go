package hangman

type Hangman struct {
	WordToGuess      string
	DisplayWord      string
	Attempts         int
	HangmanPositions []string
	LetterFile       string
	GuessedLetters   []rune
}
