{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)((cock|chicken)-?fight))(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}

{{/* Cock fight */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{$betMax := $a.betMax | toInt}}
	{{$incomeCooldown := $a.incomeCooldown | toInt}}
	{{$bal := or (dbGet $userID "cash").Value 0 | toInt}}
	{{$econData := or (dbGet $userID "userEconData").Value (sdict "settings" sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0) "cfWinChance" 50)}}
	{{$cfWC := or ($econData.Get "cfWinChance") 50}}
	{{with $.CmdArgs}}
		{{$bet := (index . 0)}}
		{{$continue := false}}
		{{if eq ($bet | toString) "all"}}
			{{$bet = $bal}}
		{{end}}
		{{if $bet = (toInt $bet)}}
			{{if gt $bet 0}}
				{{if le $bet $bal}}
					{{$continue = true}}
				{{else}}
					{{$embed.Set "description" (print "You can't bet more than you have!")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
		{{if $continue}}
			{{if $betMax}}
				{{if le $bet $betMax}}
					{{$continue = true}}
				{{else}}
					{{$continue = false}}
					{{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{end}}
			{{if $continue}}
				{{$inventory := $econData.inventory}}
				{{if $inventory.Get "chicken"}}
					{{if not ($cooldown := dbGet $userID "cockFightCooldown")}}
						{{dbSetExpire $userID "cockFightCooldown" "cooldown" $incomeCooldown}}
						{{$chance := randInt 1 101}}
						{{if le $chance $cfWC}}
							{{if not (eq $cfWC 70)}}
								{{$cfWC = $cfWC | add 1}}
							{{end}}
							{{$bal = add $bal $bet}}
							{{$embed.Set "description" (print "Your chicken won the fight!\nPlay again with `" $.Cmd " <Bet:Amount>`")}}
							{{$embed.Set "color" $successColor}}
						{{else}}
							{{$cfWC = 50}}
							{{$embed.Set "description" (print "Your chicken lost the fight and died :(\nBuy a new one to play again!")}}
							{{$embed.Set "color" $errorColor}}
							{{$bal = sub $bal $bet}}
							{{$inventory.Del "chicken"}}
							{{$econData.Set "inventory" $inventory}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "You can't bet without a chicken! Buy one from the shop with `" $prefix "buy-item chicken`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "No `bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
	{{dbSet $userID "cash" $bal}}
	{{$econData.Set "cfWinChance" $cfWC}}
	{{dbSet $userID "userEconData" $econData}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}