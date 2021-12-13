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
{{$errorColor := 0xFF0000}}

{{/* Configures economy settings */}}

{{/*
If there is no setting values
You'll be asked to set it up with default values
You can change these later
*/}}

{{/* Response */}}
{{with .CmdArgs}}
    {{if index . 0}}
        {{$setting := (lower (index . 0))}}
        {{if eq $setting "failrate"}}
            {{with (dbGet 0 "EconomySettings")}}
                {{if gt (len $.CmdArgs) 1}}
                    {{$failRate := (index $.CmdArgs 1)}}
                    {{if (toInt $failRate)}}
                        {{if lt (toInt $failRate) 20}}
                            {{sendMessage nil (cembed
                                "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                "description" (print "You cannot set your `" $setting "` below 20%")
                                "color" $errorColor
                                "timestamp" currentTime
                                )}}
                        {{else if gt (toInt $failRate) 100}}
                            {{sendMessage nil (cembed
                                "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                "description" (print "You cannot set your `" $setting "` above 100%")
                                "color" $errorColor
                                "timestamp" currentTime
                                )}}
                        {{else}}
                            {{sendMessage nil (cembed
                                "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                "description" (print "You set `" $setting "` to " $failRate "%")
                                "color" $successColor
                                "timestamp" currentTime
                                )}}
                            {{$sdict := (dbGet 0 "EconomySettings").Value}}
                            {{$sdict.Set "failRate" $failRate }}
                            {{dbSet 0 "EconomySettings" $sdict}}
                        {{end}}
                    {{else}}
                        {{sendMessage nil (cembed
                                "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                "description" (print "You're unable to set the `" $setting "` to this value, check that you used a valid number above 1")
                                "color" $errorColor
                                "timestamp" currentTime
                                )}}
                    {{end}}
                {{else}}
                    {{sendMessage nil (cembed
                                "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                                "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")
                                "color" $errorColor
                                "timestamp" currentTime
                                )}}
                {{end}}
            {{else}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $.Cmd " default`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{end}}
        {{else if eq $setting "max"}}
            {{with (dbGet 0 "EconomySettings")}}
                {{$a := sdict .Value}}
                {{$min := (toInt $a.min)}}
                {{$symbol := $a.symbol}}
                {{if gt (len $.CmdArgs) 1}}
                    {{$max := (index $.CmdArgs 1)}}
                    {{if (toInt $max)}}
                        {{if lt (toInt $max) (toInt $min)}}
                            {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You cannot set `" $setting "` to a value below `min`\n`min` is set to `" $min "`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                        {{else}}
                            {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You set `" $setting "` to " $symbol $max)
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                            {{$sdict := (dbGet 0 "EconomySettings").Value}}
                            {{$sdict.Set "max" $max }}
                            {{dbSet 0 "EconomySettings" $sdict}}
                        {{end}}
                    {{else}}
                        {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                    {{end}}
                {{else}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{end}}
            {{else}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $.Cmd " default`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{end}}
        {{else if eq $setting "min"}}
            {{with (dbGet 0 "EconomySettings")}}
                {{$a := sdict .Value}}
                {{$max := (toInt $a.max)}}
                {{$symbol := $a.symbol}}
                {{if gt (len $.CmdArgs) 1}}
                    {{$min := (index $.CmdArgs 1)}}
                    {{if (toInt $min)}}
                        {{if gt (toInt $min) (toInt $max)}}
                            {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You cannot set `" $setting "` to a value above `max`\n`max` is set to `" $max "`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                        {{else}}
                            {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You set `" $setting "` to " $symbol $min)
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                            {{$sdict := (dbGet 0 "EconomySettings").Value}}
                            {{$sdict.Set "min" $min }}
                            {{dbSet 0 "EconomySettings" $sdict}}
                        {{end}}
                    {{else}}
                        {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                    {{end}}
                {{else}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{end}}
            {{else}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $.Cmd " default`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{end}}
        {{else if eq $setting "startbalance"}}
            {{with (dbGet 0 "EconomySettings")}}
                {{$a := sdict .Value}}
                {{$symbol := $a.symbol}}
                {{$oldStartBalance := $a.startBalance }}
                {{if gt (len $.CmdArgs) 1}}
                    {{$startBalance := (index $.CmdArgs 1)}}
                    {{if (toInt $startBalance)}}
                        {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You set `" $setting "` to " $symbol $startBalance " from " $oldStartBalance)
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                        {{$sdict := (dbGet 0 "EconomySettings").Value}}
                        {{$sdict.Set "startBalance" $startBalance }}
                        {{dbSet 0 "EconomySettings" $sdict}}
                    {{else}}
                        {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to set `" $setting "` to this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                    {{end}}
                {{else}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{end}}
            {{else}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $.Cmd " default`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{end}}
        {{else if eq $setting "symbol"}}
            {{with (dbGet 0 "EconomySettings")}}
                {{if gt (len $.CmdArgs) 1}}
                    {{$symbol := (index $.CmdArgs 1)}}
                    {{$output := ""}}
                    {{if (reFind `(<a?:[A-z+]+\:\d{17,19}>)` $symbol)}}
                        {{$output = $symbol}}
                    {{else}}
                        {{$output = (print "`" $symbol "`")}}
                    {{end}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You set the server currency symbol to " $output )
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                    {{$sdict := (dbGet 0 "EconomySettings").Value}}
                    {{$sdict.Set "symbol" $symbol }}
                    {{dbSet 0 "EconomySettings" $sdict}}
                {{else}}
                    {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `value` argument passed.\nSyntax is: `" $.Cmd " " $setting " <Value:Int>`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                {{end}}
            {{else}}
                {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $.Cmd " default`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{end}}
        {{else if eq $setting "default"}}
            {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "Set the `EconomySettings` to the default values")
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
            {{dbSet 0 "EconomySettings" (sdict "min" 200 "max" 500 "failRate" 20 "symbol" "£" "startBalance" 200)}}
        {{else}}
            {{sendMessage nil (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128")) 
                            "description" (print "No valid setting argument passed.\nSyntax is: `" $.Cmd " <Setting:String> <Value:String/Int>`\nAvailable settings: `failrate`, `max`, `min`, `startbalance`, `symbol`\nAvailable value types `int` `string`\nTo set it with the default settings `" $.Cmd " default`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
        {{end}}
    {{end}}
{{else}}
    {{sendMessage nil (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128")) 
            "description" (print "No setting argument passed.\nSyntax is: `" $.Cmd " <Setting:String> <Value:String/Int>`\nAvailable settings: `failrate`, `max`, `min`, `startbalance`, `symbol`\nAvailable value types `int` `string`\nTo set it with the default settings `" $.Cmd " default`")
            "color" $errorColor
            "timestamp" currentTime
            )}}
{{end}}