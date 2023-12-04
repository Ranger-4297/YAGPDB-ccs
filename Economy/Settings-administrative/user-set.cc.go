{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(user-?set(tings)?)(\s+|\z)`

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

{{/* user Set */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{$econData := or (dbGet $userID "userEconData").Value (sdict "settings" sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
{{with (dbGet 0 "EconomySettings")}}
	{{with $.CmdArgs}}
		{{if index $.CmdArgs 0}}
			{{$setting := (index $.CmdArgs 0) | lower}}
			{{$settings := or $econData.settings (sdict "balance" "yes" "trading" "yes" "inventory" "yes" "leaderboard" "yes")}}
			{{if eq $setting "default"}}
				{{$embed.Set "description" (print "Set your account to default values")}}
				{{$embed.Set "color" $successColor}}
				{{$econData.Set "settings" $settings}}
			{{else if eq $setting "inventory" "leaderboard" "trading" "balance"}}
				{{if gt (len $.CmdArgs) 1}}
					{{$value := (index $.CmdArgs 1)}}
					{{if eq $value "yes" "no"}}
						{{$settings.Set $setting $value}}
						{{$econData.Set "settings" $settings}}
						{{$embed.Set "description" (print "`" $setting "` set to `" $value "`")}}
						{{$embed.Set "color" $successColor}}
					{{else}}
						{{$embed.Set "description" (print "Invalid value argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Yes/No>`")}}
						{{$embed.Set "color" $errorColor}}
					{{end}}
				{{else}}
					{{$embed.Set "description" (print "No value argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Yes/No>`")}}
					{{$embed.Set "color" $errorColor}}
				{{end}}
			{{else if eq $setting "view"}}
				{{$lb := $settings.leaderboard}}
				{{$trading := $settings.trading}}
				{{$inventory := $settings.inventory}}
				{{$embed.Set "description" (print "Your user settings can be found below:\n\n**Leaderboard:** `" $lb "`\n**Trading:** `" $trading "`\n**Inventory:** `" $inventory "`")}}
				{{$embed.Set "color" $successColor}}
			{{else}}
				{{$embed.Set "description" (print "Invalid setting argument passed.\nSyntax is: `" $.Cmd " <Setting> <Value:Yes/No>`")}}
				{{$embed.Set "color" $errorColor}}
			{{end}}
			{{dbSet $userID "userEconData" $econData}}
		{{end}}
	{{else}}
		{{$embed.Set "description" (print "No setting argument passed.\nSyntax is: `" $.Cmd " <Setting> <Value:Yes/No>`")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{else}}
	{{$embed.Set "description" (print "No economy database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
	{{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}