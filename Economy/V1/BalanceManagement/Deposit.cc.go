{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Deposit`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "deposit <Amount:Number/All>" (carg "string" "Amount")}}
{{$amount := ($args.Get 0)}}
{{$a := ""}}
{{$cash := ""}}
{{$bank := ""}}
{{$symbol := ""}}
{{$b := .User.ID}}
{{if not (dbGet $b "EconomyInfo")}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a = sdict .Value}}
    {{$symbol = $a.symbol}}
    {{with (dbGet $b "EconomyInfo")}}
        {{$a = sdict .Value}}
        {{$cash = $a.cash}}
        {{$bank = $a.bank}}
        {{$newBank := ""}}
        {{if eq $amount "all"}}
            {{$newBank = (add (toInt $cash) (toInt $bank))}}
            {{$depositEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You deposited " $symbol $cash " into your bank!")
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
            {{sendMessage nil $depositEmbed}}
            {{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
            {{$sdict.Set "bank" $newBank}}
            {{dbSet $b "EconomyInfo" $sdict}}
            {{$sdict.Set "cash" (toInt "0")}}
            {{dbSet $b "EconomyInfo" $sdict}}
        {{else if (toInt $amount)}}
            {{if gt (toInt $amount) $cash}}
                {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to deposit more than you have on hand")
                            "color" 0x00ff8b
                            "timestamp" currentTime
                )}}
                {{sendMessage nil $errorEmbed}}
            {{else}}
                {{$moneyToDeposit := (toInt $amount)}}
                {{$newCash := (sub $cash $amount)}}
                {{$newBank = (add $bank $amount)}}
                {{$depositEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You deposited " $symbol $moneyToDeposit " into your bank!")
                            "color" 0x00ff7b
                            "timestamp" currentTime
                            )}}
                {{sendMessage nil $depositEmbed}}
                {{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
                {{$sdict.Set "bank" $newBank}}
                {{dbSet $b "EconomyInfo" $sdict}}
                {{$sdict.Set "cash" $newCash}}
                {{dbSet $b "EconomyInfo" $sdict}}
            {{end}}
        {{else}}
            {{$errorEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You're unable to deposit this value, check that you used a valid cash amount or all")
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
            {{sendMessage nil $errorEmbed}}
        {{end}}
    {{end}}
{{end}}