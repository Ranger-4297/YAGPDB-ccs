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
{{$successColor := 0x00ff7b}}
{{/* Configures economy settings */}}

{{/* Response */}}
{{$msg := sdict}}
{{$msg.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$msg.Set "timestamp" currentTime}}
{{$msg.Set "color" 0xFF0000}}
{{$db := (dbGet 0 "EconomySettings").Value}}
{{$perms := split (index (split (exec "viewperms" ) "\n" ) 2) ", "}}
{{if or (in $perms "Administrator") (in $perms "ManageServer")}}
    {{$syntax := (print "\nAvailable settings: `max`, `min`, `startbalance`, `symbol`, `cooldown`\nTo set it with the default settings `" $.Cmd " default`")}}
    {{with .CmdArgs}}
        {{if index $.CmdArgs 0}}
            {{$setting :=  (index $.CmdArgs 0) | lower}}
            {{$unable := (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")}}
            {{if eq $setting "default"}}
                {{$msg.Set "description" (print "Set the `EconomySettings` to the default values")}}
                {{$msg.Set "color" $successColor}}
                {{dbSet 0 "EconomySettings" (sdict "min" 200 "max" 500 "symbol" "£" "startBalance" 200 "crimeCooldown" 3600 "robCooldown" 21600)}}
            {{else}}
                {{with (dbGet 0 "EconomySettings")}}
                    {{$a := sdict .Value}}
                    {{$nv := (print "No `value` argument passed.")}}
                    {{if eq $setting "max"}}
                        {{$min := toString $a.min}}
                        {{$symbol := $a.symbol}}
                        {{if gt (len $.CmdArgs) 1}}
                            {{$max := (index $.CmdArgs 1)}}
                            {{if toInt $max}}
                                {{if lt (toInt $max) (toInt $min)}}
                                    {{$msg.Set "description" (print "You cannot set `" $setting "` to a value below `min`\n`min` is set to `" $min "`")}}
                                {{else}}
                                    {{$msg.Set "description" (print "You set `" $setting "` to " $symbol $max)}}
                                    {{$msg.Set "color" $successColor}}
                                    {{$db.Set "max" $max}}
                                    {{dbSet 0 "EconomySettings" $db}}
                                {{end}}
                            {{else}}
                                {{$msg.Set "description" $unable}}
                            {{end}}
                        {{else}}
                            {{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")}}
                        {{end}}
                    {{else if eq $setting "min"}}
                        {{$max := toString $a.max}}
                        {{$symbol := $a.symbol}}
                        {{if gt (len $.CmdArgs) 1}}
                            {{$min := (index $.CmdArgs 1)}}
                            {{if toInt $min}}
                                {{if gt (toInt $min) (toInt $max)}}
                                    {{$msg.Set "description" (print "You cannot set `" $setting "` to a value above `max`\n`max` is set to `" $max "`")}}
                                {{else}}
                                    {{$msg.Set "description" (print "You set `" $setting "` to " $symbol $min)}}
                                    {{$msg.Set "color" $successColor}}
                                    {{$db.Set "min" $min}}
                                    {{dbSet 0 "EconomySettings" $db}}
                                {{end}}
                            {{else}}
                                {{$msg.Set "description" $unable}}
                            {{end}}
                        {{else}}
                            {{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")}}
                        {{end}}
                    {{else if eq $setting "startbalance"}}
                        {{$symbol := $a.symbol}}
                        {{$oldStartBalance := $a.startBalance}}
                        {{if gt (len $.CmdArgs) 1}}
                            {{$startBalance := (index $.CmdArgs 1)}}
                            {{if (toInt $startBalance)}}
                                {{$msg.Set "description" (print "You set `" $setting "` to " $symbol $startBalance " from " $oldStartBalance)}}
                                {{$msg.Set "color" $successColor}}
                                {{$db.Set "startBalance" $startBalance}}
                                {{dbSet 0 "EconomySettings" $db}}
                            {{else}}
                            {{$msg.Set "description" $unable}}
                            {{end}}
                        {{else}}
                            {{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")}}
                        {{end}}
                    {{else if eq $setting "symbol"}}
                        {{if gt (len $.CmdArgs) 1}}
                            {{$symbol := (index $.CmdArgs 1)}}
                            {{$output := ""}}
                            {{if (reFind `(<a?:[A-z+]+\:\d{17,19}>)` $symbol)}}
                                {{$output = $symbol}}
                            {{else}}
                                {{$output = (print "`" $symbol "`")}}
                            {{end}}
                            {{$msg.Set "description" (print "You set the server currency symbol to " $output )}}
                            {{$msg.Set "color" $successColor}}
                            {{$db.Set "symbol" $symbol}}
                            {{dbSet 0 "EconomySettings" $db}}
                        {{else}}
                            {{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:String>`")}}
                        {{end}}
                    {{else if eq $setting "cooldown"}}
                        {{if gt (len $.CmdArgs) 1}}
                            {{$cdType := (index $.CmdArgs 1)}}
                            {{if eq $cdType "crime"}}
                                {{if gt (len $.CmdArgs) 2}}
                                    {{$dr := toString (index $.CmdArgs 2)}}
                                    {{if toDuration $dr}}
                                        {{$dr = toDuration $dr}}
                                        {{$msg.Set "description" (print "Sucessfully set the `crimeCooldown` to `" $dr "`")}}
                                        {{$msg.Set "color" $successColor}}
                                        {{$dr = $dr.Seconds}}
                                        {{$db.Set "crimeCooldown" $dr}}
                                        {{dbSet 0 "EconomySettings" $db}}
                                    {{else}}
                                        {{$msg.Set "description" $unable}}
                                    {{end}}
                                {{else}}
                                    {{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:Duration>`")}}
                                {{end}}
                            {{else if eq $cdType "rob"}}
                                {{if gt (len $.CmdArgs) 2}}
                                    {{$dr := toString (index $.CmdArgs 2)}}
                                    {{if toDuration $dr}}
                                        {{$dr = toDuration $dr}}
                                        {{$msg.Set "description" (print "Sucessfully set the `robCooldown` to `" $dr "`")}}
                                        {{$msg.Set "color" $successColor}}
                                        {{$dr = $dr.Seconds}}
                                        {{$db.Set "robCooldown" $dr}}
                                        {{dbSet 0 "EconomySettings" $db}}
                                    {{else}}
                                        {{$msg.Set "description" $unable}}
                                    {{end}}
                                {{else}}
                                    {{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Value:Duration>`")}}
                                {{end}}
                            {{end}}
                        {{else}}
                            {{$msg.Set "description" (print $nv "\nSyntax is: `" $.Cmd " " $setting " <Type:Rob/Crime> <Value:Duration>`")}}
                        {{end}}
                    {{else}}
                        {{$msg.Set "description" (print "No valid setting argument passed.\nSyntax is: `" $.Cmd " <Setting:String> <Value:String/Int/Duration>`" $syntax)}}
                    {{end}}
                {{else}}
                    {{$msg.Set "description" (print "No database found.\nPlease set it up with the default values using `" $.Cmd " default`")}}
                {{end}}
            {{end}}
        {{end}}
    {{else}}
        {{$msg.Set "description" (print "No setting argument passed.\nSyntax is: `" $.Cmd " <Setting:String> <Value:String/Int/Duration>`" $syntax)}}
    {{end}}
{{else}}
    {{$msg.Set "description" (print "Insufficient permissions.\nTo use this command you need to have either `Administrator` or `ManageServer` permissions")}}
{{end}}
{{sendMessage nil (cembed $msg)}}