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
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{with dbGet 0 "EconomyInfoLeftGuild"}}
		{{$a := sdict .Value}}
		{{if ($a.Get (toString $.User.ID))}}
			{{dbSet $.User.ID "EconomyInfo" ($a.Get (toString $.User.ID))}}
			{{$a.Del (toString $.User.ID)}}
			{{dbSet 0 "EconomyInfoLeftGuild" $a}}
		{{end}}
	{{else}}
		{{$startBalance := (toInt $a.startBalance)}}
		{{dbSet .User.ID "EconomyInfo" (sdict "cash" $startBalance "bank" 0)}}
	{{end}}
{{end}}