{{/* 
        Made by Unknown User
        Edited by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `bannedwords`

No license
*/}}


{{/*        Notes
    If you're aware on who wrote this, please PR a their current Username and UserID
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}


{{$args := parseArgs 1 "Correct Usage:\n\n`-bannedwords <your list of bannedwords>` to set a list of bannedwords and enable the bannedwords extension\n\n`-bannedwords delete` to delete the bannedwords list and disable the bannedwords extension\n\n`-bannedwords show` to show the current list of bannedwords" (carg "string" "list of bannedwords")}}

{{if eq (lower ($args.Get 0)) "delete"}}
    {{dbDel 0 "bannedwords"}}
    {{sendMessage nil "Deleted the bannedwords list and disabled the bannedwords extension"}}
{{else if eq (lower ($args.Get 0)) "show"}}
    {{(dbGet 0 "bannedwords").Value}}
{{else}}
    {{dbSet 0 "bannedwords" ($args.Get 0)}}
    {{sendMessage nil "All set :)"}}
{{end}}

{{deleteTrigger 1}}
{{deleteResponse 5}}