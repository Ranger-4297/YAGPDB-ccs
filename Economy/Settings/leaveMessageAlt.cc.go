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
{{with dbGet .User.ID "EconomyInfo"}}
	{{$a := sdict .Value}}
	{{if not (dbGet 0 "EconomyInfoLeftGuild")}}
		{{dbSet 0 "EconomyInfoLeftGuild" sdict}}
	{{end}}
	{{with dbGet 0 "EconomyInfoLeftGuild"}}
		{{$entry := sdict .Value}}
		{{$entry.Set (toString $.User.ID) (sdict $a)}}
		{{dbSet 0 "EconomyInfoLeftGuild" $entry}}
		{{dbDel $.User.ID "EconomyInfo"}}
	{{end}}
{{end}}