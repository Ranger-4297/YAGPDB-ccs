{{/*
        Made by Ranger (765316548516380732)

        Trigger Type: `Regex`
        Trigger: `\A(-|<@!?204255221017214977>\s*)(rollnum(ber)?|rn)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1}}

{{/* Roll */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
    {{$a := sdict .Value}}
    {{$symbol := $a.symbol}}
    {{$incomeCooldown := $a.incomeCooldown | toInt}}
    {{if not (dbGet $userID "EconomyInfo")}}
        {{dbSet $userID "EconomyInfo" (sdict "cash" 200 "bank" 0)}}
    {{end}}
    {{with (dbGet $userID "EconomyInfo")}}
        {{$a = sdict .Value}}
        {{$bal := $a.cash}}
        {{with $.CmdArgs}}
            {{$bet := (index . 0)}}
            {{$continue := false}}
            {{if $bet | toInt}}
                {{if gt (toInt $bet) 0}}
                    {{if le (toInt $bet) (toInt $bal)}}
                        {{$continue = true}}
                    {{else}}
                        {{$embed.Set "description" (print "You can't bet more than you have!")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{else}}
                {{$bet = $bet | lower}}
                {{if eq $bet "all"}}
                    {{if ne (toInt $bet) 0}}
                        {{$continue = true}}
                    {{else}}
                        {{$embed.Set "description" (print "You had no money to bet.")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{end}}
            {{if $continue}}
                {{if not ($cooldown := dbGet $userID "rollCooldown")}}
                    {{dbSetExpire $userID "rollCooldown" "cooldown" $incomeCooldown}}
                    {{$roll := (randInt 1 101)}}
                    {{$amount := ""}}
                    {{$newBal := ""}}
                    {{if and (ge $roll 65) (lt $roll 90)}}
                        {{$amount = $bet}}
                        {{$newBal = (add $bal $amount)}}
                    {{else if and (ge $roll 90) (lt $roll 100)}}
                        {{$amount = (mult $bet 3)}}
                        {{$newBal = (add $bal $amount)}}
                    {{else if eq $roll 100}}
                        {{$amount = (mult $bet 5)}}
                        {{$newBal = (add $bal $amount)}}
                    {{else}}
                        {{$newBal = (sub $bal $bet)}}
                    {{end}}
                    {{if gt $newBal $bal}}
                        {{$embed.Set "description" (print "The roll landed on " $roll " and you and won " (humanizeThousands $amount) "!")}}
                        {{$embed.Set "color" $successColor}}
                    {{else}}
                        {{$embed.Set "description" (print "The roll landed on " $roll " and you and lost " (humanizeThousands $bet))}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                    {{$a.Set "cash" $newBal}}
                    {{dbSet $userID "EconomyInfo" $a}}
                {{else}}
                    {{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{end}}
        {{else}}
            {{$embed.Set "description" (print "No `bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
            {{$embed.Set "color" $errorColor}}
        {{end}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}