{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Balance`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 0 "bal [Member:Mention/ID]" (carg "member" "Member")}}
{{$errorColor := 0xFF0000}}
{{$member := ""}}
{{if ($args.Get 0).User}}
    {{$member = ($args.Get 0).User}}
{{else if not ($args.Get 0)}}
    {{$member = .User}}
{{else}}
    {{$errorEmbed := (cembed
            "description" (print "Invalid user argument passed`")
            "color" $errorColor
            )}}
    {{sendMessage nil $errorEmbed}}
{{end}}
{{if not (dbGet $member.ID "EconomyInfo")}}
    {{dbSet $member.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{with (dbGet $member.ID "EconomyInfo")}}
        {{$a = sdict .Value}}
		{{$cash := $a.cash}}
		{{$bank := $a.bank}}
		{{$balanceEmbed := (cembed
            "author" (sdict "name" $member.Username "icon_url" ($member.AvatarURL "128"))
            "description" (print $member.Mention "'s balance")
            "fields" (cslice 
                (sdict "name" "Cash" "value" (print $symbol (toInt $cash)) "inline" true)
                (sdict "name" "Bank" "value" (print $symbol (toInt $bank)) "inline" true)
                (sdict "name" "Networth" "value" (print $symbol (toString (add (toInt $cash) (toInt $bank)))) "inline" true))
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
		{{sendMessage nil $balanceEmbed}}
	{{end}}
{{end}}