{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(dep(osit)?|with(draw)?)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1}}

{{/* Deposit, Withdraw */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with .Cmd}}
	{{$cmd := $.Cmd}}
	{{with (dbGet 0 "EconomySettings")}}
		{{$a := sdict .Value}}
		{{$symbol := $a.symbol}}
		{{$bankDB := or (dbGet 0 "bank").Value sdict}}
		{{$bank := or ($bankDB.Get (toString $userID)) 0 | toInt}}
		{{$cash := or (dbGet $userID "cash").Value 0 | toInt}}
		{{if (reFind `deposit|dep` $cmd)}}
			{{with $.CmdArgs}}
				{{$amount := (index $.CmdArgs 0)}}
				{{if (toInt $amount)}}
					{{if gt (toInt $amount) (toInt $cash)}}
						{{$embed.Set "description" (print "You're unable to deposit more than you have on hand.\nYou currently have " $symbol (humanizeThousands $cash) " on you.")}}
						{{$embed.Set "color" $errorColor}}
					{{else}}
						{{$embed.Set "description" (print "You deposited " $symbol (humanizeThousands $amount) " into your bank!")}}
						{{$embed.Set "color" $successColor}}
						{{$cash = sub $cash $amount}}
						{{$bank = add $bank $amount}}
					{{end}}
				{{else if eq (lower (toString $amount)) "all"}}
					{{$embed.Set "description" (print "You deposited " $symbol (humanizeThousands $cash) " into your bank!")}}
					{{$embed.Set "color" $successColor}}
					{{$cash = (toInt 0)}}
					{{$bank = add $bank $cash}}
				{{else}}
					{{$embed.Set "description" (print "Invalid `amount` argument provided.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "No `amount` argument provided.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else if (reFind `withdraw|with` $cmd)}}
			{{with $.CmdArgs}}
				{{$amount := (index $.CmdArgs 0)}}
				{{if (toInt $amount)}}
					{{if gt (toInt $amount) (toInt $bank)}}
						{{$embed.Set "description" (print "You're unable to withdraw more than you have in your bank.\nYou currently have " $symbol (humanizeThousands $bank) " in your bank.")}}
						{{$embed.Set "color" $errorColor}}
					{{else}}
						{{$embed.Set "description" (print "You withdrew " $symbol (humanizeThousands $amount) " from your bank!")}}
						{{$embed.Set "color" $successColor}}
						{{$cash = add $cash $amount}}
						{{$bank = sub $bank $amount}}
					{{end}}
				{{else if eq (lower (toString $amount)) "all"}}
					{{$embed.Set "description" (print "You withdrew " $symbol (humanizeThousands $bank) " from your bank!")}}
					{{$embed.Set "color" $successColor}}
					{{$cash = add $bank $cash}}
					{{$bank = (toInt 0)}}
				{{else}}
					{{$embed.Set "description" (print "Invalid amount.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "No `amount` argument provided.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{end}}
		{{$bankDB.Set (toString $userID) $bank}}
		{{dbSet 0 "bank" $bankDB}}
		{{dbSet $userID "cash" $cash}}
	{{else}}
		{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{end}}
{{sendMessage nil (cembed $embed)}}