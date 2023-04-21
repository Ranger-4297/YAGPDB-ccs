{{/*
		Made by Ranger (765316548516380732)

	Trigger Type: `Regex`
	Trigger: `\A(-|<@!?204255221017214977>\s*)(set|config(ure)?)(\s+|\z)`

	©️ Ranger 2020-Present
	GNU, GPLV3 License
	Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$sC := 0x00ff7b}}
{{/* Configures economy settings */}}

{{/* Response */}}
{{$msg := sdict}}
{{$msg.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$msg.Set "timestamp" currentTime}}
{{$msg.Set "color" 0xFF0000}}
{{$db := (dbGet 0 "EconomySettings").Value}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
	{{$syntax := (print "\nAvailable settings: `max`, `min`, `betMax`, `startbalance`, `symbol`, `workCD`, `incomeCD`, `crimeCD`, `robCD`\nTo set it with the default settings `" $.Cmd " default`")}}
	{{with .CmdArgs}}
		{{if index $.CmdArgs 0}}
			{{$setting := (index $.CmdArgs 0) | lower}}
			{{$unable := (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")}}
			{{if eq $setting "default"}}
				{{$msg.Set "description" (print "Set the `EconomySettings` to default values")}}
				{{$msg.Set "color" $sC}}
				{{dbSet 0 "EconomySettings" (sdict "min" 200 "max" 500 "betMax" 5000 "symbol" "£" "startBalance" 200 "incomeCooldown" 300 "workCooldown" 7200 "crimeCooldown" 14400 "robCooldown" 21600)}}
				{{dbSet 0 "store" sdict}}
				{{dbSet 0 "russianRoulette" sdict}}
				{{dbSet 0 "bank"}}
			{{else}}
				{{with (dbGet 0 "EconomySettings")}}
					{{$a := sdict .Value}}
					{{$symbol := $a.symbol}}
					{{$nv := (print "No or invalid `value` argument passed.")}}
					{{if eq $setting "min" "max" "betmax"}}
						{{$smax := $a.max}}
						{{$smin := $a.min}}
						{{if gt (len $.CmdArgs) 1}}
							{{$val := (index $.CmdArgs 1)}}
							{{$ct := false}}
							{{$desc := ""}}
							{{if toInt $val}}
								{{if gt (toInt $val) 0}}
									{{if and (eq $setting "max") (lt (toInt $val) (toInt $smin))}}
										{{$desc = (print "You cannot set `" $setting "` to a value below `min`\n`min` is set to `" (humanizeThousands $smin) "`")}}
									{{else if and (eq $setting "min") (gt (toInt $val) (toInt $smax))}}
										{{$desc = (print "You cannot set `" $setting "` to a value above `max`\n`max` is set to `" (humanizeThousands $smax) "`")}}
									{{else}}
										{{$ct = true}}
									{{end}}
								{{else}}
									{{$desc = $unable}}
								{{end}}
							{{else}}
								{{$desc = $unable}}
							{{end}}
							{{if $ct}}
								{{$msg.Set "description" (print "You set `" $setting "` to " $symbol $val)}}
								{{$msg.Set "color" $sC}}
								{{if eq $setting "betmax"}}
									{{$setting = "betMax"}}
								{{end}}
								{{$db.Set $setting $val}}
								{{dbSet 0 "EconomySettings" $db}}
							{{else}}
								{{$msg.Set "description" $desc}}
							{{end}}
						{{else}}
							{{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")}}
						{{end}}
					{{else if eq $setting "startbalance"}}
						{{$oldStartBalance := $a.startBalance}}
						{{if gt (len $.CmdArgs) 1}}
							{{$startBalance := (index $.CmdArgs 1)}}
							{{if (toInt $startBalance)}}
								{{if gt (toInt $startBalance) 0}}
								{{$msg.Set "description" (print "You set `" $setting "` to " $symbol (humanizeThousands $startBalance) " from " (humanizeThousands $oldStartBalance))}}
								{{$msg.Set "color" $sC}}
								{{$db.Set "startBalance" $startBalance}}
								{{dbSet 0 "EconomySettings" $db}}
								{{else}}
									{{$msg.Set "description" $unable}}
								{{end}}
							{{else}}
								{{$msg.Set "description" $unable}}
							{{end}}
						{{else}}
							{{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")}}
						{{end}}
					{{else if eq $setting "symbol"}}
						{{if gt (len $.CmdArgs) 1}}
							{{$symbol = (index $.CmdArgs 1)}}
							{{$msg.Set "description" (print "You set the server currency symbol to " $symbol)}}
							{{$msg.Set "color" $sC}}
							{{$db.Set "symbol" $symbol}}
							{{dbSet 0 "EconomySettings" $db}}
						{{else}}
							{{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value>`")}}
						{{end}}
					{{else if eq $setting "workcd" "crimecd" "robcd" "incomecd"}}
						{{$cdType := reReplace "cd" $setting ""}}
						{{if gt (len $.CmdArgs) 1}}
							{{$dr := (index $.CmdArgs 1)}}
							{{if toDuration $dr}}
								{{$dr = toDuration $dr}}
								{{$msg.Set "description" (print "Sucessfully set the `" $cdType "Cooldown` to `" (humanizeDurationSeconds $dr) "`")}}
								{{$msg.Set "color" $sC}}
								{{$dr = $dr.Seconds}}
								{{$crCD := (print $cdType "Cooldown")}}
								{{$db.Set $crCD $dr}}
								{{dbSet 0 "EconomySettings" $db}}
							{{else}}
								{{$msg.Set "description" $unable}}
							{{end}}
						{{else}}
							{{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:Duration>`")}}
						{{end}}
					{{else}}
						{{$msg.Set "description" (print "No valid setting argument passed.\nSyntax is: `" $.Cmd " <Setting> <Value/>`" $syntax)}}
					{{end}}
				{{else}}
					{{$msg.Set "description" (print "No database found.\nPlease set it up with the default values using `" $.Cmd " default`")}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		{{$msg.Set "description" (print "No setting argument passed.\nSyntax is: `" $.Cmd " <Setting> <Value>`" $syntax)}}
	{{end}}
{{else}}
	{{$msg.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
{{end}}
{{sendMessage nil (cembed $msg)}}