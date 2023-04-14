{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Leave message in channel`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* 
Use this in conjunction with 'joinMessageAlt'
To retrieve a users economy data upon rejoining
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Response */}}
{{$cash := or (dbGet .User.ID "cash").Value 0 | toInt}}
{{$bank := or (((dbGet 0 "bank").Value).Get (toString $user.ID)) 0 | toInt}}
{{$economyData := or ($a := (dbGet $userID "userEconData").Value) ($a := sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
{{with or ($a := (dbGet 0 "EconomyInfoLeftGuild").Value) sdict}}
	{{$a.Set (toString $.User.ID) (sdict "cash" $cash "bank" $bank "data" $economyData)}}
	{{dbSet 0 "EconomyInfoLeftGuild" $a}}
	{{if $cash}}
		{{dbDel $.User.ID "cash"}}
	{{end}}
	{{if $bank}}
		{{$bankDB := (dbGet 0 "bank").Value}}
		{{$bankDB.Del (toString $.User.ID)}}
		{{dbSet 0 "bank" $bankDB}}
	{{end}}
	{{if $economyData}}
		{{dbDel $.User.ID "userEconData"}}
	{{end}}
{{end}}
