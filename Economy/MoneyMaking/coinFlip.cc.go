{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(coin-?flip|cf|flip-?coin|fc)(\s+|\z)`

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

{{/* Flip */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{$betMax := $a.betMax | toInt}}
	{{$incomeCooldown := $a.incomeCooldown | toInt}}
	{{$bal := or (dbGet $userID "cash").Value 0 | toInt}}
	{{with $.CmdArgs}}
		{{if (index . 0)}}
			{{$side := (index . 0) | toString | lower}}
			{{$picker1 := ""}}
			{{$win := ""}}
			{{$lose := ""}}
			{{if (reFind `(t(ails?)?|h(eads?)?)` (lower $side))}}
				{{if eq $side "t" "tails" "tail"}}
					{{$side = "tails"}}
				{{else if eq $side "h" "heads" "head"}}
					{{$side = "heads"}}
				{{end}}
				{{if gt (len .) 1}}
					{{$bet := (index . 1)}}
					{{if $bet | toInt}}
						{{$bet = $bet | toInt}}
						{{if gt $bet 0}}
							{{if le $bet $betMax}}
								{{if le $bet $bal}}
									{{if not ($cooldown := dbGet $userID "coinflipCooldown")}}
										{{dbSetExpire $userID "coinflipCooldown" "cooldown" $incomeCooldown}}
										{{$int := randInt 1 3}}
										{{if eq $int 1}} {{/* Win */}}
											{{$bal = add $bal $bet}}
											{{$embed.Set "description" (print "You flipped " $side " and won " $symbol (humanizeThousands $bet))}}
											{{$embed.Set "color" $successColor}}
										{{else}} {{/* Lose */}}
											{{$bal = sub $bal $bet}}
											{{$embed.Set "description" (print "You flipped " $side " and lost.")}}
											{{$embed.Set "color" $errorColor}}
										{{end}}
									{{else}}
										{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
										{{$embed.Set "color" $errorColor}}
									{{end}}
								{{else}}
									{{$embed.Set "description" (print "You can't bet more than you have!")}}
									{{$embed.Set "color" $errorColor}}
								{{end}}
							{{else}}
								{{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "No `Bet` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "Invalid `Side` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "No `Side` argument provided.\nSyntax is `" $.Cmd " <Side:Head/Tails> <Bet:Amount>`")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
	{{dbSet $userID "cash" $bal}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}