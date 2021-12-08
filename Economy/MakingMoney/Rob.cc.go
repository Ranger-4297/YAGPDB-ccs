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
{{if eq $victim .User.ID}}
    {{$errorEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You're can't to rob yourself, silly! Please specify a valid user")
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
            {{sendMessage nil $errorEmbed}}
{{else}}
    {{$a := ""}}
    {{$b := .User.ID}}
    {{$cash := ""}}
    {{$failRate := ""}}
    {{$symbol := ""}}
    {{if not (dbGet $b "EconomyInfo")}}
        {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
    {{end}}
    {{with (dbGet 0 "EconomySettings")}}
        {{$a = sdict .Value}}
        {{$failRate = $a.failRate}}
        {{$symbol = $a.symbol}}
        {{with (dbGet $victim "EconomyInfo")}}
            {{$a = sdict .Value}}
            {{$victimsCash := $a.cash}} {{/* Amount victim has before robbed */}}
            {{if eq (toInt $victimsCash) 0}}
                {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to rob <@!" $victim "> to a value below `0`")
                            "color" 0x00ff8b
                            "timestamp" currentTime
                )}}
                {{sendMessage nil $errorEmbed}}
            {{else}}
                {{$amount := (randInt $victimsCash)}} {{/* Amount stolen from victim */}}
                {{$victimsNewCash := (sub (toInt $victimsCash) $amount)}} {{/* Amount victim has after subtracting stolen money */}}
                {{$sdict := (dbGet $victim "EconomyInfo").Value}}
                {{$sdict.Set "cash" $victimsNewCash}}
                {{dbSet $victim "EconomyInfo" $sdict}}
                {{with (dbGet $b "EconomyInfo")}}
                    {{$a = sdict .Value}}
                    {{$yourCash := $a.cash}} {{/* Amount you have before robbing */}}
                    {{$yourNewCash := (add (toInt $yourCash) $amount)}} {{/* Amount you have after adding stolen money */}}
                    {{$crimeEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))
                            "description" (print "You robbed " $symbol $amount " from <@!" $victim ">")
                            "color" 0x00ff7b
                            "timestamp" currentTime
                    )}}
                    {{sendMessage nil $crimeEmbed}}
                    {{$sdict = (dbGet $b "EconomyInfo").Value}}
                    {{$sdict.Set "cash" $yourNewCash}}
                    {{dbSet $b "EconomyInfo" $sdict}}
                {{end}}
            {{end}}
        {{end}}
    {{end}}
{{end}}