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
{{$sC := 0x00ff7b}}

{{/* server set */}}

{{/* Response */}}
{{$msg := sdict}}
{{$msg.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$msg.Set "timestamp" currentTime}}
{{$msg.Set "color" 0xFF0000}}
{{$db := (dbGet 0 "EconomySettings").Value}}
{{$perms := split (index (split (exec "viewperms") "\n") 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
	{{$syntax := (print "\nAvailable settings: `max`, `min`, `betMax`, `startbalance`, `symbol`, `workCD`, `incomeCD`, `crimeCD`, `robCD`, `responses`\nTo set it with the default settings `" $.Cmd " default`")}}
	{{with .CmdArgs}}
		{{if index $.CmdArgs 0}}
			{{$setting := (index $.CmdArgs 0) | lower}}
			{{$unable := (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")}}
			{{if eq $setting "default"}}
				{{$msg.Set "description" (print "Set the `EconomySettings` to default values")}}
				{{$msg.Set "color" $sC}}
				{{dbSet 0 "EconomySettings" (sdict "min" 200 "max" 500 "betMax" 5000 "symbol" "£" "startBalance" 200 "incomeCooldown" 300 "workCooldown" 7200 "crimeCooldown" 14400 "robCooldown" 21600 "enable-responses" false "responses" (sdict "work" cslice "crime" cslice))}}
				{{dbSet 0 "store" (sdict "items" sdict)}}
				{{dbSet 0 "russianRoulette" sdict}}
				{{dbSet 0 "bank" sdict}}
				{{dbSet 0 "roulette" (sdict "game" sdict "storage" sdict)}}
				{{dbSet 0 "accounts" sdict}}
				{{dbSet 0 "cardGames" (sdict "cards" (sdict "AC" "<:AC:1182102116190392360>" "AS" "<:AS:1182102121886273648>" "AH" "<:AH:1182102120644743258>" "AD" "<:AD:1182102119063490611>" "2C" "<:2C:1182102480218239086>" "2S" "<:2S:1182102485247201400>" "2H" "<:2H:1182102483909226586>" "2D" "<:2D:1182102481715605595>" "3C" "<:3C:1182102658857840780>" "3S" "<:3S:1182102664125878422>" "3H" "<:3H:1182102662834044968>" "3D" "<:3D:1182102661328289924>" "4C" "<:4C:1182102790953238558>" "4S" "<:4S:1182102795923488869>" "4H" "<:4H:1182102795051089961>" "4D" "<:4D:1182102793767637102>" "5C" "<:5C:1182102882833670245>" "5S" "<:5S:1182102887980081193>" "5H" "<:5H:1182102886612742164>" "5D" "<:5D:1182102885148938351>" "6C" "<:6C:1182102934088065186>" "6S" "<:6S:1182102939016384553>" "6H" "<:6H:1182102937502228550>" "6D" "<:6D:1182102936193605703>" "7C" "<:7C:1182102982372896879>" "7S" "<:7S:1182102987917774868>" "7H" "<:7H:1182102986411999252>" "7D" "<:7D:1182102984746881104>" "8C" "<:8C:1182103038341693511>" "8S" "<:8S:1182103045874651246>" "8H" "<:8H:1182103043072860190>" "8D" "<:8D:1182103039696457811>" "9C" "<:9C:1182103085594706000>" "9S" "<:9S:1182103094511796254>" "9H" "<:9H:1182103093098332311>" "9D" "<:9D:1182103088530731101>" "10C" "<:10C:1183212851830210620>" "10S" "<:10D:1183212854246113291>" "10H" "<:10H:1183212856083222538>" "10D" "<:10D:1183212854246113291>" "JC" "<:JC:1182103188636184617>" "JS" "<:JS:1182103197700083792>" "JH" "<:JH:1182103194973769778>" "JD" "<:JD:1182103192000008243>" "QC" "<:QC:1182103800820011059>" "QS" "<:QS:1182103390008913960>" "QH" "<:QH:1182103415644491776>" "QD" "<:QD:1182103437383577601>" "KC" "<:KC:1182754284878700696>" "KS" "<:KS:1182103911415431219>" "KH" "<:KH:1182754289500827718>" "KD" "<:KD:1182754287525302383>" "CB" "<:cardback:1248784417598603284>"))}}
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
							{{else if and (eq $setting "betmax") (eq (toInt $val) 0)}}
								{{$ct = true}}
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
					{{else if eq $setting "responses"}}
						{{if gt (len $.CmdArgs) 1}}
							{{$value := (index $.CmdArgs 1) | lower}}
							{{if eq $value "yes" "enable" "enabled" "no" "disable" "disabled"}}
								{{$status := ""}}
								{{if eq $value "yes" "enable" "enabled"}}
									{{$status = "enabled"}}
									{{$value = true}}
								{{else}}
									{{$status = "disabled"}}
									{{$value = false}}
								{{end}}
								{{$msg.Set "description" (print "You " $status " custom responses")}}
								{{$msg.Set "color" $sC}}
								{{$db.Set "enable-responses" $value}}
								{{dbSet 0 "EconomySettings" $db}}
							{{else}}
								{{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value>`")}}
							{{end}}
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