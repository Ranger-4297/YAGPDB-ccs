{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Join message in channel`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* 
Use this in conjunction with 'leaveMessageAlt'
To retrieve a users economy data upon rejoining
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Response */}}
{{with or ($a := (dbGet 0 "EconomyInfoLeftGuild").Value) ($a := sdict)}}
	{{if ($user := $a.Get (toString $.User.ID))}}
		{{$cash := $user.cash}}
		{{$bank := $user.bank}}
		{{$data := $user.data}}
		{{dbSet $.User.ID "cash" $cash}}
		{{dbSet $.User.ID "userEconData" $data}}
		{{$bankDB := or (dbGet 0 "bank").Value sdict}}
		{{$bankDB.Set (toString .User.ID) $bank}}
		{{dbSet 0 "bank" $bankDB}}
	{{else}}
		{{$startBalance := (dbGet 0 "EconomySettings").Value.startBalance}}
		{{dbSet $.User.ID "cash" $startBalance}}
	{{end}}
{{end}}