{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Command`
	Trigger: `Work`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$a := ""}}
{{$cash := ""}}
{{$b := .User.ID}}
{{if not (dbGet $b "EconomyInfo")}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{$min := ""}}
{{$max := ""}}
{{with (dbGet 0 "EconomySettings")}}
{{$a = sdict .Value}}
    {{$min = $a.min}}
    {{$max = $a.max}}
    {{with (dbGet $b "EconomyInfo")}}
        {{$cash = $a.cash}}
        {{$pay := (randInt (toInt $min) (toInt $max)}}
        {{$newCash := (add (toInt $cash) $pay)}}
        {{$workEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))
            "description" (print "You decided to work today! You got paid a hefty £" $pay)
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $workEmbed}}
        {{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
        {{$sdict.Set "cash" $newCash}}
        {{dbSet $b "EconomyInfo" $sdict}}
    {{end}}
{{end}}
