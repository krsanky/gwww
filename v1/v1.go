package v1

import (
	"net/http"

	"oldcode.org/gow/lg"
	"oldcode.org/gow/web"
)

var tab = `
Intro
-----

Am   C    Em   D
Am   C    Em   D

Verse
-----

Am   C    Em .........repeat 4X
Am   C    Em   D .....repeat 3X
Am   C    Em .........repeat 1X

Bridge
------

D   Dm7         G(D bass)
D   E(D bass)   G(D bass)   Dm7

-------------------------------------------------------

This whole progression is then repeated.
All chords are open position.
To play the Dm7, slide the open D position up to the 5th fret.
To play the G(D bass) slide open D to 7th fret
To play the E(D bass) slide open D to 4th fret.

-------------EOF------------------------------------------
`

func Index(w http.ResponseWriter, r *http.Request) {
	//func RenderPage(w http.ResponseWriter, page string, data interface{}) {
	lg.Log.Printf("v1/index.....")
	web.RenderPage(w, "v1/index", nil)
}

var tab2 = `
[Intro] (x2) 

Am - C – Em – D



[Verse]

Am                 C

A strange kind of love

Em                            

A strange kind of feeling

Am      C           Em

Swims through your eyes

Am           Fmaj7

And like the doors

     Em

To a wide vast dominion

Am       Fmaj7      Em

They open to your prize

Am           C

This is no terror ground

Em                 D

Or place for the rage

Am         C

No broken hearts

            Em      D

White wash lies

Am                    C

Just a taste for the truth

         Em                 D

Perfect taste choice and meaning

Am       C        Em

A look into your eyes



[Bridge]      

D – Dm7 (D 5th) – G (D 7th) 

D – E (D 4th) – G (D 7th) – Dm7 (D 5th)



[Verse]

Am           C          Em

Blind to the gemstone alone

(Em)

A smile from a frown circles round

Am         C                Em

Should he stay or should he go

Am              Fmaj7

Let him shout a rage so strong

  Em

A rage that knows no right or wrong

Am         Fmaj7           Em

And take a little piece of you

Am                   C          

There is no middle ground

Em                 D

Or that's how it seems

Am         C         Em       D

For us to walk or to take

Am                 C

Instead we tumble down

       Em            D

Either side left or right

Am      C      Em

To love or to hate 



[Outro]   

D – Dm7 (D 5th) – G (D 7th) 

D – E (D 4th) – G (D 7th) – Dm7 (D 5th)

Am – C – Em – D              (x4)

Am – Fmaj7 – Em – D          (x2)
`
