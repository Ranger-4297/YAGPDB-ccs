{{$languageDB := (dbGet 0 "languageDB").Value}}
{{$languages := sdict
"italian" (dict
	1 "Nessuna alleanza? Digita \"skip\""
	2 "Eccezionale! <@!> Hai appena ricevuto il tuo Tag Server! Ora procedi a <#> <1/5>"
	3 "<@!> Inserisci qui il tag della tua alleanza <2/5>"	
	4 "È fantastico! <@!> ora hai un tag dell'Alleanza! Continua e procedi con <#> <2/5>"
	5 "<@!> digita il nome del gioco del tuo personaggio qui! <3/5>"
	6 "Perfetto! <@!> hai aggiornato il tuo nome visualizzato con il nome del gioco del tuo personaggio! Altri due passaggi, procedi a <#> <3/5>"
	7 "<@!> seleziona il grado della tua alleanza nel gioco <4/5>"
	8 "Impossibile aggiungere una reazione all'utente che ha bloccato il bot. Soprannome aggiornato"
	9 "Inserisci un tag di 3-4 cifre"
	10 "Inserisci un tag numerico"
	11 "Inserisci un tag di alleanza di 3-4 caratteri"
	12 "Per favore non usare caratteri speciali"
	13 "Inserisci un nome utente di 3-15 caratteri"	
)
"romanian" (dict
	1 "Fără alianță? Tastați \"skip\""
	2 "Minunat! <@!> Tocmai ați primit eticheta de server! Acum treceți la <#> <1/5>"
	3 "<@!> Introdu eticheta de alianță aici <2/5>"
	4 "Aceasta este cool! <@!> acum aveți o etichetă de alianță! Continuați și treceți la <#> <2/5>"
	5 "<@!> introduceți numele jocului dvs. de personaj aici! <3/5>"
	6 "Perfect! <@!> ți-ai actualizat numele afișat la Numele jocului de personaj! Încă doi pași, treceți la <#> <3/5>"
	7 "<@!> selectează rangul alianței în joc <4/5>"
	8 "Nu se poate adăuga reacție utilizatorului care a blocat bot. Pseudomul a fost actualizat"
	9 "Vă rugăm să introduceți o etichetă de 3-4 cifre"
	10 "Vă rugăm să introduceți o etichetă "
	11 "Vă rugăm să introduceți o etichetă de alianță de 3-4 caractere"
	12 "Vă rugăm să nu folosiți caractere speciale"
	13 "Vă rugăm să introduceți un nume de utilizator de 3-15 caractere"	
)
"danish" (dict
	1 "Ingen alliance? Skriv \"skip\""
	2 "Fantastisk! <@!> Du har lige modtaget dit server-tag! Fortsæt nu til <#> <1/5>"
	3 "<@!> Indtast dit alliance-tag her <2/5>"
	4 "Det her er sejt! <@!> du har nu et Alliance-tag! Fortsæt og fortsæt til <#> <2/5>"
	5 " <@!> skriv dit karakterspils navn her! <3/5>"
	6 "Perfekt! <@!> du har opdateret dit viste navn til dit karakterspilsnavn! To trin mere, fortsæt til <#> <3/5>"
	7 "<@!> vælg din alliancerangering i spillet <4/5>"
	8 "Kan ikke tilføje reaktion til bruger, der har blokeret bot. Kaldenavnet er opdateret"
	9 "Indtast venligst et 3-4-cifret tag"
	10 "Indtast venligst et numerisk tag"
	11 "Indtast venligst et alliancemærke på 3-4 tegn"
	12 "Brug venligst ikke specialtegn"
	13 "Indtast venligst et brugernavn på 3-15 tegn"
)
"polish" (dict
	1 "Brak sojuszu? Wpisz \"skip\""
	2 "Wspaniały! <@!> Właśnie otrzymałeś swój tag serwera! Teraz przejdź do <#> <1/5>"
	3 "<@!> Wpisz tutaj swój tag sojuszu <2/5>"
	4 "To jest fajne! <@!> masz teraz tag sojuszu! Kontynuuj i przejdź do <#> <2/5>"
	5 "<@!> wpisz tutaj swoją nazwę gry postaci! <3/5>"
	6 "Doskonały! <@!> zaktualizowałeś swoją nazwę wyświetlaną do nazwy gry postaci! Jeszcze dwa kroki i przejdź do <#> <3/5>"
	7 "<@!> wybierz swoją rangę sojuszu w grze <4/5>"
	8 "Nie można dodać reakcji do użytkownika, który zablokował bota. Pseudonim zaktualizowany"
	9 "Proszę wprowadzić 3-4 cyfrowy znacznik"
	10 "Proszę wprowadzić znacznik numeryczny"
	11 "Proszę wprowadzić tag sojuszu składający się z 3-4 znaków"
	12 "Proszę nie używać znaków specjalnych"
	13 "Wprowadź nazwę użytkownika zawierającą od 3 do 15 znaków"
)
"hebrew" (dict
	1 "אין ברית? הקלד \"skip\""
	2 "מדהים! <@!> זה עתה קיבלת את תג השרת שלך! כעת המשך ל-<#> <1/5>"
	3 "<@!> הזן את תג הברית שלך כאן <2/5>"
	4 "זה מגניב! <@!> יש לך כעת תג ברית! המשך והמשיך אל <#> <2/5>"
	5 "<@!> הקלד את שם משחק הדמות שלך כאן! <3/5>"
	6 "מושלם! <@!> עדכנת את שם התצוגה שלך לשם משחק הדמות שלך! שני שלבים נוספים, המשך אל <#> <3/5>"
	7 "<@!> בחר את דירוג הברית שלך במשחק. <4/5>"
	8 "לא ניתן להוסיף תגובה למשתמש שחסם בוט. הכינוי עודכן"
	9 "נא להזין תג בן 3-4 ספרות"
	10 "נא להזין תג מספרי"
	11 "אנא הזן תג ברית בן 3-4 תווים"
	12 "נא לא להשתמש בתווים מיוחדים"
	13 "נא להזין שם משתמש בן 3-15 תווים"
)
}}
{{range $k, $v := $languages}}
	{{$languageDB.Set $k $v}}
{{end}}
{{dbSet 0 "languageDB" $languageDB}}