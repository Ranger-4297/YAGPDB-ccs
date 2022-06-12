{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Decrypt`

    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{$a := (reReplace `(\d{1,2}[}{[]\W[A-z][}{\]] )(\d)(\.)([A-z]\d\W[A-z]{1,2})(\W{1,2})` .StrippedMsg "$2")}}
{{sendMessage nil $a}}