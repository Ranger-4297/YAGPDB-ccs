{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(give-?money|loan-?money|pay)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}

{{/* Give money */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{$cash := or (dbGet $userID "cash").Value 0 | toInt}}
	{{with $.CmdArgs}}
		{{if index . 0}}
			{{if index . 0 | getMember}}
				{{$user := getMember (index . 0)}}
				{{$receivingUser := $user.User.ID}}
				{{if eq $receivingUser $.User.ID}}
					{{$embed.Set "description" (print "You cannot give money to yourself.")}}
					{{$embed.Set "color" $errorColor}}
				{{else}}
					{{if gt (len $.CmdArgs) 1}}
						{{$amount := (index $.CmdArgs 1)}}
						{{if (toInt $amount)}}
							{{if gt (toInt $amount) 0}}
								{{$rCash := or (dbGet $receivingUser "cash").Value 0 | toInt}}
								{{if gt (toInt $amount) (toInt $cash)}}
									{{$embed.Set "description" (print "You cannot give more than you have.")}}
									{{$embed.Set "color" $errorColor}}
								{{else}}
									{{$rCash = add $rCash $amount}}
									{{$cash = sub $cash $amount}}
									{{$embed.Set "description" (print "You gave " $symbol (humanizeThousands $amount) " to <@!" $receivingUser ">\nThey now have " $symbol (humanizeThousands $rCash) " in cash!")}}
									{{$embed.Set "color" $successColor}}
								{{end}}
								{{dbSet $receivingUser "cash" $rCash}}
							{{else}}
								{{$embed.Set "description" (print "You're unable to give this value, check that you used a valid number above 1")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
						{{else}}
							{{$embed.Set "description" (print "Invalid `Amount` argument passed.\nCheck that you used a valid number above 1")}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "No `Amount` argument passed.\nSyntax is: `" $.Cmd " <Member:Mention/ID> <Amount:Amount>`")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{end}}
			{{else}}
				{{$embed.Set "description" (print "Invalid `user` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Amount:Amount>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "No `User` argument provided.\nSyntax is `" $.Cmd " <User:Mention/ID> <Amount:Amount>`")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
	{{dbSet $userID "cash" $cash}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}