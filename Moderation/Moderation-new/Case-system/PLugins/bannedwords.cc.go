{{/* 
        Made by Unknown User

    Trigger Type: `Command`
    Trigger: `bannedwordd`

    Usage: `bannedwords` | `bannedwords delete` | `bannedwords show`
*/}}


{{/*        Notes
    If you're aware on who wrote this, please PR a their current Username and UserID
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}


{{$args := parseArgs 1 "Correct Usage:\n\n`-bannedwords <your list of banned words>` to set a list of banned words and enable the banned words extension\n\n`-bannedwords delete` to delete the banned words list and disable the banned words extension\n\n`-bannedwords show` to show the current list of banned words" (carg "string" "list of banned words")}}

{{if eq (lower ($args.Get 0)) "delete"}}
    {{dbDel 0 "banned words"}}
    {{sendMessage nil "Deleted the banned words list and disabled the banned words extension"}}
{{else if eq (lower ($args.Get 0)) "show"}}
    {{(dbGet 0 "banned words").Value}}
{{else}}
    {{dbSet 0 "banned words" ($args.Get 0)}}
    {{sendMessage nil "All set :)"}}
{{end}}

{{deleteTrigger 1}}
{{deleteResponse 5}}