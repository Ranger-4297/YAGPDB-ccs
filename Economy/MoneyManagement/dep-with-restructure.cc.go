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
{{$user := .User}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}

{{/* Deposit, Withdraw */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024")) "timestamp" currentTime}}
{{with .Cmd}}
	{{$cmd := $.Cmd}}
	{{with (dbGet 0 "EconomySettings")}}
		{{$a := sdict .Value}}
		{{$symbol := $a.symbol}}
		{{$dbecoInfo := dbGet $user.ID "EconomyInfo"}}
		{{if not $dbecoInfo}}
			{{$dbecoInfo = sdict "cash" 200 "bank" 0}}
			{{dbSet $user.ID "EconomyInfo" $dbecoInfo}}
		{{end}}
		{{with $.CmdArgs}}
			{{with $dbecoInfo}}	
				{{$a = sdict .Value}}
				{{$cash := toInt $a.cash}}
				{{$bank := toInt $a.bank}}
				{{$amount := index $.CmdArgs 0}}
				{{if (reFind `deposit|dep` $cmd)}}
					{{if toInt $amount}}
						{{if gt (toInt $amount) (toInt $cash)}}
							{{$embed.Set "description" (print "You're unable to deposit more than you have on hand.\nYou currently have " $symbol (humanizeThousands $cash) " on you.")}}
							{{$embed.Set "color" $errorColor}}
						{{else}}
							{{$newCashBalance := $amount | sub $cash}}
							{{$newBankBalance := $amount | add $bank}}
							{{$embed.Set "description" (print "You deposited " $symbol (humanizeThousands $amount) " into your bank!")}}
							{{$embed.Set "color" $successColor}}
							{{dbSet $user.ID "EconomyInfo" (sdict "cash" $newCashBalance "bank" $newBankBalance)}}
						{{end}}
					{{else if eq (lower (toString $amount)) "all"}}
						{{$newCashBalance := (toInt 0)}}
						{{$newBankBalance := $bank | add $cash}}
						{{$embed.Set "description" (print "You deposited " $symbol (humanizeThousands $cash) " into your bank!")}}
						{{$embed.Set "color" $successColor}}
						{{dbSet $user.ID "EconomyInfo" (sdict "cash" $newCashBalance "bank" $newBankBalance)}}
					{{else}}
						{{$embed.Set "description" (print "Invalid `amount` argument provided.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else if (reFind `withdraw|with` $cmd)}}
					{{if toInt $amount}}
						{{if gt (toInt $amount) (toInt $bank)}}
							{{$embed.Set "description" (print "You're unable to withdraw more than you have in your bank.\nYou currently have " $symbol (humanizeThousands $bank) " in your bank.")}}
							{{$embed.Set "color" $errorColor}}
						{{else}}
							{{$newCashBalance := $amount | add $cash}}
							{{$newBankBalance := $amount | sub $bank}}
							{{$embed.Set "description" (print "You withdrew " $symbol (humanizeThousands $amount) " from your bank!")}}
							{{$embed.Set "color" $successColor}}
							{{dbSet $user.ID "EconomyInfo" (sdict "cash" $newCashBalance "bank" $newBankBalance)}}
						{{end}}
					{{else if eq (lower (toString $amount)) "all"}}
						{{$newCashBalance := $bank | add $cash}}
						{{$newBankBalance := (toInt 0)}}
						{{$embed.Set "description" (print "You withdrew " $symbol (humanizeThousands $bank) " from your bank!")}}
						{{$embed.Set "color" $successColor}}
						{{dbSet $user.ID "EconomyInfo" (sdict "cash" $newCashBalance "bank" $newBankBalance)}}
					{{else}}
						{{$embed.Set "description" (print "Invalid amount.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{end}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "No `amount` argument provided.\nCommand syntax is `" $.Cmd " <Amount:Amount>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{end}}
{{sendMessage nil (cembed $embed)}}