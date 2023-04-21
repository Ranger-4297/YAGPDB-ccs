{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(rollnum(ber)?|rn|snake?-?eyes)(\s+|\z)`

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

{{/* Roll, SnakeEyes */}}

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
			{{if le $bet $betMax}}
				{{if (reFind `rollnum(ber)?|rn` $.Cmd)}}
					{{if not ($cooldown := dbGet $userID "rollCooldown")}}
						{{dbSetExpire $userID "rollCooldown" "cooldown" $incomeCooldown}}
						{{$roll := (randInt 1 101)}}
						{{$rs := 1}}
						{{$amount := ""}}
						{{if and (ge $roll 65) (lt $roll 90)}}
							{{$amount = $bet}}
							{{$bal = add $bal $amount}}
						{{else if and (ge $roll 90) (lt $roll 100)}}
							{{$amount = (mult $bet 3)}}
							{{$bal = add $bal $amount}}
						{{else if eq $roll 100}}
							{{$amount = (mult $bet 5)}}
							{{$bal = add $bal $amount}}
						{{else}}
							{{$bal = sub $bal $bet}}
							{{$rs = 0}}
						{{end}}
						{{if $rs}}
							{{$embed.Set "description" (print "The roll landed on " $roll " and you and won " $symbol (humanizeThousands $amount) "!")}}
							{{$embed.Set "color" $successColor}}
						{{else}}
							{{$embed.Set "description" (print "The roll landed on " $roll " and you and lost " $symbol (humanizeThousands $bet))}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else if (reFind `snake?-?eyes` $.Cmd)}}
					{{if not ($cooldown := dbGet $userID "snakeeyesCooldown")}}
						{{dbSetExpire $userID "snakeeyesCooldown" "cooldown" $incomeCooldown}}
						{{$die1 := (randInt 1 7)}}
						{{$die2 := (randInt 1 7)}}
						{{if and (eq $die1 1) (eq $die2 1)}}
							{{$embed.Set "description" (print "You rolled snake eyes (" $die1 "&" $die2 ")\nAnd won " $symbol (humanizeThousands (mult $bet 36)))}}
							{{$embed.Set "color" $successColor}}
							{{$bal = add $bal (mult $bet 36)}}
						{{else}}
							{{$embed.Set "description" (print "You rolled " $die1 "&" $die2 " and lost " $symbol (humanizeThousands $bet) ".")}}
							{{$embed.Set "color" $errorColor}}
							{{$bal = sub $bal $bet}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "No `bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
	{{dbSet $userID "cash" $bal}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}