{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(roulette)(\s+|\z)`

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

{{/* Sell  */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with dbGet 0 "EconomySettings"}}
    {{$a := sdict .Value}}
	{{$symbol := $a.symbol}}
	{{$betMax := $a.betMax | toInt}}
	{{$incomeCooldown := $a.incomeCooldown | toInt}}
	{{$bal := or (dbGet $userID "cash").Value 0 | toInt}}
	{{with $.CmdArgs}}
        {{$numbers := seq 1 37}}
        {{$d1 := seq 1 13}}
        {{$d2 := seq 13 25}}
        {{$d3 := seq 25 37}}
        {{$c1 := cslice 1 4 7 10 13 16 19 22 25 28 31 34}}
        {{$c2 := cslice 2 5 8 11 14 17 20 23 26 29 32 35}}
        {{$c3 := cslice 3 6 9 12 15 18 21 24 27 30 33 36}}
        {{$h1 := seq 1 19}}
        {{$h2 := seq 19 37}}
        {{$red := cslice 1 3 5 7 9 12 14 16 18 21 23 25 27 30 32 34 36}}
        {{$black := cslice 2 4 6 8 10 11 13 15 17 20 22 24 26 28 29 31 33 35}}
        {{$even := cslice 2 4 6 8 10 12 14 16 18 20 22 24 26 28 30 32 34 36}}
        {{$odd := cslice 1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35}}
        {{if (index . 0)}}
            {{$side := (index . 0)}}
            {{if or (in $numbers $side) (eq (str $side) "1-12") (eq (str $side) "13-24") (eq (str $side) "25-36") (eq (lower (str $side)) "even") (eq (lower (str $side)) "odd") (eq (lower (str $side)) "red") (eq (lower (str $side)) "black") (eq (lower (str $side)) "1st") (eq (lower (str $side)) "2nd") (eq (lower (str $side)) "3rd") (eq (str $side) "1-18") (eq (str $side) "19-36")}}
                {{if gt (len .) 1}}
                    {{$bet := (index . 1)}}
                    {{$cont := false}}
                    {{if eq ($bet | toString) "all"}}
                        {{$bet = $bal}}
                    {{else if and $betMax (eq (toString $bet) "max")}}
                        {{$bet = $betMax}}
                    {{end}}
                    {{if $bet | toInt}}
                        {{$bet = $bet | toInt}}
                        {{if gt $bet 0}}
                            {{if le $bet $bal}}
                                {{$cont = true}}
                            {{else}}
                                {{$embed.Set "description" (print "You can't bet more than you have!")}}
                                {{$embed.Set "color" $errorColor}}
                            {{end}}
                            {{if $betMax}}
                                {{if gt $bet $betMax}}
                                    {{$embed.Set "description" (print "You can't bet more than " $symbol $betMax)}}
                                    {{$embed.Set "color" $errorColor}}
                                    {{$cont = false}}
                                {{end}}
                            {{end}}
                            {{if $cont}}
                                {{$land := randInt 37}}
                                {{if eq (toInt $side) $land}}
                                    {{/* bet*35 + bet */}}
                                    {{$bet = mult $bet 35 | add $bet }}
                                {{else if (or (and (eq (str $side) "1-12") (in $d1 $land)) (and (eq (str $side) "13-24") (in $d2 $land)) (and (eq (str $side) "25-36") (in $d3 $land)))}}
                                    {{/* bet*2 + bet */}}
                                    {{$bet = mult $bet 2 | add $bet}}
                                {{else if (or (and (eq (str $side) "1st") (in $c1 $land)) (and (eq (str $side) "2nd") (in $c2 $land)) (and (eq (str $side) "3rd") (in $c3 $land)))}}
                                    {{/* bet*2 + bet */}}
                                    {{$bet = mult $bet 2 | add $bet}}
                                {{else if (or (and (eq (str $side) "1-18") (in $h1 $land)) (and (eq (str $side) "19-36") (in $c2 $land)))}}
                                    {{/* bet*2 + bet */}}
                                    {{$bet = mult $bet 2 | add $bet}}
                                {{else if (or (and (eq (str $side) "even") (in $even $land)) (and (eq (str $side) "odd") (in $odd $land)))}}
                                    {{/* bet + bet */}}
                                    {{$bet = mult $bet 2}}
                                {{else if (or (and (eq (str $side) "red") (in $red $land)) (and (eq (str $side) "black") (in $black $land)))}}
                                    {{/* bet + bet */}}
                                    {{$bet = mult $bet 2}}
                                {{end}}
                                {{$space := ""}}
                                {{if in $black $land}}
                                    {{$space = "black"}}
                                {{else if in $red $land}}
                                    {{$space = "red"}}
                                {{end}}
                                {{$embed.Set "description" (print "The ball landed on " $space " " $land)}}
                                {{$embed.Set "color" $successColor}}
                            {{end}}
                        {{else}}
                            {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount> <Space>`")}}
                            {{$embed.Set "color" $errorColor}}
                        {{end}}
                    {{else}}
                        {{$embed.Set "description" (print "Invalid `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount> <Space>`")}}
                        {{$embed.Set "color" $errorColor}}
                    {{end}}
                {{else}}
                    {{$embed.Set "description" (print "No `Bet` argument provided.\nSyntax is `" $.Cmd " <Bet:Amount> <Space>`")}}
                    {{$embed.Set "color" $errorColor}}
                {{end}}
            {{else}}
                {{$embed.Set "description" (print "Invalid `Space` argument provided.\nSyntax is `" $.Cmd " <Space> <Bet:Amount>`")}}
                {{$embed.Set "color" $errorColor}}
            {{end}}
        {{end}}
    {{else}}
        {{$embed.Set "description" (print "No `Space` argument provided.\nSyntax is `" $.Cmd " <Space> <Bet:Amount>`")}}
        {{$embed.Set "color" $errorColor}}
    {{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "server-set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}