{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(work|job|get-?paid|(commit-?)?crime|rob|steal)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := (index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 100)}}

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
	{{$workCooldown := $a.workCooldown | toInt}}
	{{$robCooldown := $a.robCooldown | toInt}}
	{{$crimeCooldown := $a.crimeCooldown | toInt}}
	{{if not (dbGet $userID "EconomyInfo")}}
		{{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
	{{end}}
	{{with (dbGet $userID "EconomyInfo")}}
		{{$a = sdict .Value}}
		{{$cash := $a.cash}}
		{{$cmd := $.Cmd | toString | lower}}
		{{if (reFind `(work|job|get-?paid|labor)` $cmd)}}
			{{if not ($cooldown := dbGet $userID "workCooldown")}}
				{{dbSetExpire $userID "workCooldown" "cooldown" $workCooldown}}
				{{$workPay := (mult (randInt $min $max) (randInt 1 3))}}
				{{$newCashBalance := $cash | add $workPay}}
				{{$embed.Set "description" (print "You decided to work today! You got paid a hefty " $symbol (humanizeThousands $workPay))}}
				{{$embed.Set "color" 0x00ff7b}}
				{{$a.Set "cash" $newCashBalance}}
				{{dbSet $userID "EconomyInfo" $a}}
			{{else}}
				{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else if (reFind `(commit-?)?crime` $cmd)}}
			{{if not ($cooldown := dbGet $userID "crimeCooldown")}}
				{{dbSetExpire $userID "crimeCooldown" "cooldown" $crimeCooldown}}
				{{$amount := (mult (randInt $min $max) (randInt 1 5))}}
				{{$newCash := ""}}
				{{$int := randInt 1 3}}
				{{if eq $int 1}}
					{{$newCash = $cash | add $amount}}
					{{$embed.Set "description" (print "You broke the law for a pretty penny! You made " $symbol (humanizeThousands $amount) " in your crime spree today")}}
					{{$embed.Set "color" $successColor}}
				{{else}}
					{{$newCash = $amount | sub $cash}}
					{{$embed.Set "description" (print "You broke the law trying to commit a felony! You were arrested and lost " $symbol (humanizeThousands $amount) " due to your bail.")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
				{{$a.Set "cash" $newCash}}
				{{dbSet $userID "EconomyInfo" $a}}
			{{else}}
				{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else if (reFind `(rob|steal)` $cmd)}}
			{{with $.CmdArgs}}
				{{if (index . 0)}}
					{{if (index . 0) | getMember}}
						{{$user (index . 0) | getMember}}
						{{$victim := $user.User.ID}}
						{{if not (eq $victim $userID)}}
							{{if not ($cooldown := dbGet $userID "robCooldown")}}
								{{dbSetExpire $userID "robCooldown" "cooldown" $robCooldown}}
								{{with (dbGet $victim "EconomyInfo")}}
									{{$b := sdict .Value}}
									{{$victimsCash := $b.cash | toInt}}
									{{if not (eq $victimsCash 0)}}
										{{$amount := (randInt $victimsCash)}} {{/* Amount stolen from victim */}}
										{{$victimsNewCash := (sub $victimsCash $amount)}} {{/* Amout victim will have after being robbed */}}
										{{$yourNewCash := (add $cash $amount)}} {{/* Amount you will have after robbing vitim */}}
										{{$embed.Set "description" (print "You robbed " $symbol (humanizeThousands $amount) " from <@!" $victim ">")}}
										{{$embed.Set "color" $successColor}}
										{{$b.Set "cash" $victimsNewCash}}
										{{dbSet $victim "EconomyInfo" $b}}
										{{$a.Set "cash" $yourNewCash}}
										{{dbSet $userID "EconomyInfo" $a}}
									{{else}}
										{{$embed.Set "description" (print "<@!" $victim "> has no money left for you to rob!")}}
										{{$embed.Set "color" 0x00ff8b}}
									{{end}}
								{{else}}
									{{dbSet $victim "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
									{{$embed.Set "description" (print "<@!" $victim "> doesn't have any money for you to rob!")}}
									{{$embed.Set "color" $errorColor}}
								{{end}}
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
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}