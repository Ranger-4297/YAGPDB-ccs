{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(user-?set(tings)?)(\s+|\z)`

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

{{/* user Set */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$econData := or (dbGet $userID "userEconData").Value (sdict "settings" sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
{{$economySettings := (dbGet 0 "EconomySettings").Value}}
{{if not $economySettings}}
	{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$settings := print "\nAvailable settings: `balance`, `trading`, `inventory`, `leaderboard`\nTo set it with the default settings `" .Cmd " default`"}}
{{if not .CmdArgs}}
	{{$embed.Set "description" (print "No `Setting` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value:Yes/No>`" $settings)}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$setting := index .CmdArgs 0 | lower}}
{{if eq $setting "default"}}
	{{$embed.Set "description" (print "Set your account to default values")}}
	{{$embed.Set "color" $successColor}}
	{{$econData.Set "settings" (sdict "balance" "yes" "trading" "yes" "inventory" "yes" "leaderboard" "yes")}}
{{else}}
	{{if not (eq $setting "inventory" "leaderboard" "trading" "balance" "view")}}
		{{$embed.Set "description" (print "Invalid `Setting` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value:Yes/No>`" $settings)}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{if not (eq $setting "view")}}
		{{if lt (len .CmdArgs) 2}}
			{{$embed.Set "description" (print "No `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value:Yes/No>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$value := index .CmdArgs 1 | lower}}
		{{if not (eq $value "yes" "no")}}
			{{$embed.Set "description" (print "Invalid `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value:Yes/No>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$econData.settings.Set $setting $value}}
		{{$embed.Set "description" (print "`" $setting "` set to `" $value "`")}}
		{{$embed.Set "color" $successColor}}
	{{else}}
		{{$lb := or $econData.settings.leaderboard "yes"}}
		{{$cash := or $econData.settings.balance "yes"}}
		{{$trading := or $econData.settings.trading "yes"}}
		{{$inventory := or $econData.settings.inventory "yes"}}
		{{$embed.Set "description" (print "Your user settings can be found below:\n\n**Leaderboard:** `" $lb "`\n**Balance:** `" $cash "`\n**Trading:** `" $trading "`\n**Inventory:** `" $inventory "`")}}
		{{$embed.Set "color" $successColor}}
	{{end}}
{{end}}
{{dbSet $userID "userEconData" $econData}}
{{sendMessage nil (cembed $embed)}}