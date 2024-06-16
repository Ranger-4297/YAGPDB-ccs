{{/*
		Made by ranger_4297 (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)server-?(set|config(ure)?)(\s+|\z)`

	©️ RhykerWells 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}

{{/* server set */}}

{{/* Response */}}
{{$embed := sdict "author" (sdict "name" .User.Username "icon_url" (.User.AvatarURL "1024")) "timestamp" currentTime "color" $errorColor}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if not (or (in $perms "Administrator") (in $perms "ManageServer"))}}
	{{$embed.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$settings := print "\nAvailable settings: `max`, `min`, `betMax`, `startbalance`, `symbol`, `workCD`, `incomeCD`, `crimeCD`, `robCD`, `responses`\nTo set it with the default settings `" .Cmd " default`"}}
{{if not .CmdArgs}}
	{{$embed.Set "description" (print "No `Setting` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
	{{sendMessage nil (cembed $embed)}}
	{{return}}
{{end}}
{{$setting := index .CmdArgs 0 | lower}}
{{if eq $setting "default"}}
	{{$embed.Set "description" (print "Set the `EconomySettings` to default values")}}
	{{$embed.Set "color" $successColor}}
	{{dbSet 0 "EconomySettings" (sdict "min" 200 "max" 500 "betMax" 5000 "symbol" "£" "startBalance" 200 "incomeCooldown" 300 "workCooldown" 7200 "crimeCooldown" 14400 "robCooldown" 21600 "enable-responses" false "responses" (sdict "work" cslice "crime" cslice))}}
	{{dbSet 0 "store" (sdict "items" sdict)}}
	{{dbSet 0 "russianRoulette" sdict}}
	{{dbSet 0 "bank" sdict}}
	{{dbSet 0 "roulette" (sdict "game" sdict "storage" sdict)}}
	{{dbSet 0 "accounts" sdict}}
{{else}}
	{{if not (reFind `(work|crime|rob|income)cd|m(in|ax)|s(ymbol|tartBalance)|responses|betMax` $setting)}}
		{{$embed.Set "description" (print "Invalid `Setting` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{$economySettings := (dbGet 0 "EconomySettings").Value}}
	{{if not $economySettings}}
		{{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" .ServerPrefix "server-set default`")}}
		{{sendMessage nil (cembed $embed)}}
		{{return}}
	{{end}}
	{{$symbol := $economySettings.symbol}}
	{{if eq $setting "min" "max" "betmax" "startbalance"}}
		{{$max := $economySettings.max}}
		{{$min := $economySettings.min}}
		{{if lt (len .CmdArgs) 2}}
			{{$embed.Set "description" (print "No `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$val := index .CmdArgs 1}}
		{{if and (not ($val := toInt $val)) (lt $val 1) }}
			{{$embed.Set "description" (print "Invalid `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{if and (eq $setting "max") (lt (toInt $val) (toInt $min))}}
			{{$embed.Set "description" (print "You cannot set `" $setting "` to a value below `min`\n`min` is set to `" (humanizeThousands $min) "`")}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{else if and (eq $setting "min") (gt (toInt $val) (toInt $max))}}
			{{$embed.Set "description" (print "You cannot set `" $setting "` to a value above `max`\n`max` is set to `" (humanizeThousands $max) "`")}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$embed.Set "description" (print "You set `" $setting "` to " $symbol (humanizeThousands $val))}}
		{{$embed.Set "color" $successColor}}
		{{if eq $setting "betmax"}}{{$setting = "betMax"}}{{else if eq $setting "startbalance"}}{{$setting = "startBalance"}}{{end}}
		{{$economySettings.Set $setting $val}}
		{{dbSet 0 "EconomySettings" $economySettings}}
	{{else if eq $setting "symbol"}}
		{{if lt (len .CmdArgs) 2}}
			{{$embed.Set "description" (print "No `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$symbol = index .CmdArgs 1}}
		{{$embed.Set "description" (print "You set the server currency symbol to " $symbol)}}
		{{$embed.Set "color" $successColor}}
		{{$economySettings.Set "symbol" $symbol}}
		{{dbSet 0 "EconomySettings" $economySettings}}
	{{else if eq $setting "responses"}}
		{{if lt (len .CmdArgs) 2}}
			{{$embed.Set "description" (print "No `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$value := index .CmdArgs 1 | lower}}
		{{if not (eq $value "yes" "enable" "enabled" "no" "disable" "disabled")}}
			{{$embed.Set "description" (print "Invalid `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$status := "enabled"}}{{$value = true}}
		{{if not (eq $value "yes" "enable" "enabled")}}
			{{$status = "disabled"}}
			{{$value = false}}
		{{end}}
		{{$embed.Set "description" (print "You " $status " custom responses")}}
		{{$embed.Set "color" $successColor}}
		{{$economySettings.Set "enable-responses" $value}}
		{{dbSet 0 "EconomySettings" $economySettings}}
	{{else if eq $setting "workcd" "crimecd" "robcd" "incomecd"}}
		{{$cdType := reReplace "cd" $setting ""}}
		{{if lt (len .CmdArgs) 2}}
			{{$embed.Set "description" (print "No `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$duration := index .CmdArgs 1}}
		{{if not ($duration = toDuration $duration)}}
			{{$embed.Set "description" (print "Invalid `Value` argument provided.\nSyntax is: `" .Cmd " <Setting> <Value>`" $settings)}}
			{{sendMessage nil (cembed $embed)}}
			{{return}}
		{{end}}
		{{$embed.Set "description" (print "Sucessfully set the `" $cdType "Cooldown` to `" (humanizeDurationSeconds $duration) "`")}}
		{{$embed.Set "color" $successColor}}
		{{$duration = $duration.Seconds}}
		{{$crCD := print $cdType "Cooldown"}}
		{{$economySettings.Set $crCD $duration}}
		{{dbSet 0 "EconomySettings" $economySettings}}
	{{end}}
{{end}}
{{sendMessage nil (cembed $embed)}}