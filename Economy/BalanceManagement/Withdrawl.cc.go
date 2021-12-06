{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Withdraw`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 1 "deposit <Amount:Number/All>" (carg "string" "Amount")}}
{{$amount := ($args.Get 0)}}
{{$a := ""}}
{{$cash := ""}}
{{$bank := ""}}
{{$EconomySymbol := ""}}
{{$b := .User.ID}}
{{if not (dbGet $b "EconomyInfo")}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a = sdict .Value}}
	{{$EconomySymbol = $a.EconomySymbol}}
	{{with (dbGet $b "EconomyInfo")}}
		{{$a = sdict .Value}}
		{{$cash = $a.cash}}
		{{$bank = $a.bank}}
		{{$newCash := ""}}
		{{if eq $amount "all"}}
			{{$newCash = (add (toInt $bank) (toInt $cash))}}
			{{$withdrawEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You withdrew " $EconomySymbol $bank " from your bank!")
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
			{{sendMessage nil $withdrawEmbed}}
			{{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
			{{$sdict.Set "bank" (toInt "0")}}
			{{dbSet $b "EconomyInfo" $sdict}}
			{{$sdict.Set "cash" $newCash}}
			{{dbSet $b "EconomyInfo" $sdict}}
		{{else if (toInt $amount)}}
			{{if gt (toInt $amount) $bank}}
				{{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to withdraw more than you have in your bank")
                            "color" 0x00ff8b
                            "timestamp" currentTime
                )}}
                {{sendMessage nil $errorEmbed}}
			{{else}}
				{{$moneyToWithdraw := (toInt $amount)}}
				{{$newCash = (add $cash $amount)}}
				{{$newBank := (sub $bank $amount)}}
				{{$withdrawEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You withdrew " $EconomySymbol $moneyToWithdraw " from your bank!")
                            "color" 0x00ff7b
                            "timestamp" currentTime
                )}}
                {{sendMessage nil $withdrawEmbed}}
                {{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
                {{$sdict.Set "bank" $newBank}}
                {{dbSet $b "EconomyInfo" $sdict}}
                {{$sdict.Set "cash" $newCash}}
                {{dbSet $b "EconomyInfo" $sdict}}
			{{end}}
		{{else}}
			{{$errorEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
            "description" (print "You're unable to withdraw this value, check that you used a valid cash amount or all")
            "color" 0x00ff8b
            "timestamp" currentTime
            )}}
            {{sendMessage nil $errorEmbed}}
		{{end}}
	{{end}}
{{end}}