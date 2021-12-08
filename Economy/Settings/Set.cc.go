{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `set`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$args := parseArgs 0 "" (carg "string" "Setting") (carg "string" "Value")}}
{{$errorColor := 0xFF0000}}
{{$successColor := 0x00ff8b}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 }}
{{if not ($args.Get 0)}}
    {{$errorEmbed := (cembed
            "description" (print "Not enough arguments passed.\nSyntax is: `" $prefix "set <Setting:String> <Value:String/Int>`\nAvailable settings: `failrate`, `max`, `min`, `startbalance`, `symbol`\nAvailable value types `int` `string`")
            "color" $errorColor
            )}}
    {{sendMessage nil $errorEmbed}}
{{end}}

{{with and ($args.Get 0)}}
    {{$Setting := (lower ($args.Get 0))}}
    {{with (dbGet 0 "EconomySettings")}}
        {{if eq $Setting "failrate"}}
            {{if not ($args.Get 1)}}
                {{$errorEmbed := (cembed
                            "description" (print "Not enough arguments passed.\nSyntax is: `" $prefix "set " $Setting " <Value:Int>`")
                            "color" $errorColor
                )}}
                {{sendMessage nil $errorEmbed}}
            {{else}}
                {{$a := sdict .Value}}
                {{$failRate := $a.failRate}}
                {{$newFailRate := ($args.Get 1)}}
                {{if (toInt $newFailRate)}}
                    {{if lt (toInt $newFailRate) 20}}
                        {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You cannot set your fail-rate below 20%")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                        {{sendMessage nil $errorEmbed}}
                    {{else if gt (toInt $newFailRate) 100}}
                        {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You cannot set your fail-rate above 100%")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                        {{sendMessage nil $errorEmbed}}
                    {{else}}
                        {{$updateEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You set your fail-rate to " $newFailRate "% from " $failRate "%")
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                        {{sendMessage nil $updateEmbed}}
                        {{$sdict := (dbGet 0 "EconomySettings").Value}}
                        {{$sdict.Set "failRate" $newFailRate}}
                        {{dbSet 0 "EconomySettings" $sdict}}
                    {{end}}
                {{else}}
                    {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to set the fail-rate to this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                    {{sendMessage nil $errorEmbed}}
                {{end}}
            {{end}}
        {{else if eq $Setting "max"}}
            {{$a := sdict .Value}}
            {{$symbol := $a.symbol}}
            {{$min := $a.min}}
            {{$newMax := ($args.Get 1)}}
            {{if not ($args.Get 1)}}
                {{$errorEmbed := (cembed
                            "description" (print "Not enough arguments passed.\nSyntax is: `" $prefix "set " $Setting " <Value:Int>`")
                            "color" $errorColor
                )}}
                {{sendMessage nil $errorEmbed}}
            {{else}}
                {{if (toInt $newMax)}}
                    {{if lt (toInt $newMax) (toInt $min)}}
                        {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You cannot set `max` to a value below `min`.\nYour min is set to `" $symbol $min "`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                        {{sendMessage nil $errorEmbed}}
                    {{else}}
                        {{$updateEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "Successfully set `max` to  " $symbol $newMax)
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                        {{sendMessage nil $updateEmbed}}
                        {{$sdict := (dbGet 0 "EconomySettings").Value}}
                        {{$sdict.Set "max" $newMax}}
                        {{dbSet 0 "EconomySettings" $sdict}}
                    {{end}}
                {{else}}
                    {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to set the max this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                    {{sendMessage nil $errorEmbed}}
                {{end}}
            {{end}}
        {{else if eq $Setting "min"}}
            {{$a := sdict .Value}}
            {{$symbol := $a.symbol}}
            {{$max := $a.max}}
            {{$newMin := ($args.Get 1)}}
            {{if not ($args.Get 1)}}
                {{$errorEmbed := (cembed
                            "description" (print "Not enough arguments passed.\nSyntax is: `" $prefix "set " $Setting " <Value:Int>`")
                            "color" $errorColor
                )}}
                {{sendMessage nil $errorEmbed}}
            {{else}}
                {{if (toInt $newMin)}}
                    {{if gt (toInt $newMin) (toInt $max)}}
                        {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You cannot set `min` to a value above `max`.\nYour max is set to `" $symbol $max "`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                        {{sendMessage nil $errorEmbed}}
                    {{else}}
                        {{$updateEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "Successfully set `min` to " $symbol $newMin)
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                        {{sendMessage nil $updateEmbed}}
                        {{$sdict := (dbGet 0 "EconomySettings").Value}}
                        {{$sdict.Set "min" $newMin}}
                        {{dbSet 0 "EconomySettings" $sdict}}
                    {{end}}
                {{else}}
                    {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to set the max this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                    {{sendMessage nil $errorEmbed}}
                {{end}}
            {{end}}
        {{else if eq $Setting "startbalance"}}
            {{$a := sdict .Value}}
            {{$symbol := $a.symbol}}
            {{$startBalance := $a.startBalance}}
            {{$newStartBalance := ($args.Get 1)}}
            {{if not ($args.Get 1)}}
                {{$errorEmbed := (cembed
                            "description" (print "Not enough arguments passed.\nSyntax is: `" $prefix "set " $Setting " <Value:Int>`")
                            "color" $errorColor
                )}}
                {{sendMessage nil $errorEmbed}}
            {{else}}
                {{if (toInt $newStartBalance)}}
                    {{$updateEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You set the server start-balance to " $symbol $newStartBalance " from " $symbol $startBalance)
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                    {{sendMessage nil $updateEmbed}}
                    {{$sdict := (dbGet 0 "EconomySettings").Value}}
                    {{$sdict.Set "startBalance" $newStartBalance}}
                    {{dbSet 0 "EconomySettings" $sdict}}
                {{else}}
                    {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You're unable to set the start-balance this value, check that you used a valid number above 1")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
                    {{sendMessage nil $errorEmbed}}
                {{end}}
            {{end}}
        {{else if eq $Setting "symbol"}}
            {{$a := sdict .Value}}
            {{$newSymbol := ($args.Get 1)}}
            {{if not ($args.Get 1)}}
                {{$errorEmbed := (cembed
                            "description" (print "Not enough arguments passed.\nSyntax is: `" $prefix "set " $Setting " <Value:Char/Emoji>`")
                            "color" $errorColor
                )}}
                {{sendMessage nil $errorEmbed}}
            {{else}}
                {{$symbol := $a.symbol}}
                {{$updateEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You set the server currency symbol to " $newSymbol " from " $symbol)
                            "color" $successColor
                            "timestamp" currentTime
                            )}}
                {{sendMessage nil $updateEmbed}}
                {{$sdict := (dbGet 0 "EconomySettings").Value}}
                {{$sdict.Set "symbol" $newSymbol}}
                {{dbSet 0 "EconomySettings" $sdict}}
            {{end}}
        {{else}}
            {{$errorEmbed := (cembed
                            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "128"))
                            "description" (print "You have no provide a valid setting or value.\nSyntax is: `" $prefix "set <Setting:String> <Value:String/Int>`\nAvailable settings: `failrate`, `max`, `min`, `startbalance`, `symbol`\nAvailable value types `int` `string`")
                            "color" $errorColor
                            "timestamp" currentTime
                            )}}
            {{sendMessage nil $errorEmbed}}
        {{end}}
    {{end}}
{{end}}
