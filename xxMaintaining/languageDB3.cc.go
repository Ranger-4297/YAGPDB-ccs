{{$languageDB := (dbGet 0 "languageDB").Value}}
{{$languages := sdict
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