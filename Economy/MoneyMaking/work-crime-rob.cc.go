{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(work|job|get-?paid|(commit-?)?crime|rob|steal)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix}}

{{/* Work, crime, rob */}}

{{/* Resoponse */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
	{{$a := sdict .Value}}
	{{$min := $a.min}}
	{{$max := $a.max}}
	{{$symbol := $a.symbol}}
	{{$responses := $a.responses}}
	{{$enabledResponses := $a.Get "enable-responses"}}
	{{$workCooldown := $a.workCooldown | toInt}}
	{{$robCooldown := $a.robCooldown | toInt}}
	{{$crimeCooldown := $a.crimeCooldown | toInt}}
	{{$cash := or (dbGet $userID "cash").Value 0 | toInt}}
	{{$cmd := $.Cmd | toString | lower}}
	{{if (reFind `(work|job|get-?paid|labor)` $cmd)}}
		{{if not ($cooldown := dbGet $userID "workCooldown")}}
			{{dbSetExpire $userID "workCooldown" "cooldown" $workCooldown}}
			{{$workPay := randInt $min $max}}
			{{$response := (print "You decided to work today! You got paid a hefty " $symbol (humanizeThousands $workPay))}}
			{{if and $enabledResponses $responses.work}}
				{{$response = (reReplace `\(amount\)` (index (shuffle $responses.work) 0) (print $symbol (humanizeThousands $workPay)))}}
			{{end}}
			{{$cash =  add $cash $workPay}}
			{{$embed.Set "description" $response}}
			{{$embed.Set "color" 0x00ff7b}}
		{{else}}
			{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{else if (reFind `(commit-?)?crime` $cmd)}}
		{{if not ($cooldown := dbGet $userID "crimeCooldown")}}
			{{dbSetExpire $userID "crimeCooldown" "cooldown" $crimeCooldown}}
			{{$amount := (mult (randInt $min $max) (randInt 1 5))}}
			{{$int := randInt 1 3}}
			{{if eq $int 1}}
				{{$cash = add $cash $amount}}
				{{$response := (print "You broke the law for a pretty penny! You made " $symbol (humanizeThousands $amount) " in your crime spree today")}}
				{{if and $enabledResponses $responses.crime}}
					{{$response = (reReplace `\(amount\)` (index (shuffle $responses.crime) 0) (print $symbol (humanizeThousands $amount)))}}
				{{end}}
				{{$embed.Set "description" $response}}
				{{$embed.Set "color" $successColor}}
			{{else}}
				{{$cash = sub $cash $amount}}
				{{$embed.Set "description" (print "You broke the lawm and got caught! You were arrested and lost " $symbol (humanizeThousands $amount) " due to your bail.")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{else if (reFind `(rob|steal)` $cmd)}}
		{{with $.CmdArgs}}
			{{if (index . 0)}}
				{{if (index . 0) | getMember}}
					{{$user := (index . 0) | getMember}}
					{{$victim := $user.User.ID}}
					{{if not (eq $victim $userID)}}
						{{if not ($cooldown := dbGet $userID "robCooldown")}}
							{{dbSetExpire $userID "robCooldown" "cooldown" $robCooldown}}
							{{$vicCash := or (dbGet $victim "cash").Value 0 | toInt}}
							{{if $vicCash}}
								{{$amount := (randInt $vicCash)}} {{/* Amount stolen from victim */}}
								{{$vicCash = sub $vicCash $amount}} {{/* Amout victim will have after being robbed */}}
								{{$cash = add $cash $amount}} {{/* Amount you will have after robbing vitim */}}
								{{$embed.Set "description" (print "You robbed " $symbol (humanizeThousands $amount) " from <@!" $victim ">")}}
								{{$embed.Set "color" $successColor}}
							{{else}}
								{{$embed.Set "description" (print "<@!" $victim "> doesn't have any money for you to rob!")}}
								{{$embed.Set "color" $errorColor}}
							{{end}}
							{{dbSet $victim "cash" $vicCash}}
						{{else}}
							{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
							{{$embed.Set "color" $errorColor}}
						{{end}}
					{{else}}
						{{$embed.Set "description" (print "You can't rob yourself.")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "Invalid `User` argument provided.\nSyntaxt is `" $.Cmd " <User:Mention/ID>`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "No `User` argument provided.\nSyntaxt is `" $.Cmd " <User:Mention/ID>`")}}
			{{$embed.Set "color" $errorColor}}
		{{end}}
	{{end}}
	{{dbSet $userID "cash" $cash}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}