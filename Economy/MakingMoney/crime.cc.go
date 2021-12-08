{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Command`
	Trigger: `crime`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}


{{$a := ""}}
{{$cash := ""}}
{{$b := .User.ID}}
{{$min := ""}}
{{$max := ""}}
{{$failRate := ""}}
{{$symbol := ""}}
{{if not (dbGet $b "EconomyInfo")}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
    {{$min = $a.min}}
    {{$max = $a.max}}
    {{$failRate = $a.failRate}}
    {{$symbol = $a.symbol}}
    {{with (dbGet $b "EconomyInfo")}}
        {{$a = sdict .Value}}
        {{$cash = $a.cash}}
        {{$pay := (randInt (toInt $min) (toInt $max))}}
        {{$loss := (randInt (toInt $min) (toInt $max))}}
        {{$newCash := ""}}
        {{$int := randInt $failRate}}
        {{if gt $int (div $failRate 2)}}
            {{$newCash = (add (toInt $cash) (toInt $pay))}}
            {{$crimeEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You broke the law for a pretty penny! You made " $symbol $pay " in your crime spree today")
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
            {{sendMessage nil $crimeEmbed}}
            {{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
            {{$sdict.Set "cash" $newCash}}
            {{dbSet $b "EconomyInfo" $sdict}}
        {{else}}
            {{$newCash := (sub (toInt $cash) (toInt $loss))}}
            {{$crimeEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You broke the law to try squander a pretty penny! You were caught and lost " $symbol $loss " in your crime spree today")
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
            {{sendMessage nil $crimeEmbed}}
            {{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
            {{$sdict.Set "cash" $newCash}}
            {{dbSet $b "EconomyInfo" $sdict}}
        {{end}}
    {{end}}
{{end}}