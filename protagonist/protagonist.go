package protagonist

type Player struct {
	fname			string
	lname			string
	occupation		string
	intelligence	int
	knowledge		int
	occult			int
	observation		int
	deduction		int
	induction		int
	abduction		int
	dexterity		int
	strength		int
	stamina			int
	speed			int
}

func NewPlayer(fn string, ln string, occup string, intl int, know int, occ int, obsv int, ded int, ind int, abd int, dext int, str int, stmn int, spd int) *Player {
	p := Player{
		fname: fn, 
		lname: ln, 
		occupation: occup, 
		intelligence: intl, 
		knowledge: know, 
		occult: occ, 
		observation: obsv, 
		deduction: ded, 
		induction: ind, 
		abduction: abd, 
		dexterity: dext, 
		strength: str,
		stamina: stmn,
		speed: spd,
	}
	
	return p
}
