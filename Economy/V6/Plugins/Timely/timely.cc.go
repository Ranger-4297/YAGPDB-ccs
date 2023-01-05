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
		{{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0 "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
	{{end}}
	{{with (dbGet $userID "EconomyInfo")}}
		{{$a = sdict .Value}}
		{{$streaks := $a.streaks}}
		{{$streak := $streaks.daily}}
		{{$cash := $a.cash}}
		{{if (reFind `daily` $.Cmd)}}
			{{if not ($cd := dbGet $userID "dCooldown")}}
				{{dbSetExpire $userID "dCooldown" "cooldown" 86400}}
				{{if (dbGet $userID "dGraceCooldown")}}
					{{if not (eq $streaks.daily 9)}}
						{{$streak = add $streak 1}}
						{{$streaks.Set "daily" $streak}}
						{{$a.Set "streaks" $streaks}}
					{{end}}
				{{else}}
					{{$streak = 0}}
					{{$streaks.Set "daily" $streak}}
					{{$a.Set "streaks" $streaks}}
				{{end}}
				{{$streak = (print "1." $streak)}}
				{{$daily = toInt (mult (toFloat $streak) $daily)}}
				{{$embed.Set "description" (print "You've just claimed your " $symbol $daily " daily! Come back in 1 day")}}
				{{$embed.Set "color" $successColor}}
				{{$newBalance := add $cash $daily}}
				{{$a.Set "cash" $newBalance}}
				{{dbSetExpire $userID "dGraceCooldown" "cooldown" 129600}}
				{{dbSet $userID "EconomyInfo" $a}}
			{{else}}
				{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cd.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else if (reFind `weekly` $.Cmd)}}
			{{if not ($cd := dbGet $userID "wCooldown")}}
				{{dbSetExpire $userID "wCooldown" "cooldown" 604800}}
				{{if (dbGet $userID "wGraceCooldown")}}
					{{if not (eq $streaks.weekly 9)}}
						{{$streak = add $streak 1}}
						{{$streaks.Set "weekly" $streak}}
						{{$a.Set "streaks" $streaks}}
					{{end}}
				{{else}}
					{{$streak = 0}}
					{{$streaks.Set "weekly" $streak}}
					{{$a.Set "streaks" $streaks}}
				{{end}}
				{{$streak := $streaks.weekly}}
				{{$streak = (print "1." $streak)}}
				{{$weekly = toInt (mult (toFloat $streak) $weekly)}}
				{{$embed.Set "description" (print "You've just claimed your " $symbol $weekly " weekly! Come back in 1 week")}}
				{{$embed.Set "color" $successColor}}
				{{$newBalance := add $cash $weekly}}
				{{$a.Set "cash" $newBalance}}
				{{dbSetExpire $userID "wGraceCooldown" "cooldown" 691200}}
				{{dbSet $userID "EconomyInfo" $a}}
			{{else}}
				{{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cd.ExpiresAt.Sub currentTime)))}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
		{{else if (reFind `monthly` $.Cmd)}}
			{{if not ($cd := dbGet $userID "mCooldown")}}
				{{dbSetExpire $userID "mCooldown" "cooldown" 2419200}}
				{{if (dbGet $userID "mGraceCooldown")}}
					{{if not (eq $streaks.daily 9)}}
						{{$streak = add $streak 1}}
						{{$streaks.Set "monthly" $streak}}
						{{$a.Set "streaks" $streaks}}
					{{end}}
				{{else}}
					{{$streak = 0}}
					{{$streaks.Set "monthly" $streak}}
					{{$a.Set "streaks" $streaks}}
				{{end}}
				{{$streak := $streaks.monthly}}
				{{$streak = (print "1." $streak)}}
				{{$monthly = toInt (mult (toFloat $streak) $monthly)}}
				{{$embed.Set "description" (print "You've just claimed your " $symbol $monthly " monthly! Come back in 1 month")}}
				{{$embed.Set "color" $successColor}}
				{{$newBalance := add $cash $monthly}}
				{{$a.Set "cash" $newBalance}}
				{{dbSetExpire $userID "mGraceCooldown" "cooldown" 2505600}}
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
{{sendMessage nil (cembed $embed)}}