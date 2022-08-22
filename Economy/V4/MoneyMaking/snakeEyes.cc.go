{{/*
        Made by Ranger (765316548516380732)

        Trigger Type: `Regex`
        Trigger: `\A(-|<@!?204255221017214977>\s*)(snake?-?eyes)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := (index (reFindAllSubmatches `.*?: \x60(.*)\x60\z` (execAdmin "Prefix")) 0 1 )}}

{{/* Snake Eyes */}}

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
    {{with dbGet $userID "EconomyInfo"}}
        {{$a = sdict .Value}}
        {{$bal := $a.cash}}
        {{$die1 := (randInt 1 7)}}
        {{$die2 := (randInt 1 7)}}
        {{with $.CmdArgs}}
            {{$bet := (index . 0)}}
            {{if (toInt $bet)}}
                {{if gt (toInt $bet) 0}}
                    {{if le (toInt $bet) (toInt $bal)}}
                        {{if not ($cooldown := dbGet $userID "snakeeyesCooldown")}}
                            {{dbSetExpire $userID "snakeeyesCooldown" "cooldown" $incomeCooldown}}
                            {{$newCash := (sub $bal $bet)}}
                            {{if and (eq $die1 1) (eq $die2 1)}}
                                {{$embed.Set "description" (print "You rolled snake eyes (" $die1 "&" $die2 ")\nAnd won " $symbol (humanizeThousands (mult $bet 36)))}}
                                {{$embed.Set "color" $successColor}}
                                {{$newCash = (add $bal (mult $bet 36))}}
                            {{else}}
                                {{$embed.Set "description" (print "You rolled " $die1 "&" $die2 " and lost " $symbol (humanizeThousands $bet) ".")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                            {{$a.Set "cash" $newCash}}
                            {{dbSet $userID "EconomyInfo" $a}}
                        {{else}}
                            {{$embed.Set "description" (print "This command is on cooldown for " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)))}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{else}}
                        {{$embed.Set "description" (print "You can't bet more than you have")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{else}}
                {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
                {{$embed.Set "color" $errorColor}}
            {{end}}
        {{else}}
            {{$embed.Set "description" (print "No `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount>`")}}
            {{$embed.Set "color" $errorColor}}
        {{end}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}