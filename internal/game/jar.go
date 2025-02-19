package game

// A Jar is an object that contains a die.
type Jar struct {
	Berry  Berry
	Locked bool
	Rolled bool
}

// Roll will roll the berry value if a Jar is unlocked, and leave the berry value unchanged if it is locked.
func (j *Jar) Roll() {
	if j.Rolled && j.Locked {
		return
	}

	j.Berry = doRoll()
	j.Rolled = true
}

// Lock will lock the jar
func (j *Jar) Lock() {
	j.Locked = true
}

// Unlock will unlock the jar
func (j *Jar) Unlock() {
	j.Locked = false
}

// Reset will reset a Jar so that it hasn't been rolled and is unlocked.
func (j *Jar) Reset() {
	j.Rolled = false
	j.Locked = false
}

func (j *Jar) String() string {
	if !j.Rolled {
		return "EMPTY"
	}

	if j.Locked {
		return j.Berry.String() + "🛑"
	}

	return j.Berry.String() + "✅"
}