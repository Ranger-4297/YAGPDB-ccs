{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Command`
	Trigger: `Rob`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "rob <User:Mention/ID>" (carg "member" "target user")}}
{{$user := ($args.Get 0).User}}
{{$victim := $user.ID}}
{{$a := ""}}
{{$b := .User.ID}}
{{$c := ""}}
{{$cash := ""}}
{{$failRate := ""}}
{{if not (dbGet $b "EconomyInfo")}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a = sdict .Value}}
	{{$failRate = $a.failRate}}
    {{with (dbGet $victim "EconomyInfo")}}
        {{$victimsCash := $a.cash}}
        {{$amount := (randInt $victimsCash}}
        {{$yourNewCash := (add (toInt $victimsCash) $amount)}}
		{{$victimsNewCash := (sub (toInt $victimsCash) $amount)}}
		{{$crimeEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))
            "description" (print "You robbed £ " $amount " from <@! " $victim ">")
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
        {{sendMessage nil $crimeEmbed}}
        {{$sdict := (dbGet $victim "EconomyInfo").Value}}
        {{$sdict.Set "cash" $victimsNewCash}}
        {{dbSet $victim "EconomyInfo" $sdict}}
		{{with (dbGet $b "EconomyInfo")}}
            {{$c = sdict .Value}}
            {{$cash = $c.cash}}
            {{$newCash := (add (toInt $cash) $amount)}}
            {{$sdict2 := (dbGet $b "EconomyInfo").Value}}
            {{$sdict2.Set "cash" $newCash}}
            {{dbSet $b "EconomyInfo" $sdict2}}
        {{end}}
    {{end}}
{{end}}
