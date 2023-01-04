{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(daily|weekly|monthly)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$daily := 2500}}
{{$weekly := 5000}}
{{$monthly := 10000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1}}


{{/* Timely */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
	{{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{if not (dbGet $userID "EconomyInfo")}}
		{{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
	{{end}}
	{{with (dbGet $userID "EconomyInfo")}}
		{{$a = sdict .Value}}
		{{$balance = $a.cash}}
		{{$cmd := $.Cmd | lower | toString}}
		{{if eq $cmd "daily"}}
			{{if not ($cd := dbGet $userID "dCooldown")}}
				{{dbSetExpire $userID "dCooldown" "cooldown" 86400}}
				{{$embed.Set "description" (print "You've just claimed your " $symbol $daily " " $cmd"! Come back in " (humanizeDurationSeconds ($cd.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $successColor}}
				{{$newBalance := add $cash $daily}}
				{{$a.Set "cash" $newBalance}}
				{{dbSet $userID "EconomyInfo" $a}}
			{{else}}
				{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cd.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else if eq $cmd "weekly"}}
			{{if not ($cd := dbGet $userID "wCooldown")}}
				{{dbSetExpire $userID "wCooldown" "cooldown" 691200}}
				{{$embed.Set "description" (print "You've just claimed your " $symbol $weekly " " $cmd"! Come back in " (humanizeDurationSeconds ($cd.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $successColor}}
				{{$newBalance := add $cash $weekly}}
				{{$a.Set "cash" $newBalance}}
				{{dbSet $userID "EconomyInfo" $a}}
			{{else}}
				{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cd.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else if eq $cmd "monthly"}}
			{{if not ($cd := dbGet $userID "mCooldown")}}
				{{dbSetExpire $userID "mCooldown" "cooldown" 2716400}}
				{{$embed.Set "description" (print "You've just claimed your " $symbol $monthly " " $cmd"! Come back in " (humanizeDurationSeconds ($cd.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $successColor}}
				{{$newBalance := add $cash $monthly}}
				{{$a.Set "cash" $newBalance}}
				{{dbSet $userID "EconomyInfo" $a}}
			{{else}}
				{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cd.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{end}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}