{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Decrypt`
©️ Ranger 2021
MIT License
*/}}

{{$a := (reReplace `(\()(\d)(\[;s\] \d-\[[0-9]+/\)[a-z])(\$\^)` .StrippedMsg "$2")}}
{{sendMessage nil $a}}